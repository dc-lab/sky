package job_manager

import (
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	_ "gorm.io/gorm/clause"
)

type Repo struct {
	db *gorm.DB
}

func OpenRepo(dsn string) (*Repo, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	err = db.AutoMigrate(&Job{}, &Task{}, &Agent{}, &TaskExecution{})
	if err != nil {
		return nil, err
	}

	return &Repo{db}, nil
}

func (r *Repo) Transaction(cb func(tx *Repo) error) error {
	return r.db.Transaction(func(dbtx *gorm.DB) error {
		return cb(&Repo{dbtx})
	})
}

func (r *Repo) AddJob(job *Job) error {
	if len(job.Tasks) < 1 {
		return errors.New("Cannot create job without tasks")
	}

	log.WithField("job_name", job.Name).Infoln("Add job")

	job.Status = JobCreated
	taskIdByName := make(map[string]int)
	for i, task := range job.Tasks {
		if _, found := taskIdByName[task.Name]; found {
			return fmt.Errorf("Task names should be unique; found duplicate task %s", task.Name)
		}
		taskIdByName[task.Name] = i
		if len(task.Dependencies) == 0 {
			job.Tasks[i].Status = JobCreated
		} else {
			job.Tasks[i].Status = JobWaiting
		}
	}

	for i, task := range job.Tasks {
		job.Tasks[i].PendingDependencies = make([]*Task, len(task.Dependencies))
		for j, dep := range task.Dependencies {
			depTask, found := taskIdByName[dep]
			if !found {
				return fmt.Errorf("Unknown dependency %s for task %s", dep, task.Name)
			}
			job.Tasks[i].PendingDependencies[j] = &job.Tasks[depTask]
		}
	}
	loop := validateGraphIsAcyclic(job)
	if loop != nil {
		var str strings.Builder
		str.WriteString("Found dependency loop: ")
		for i, name := range *loop {
			if i > 0 {
				str.WriteString(" ")
			}
			str.WriteString(name)
		}
		return errors.New(str.String())
	}

	res := r.db.Debug().Create(job)
	if err := res.Error; err != nil {
		return err
	}

	log.WithField("job_id", job.ID).Infoln("Created job")
	for _, task := range job.Tasks {
		log.WithFields(log.Fields{
			"task_id": task.ID,
			"job_id":  task.JobID,
		}).Infoln("Created task")
	}

	return nil
}

type NodeVisitedType int

const (
	NodeNotVisited NodeVisitedType = iota
	NodeInCurrentPath
	NodeVisited
)

func validateGraphIsAcyclic(job *Job) (loop *[]string) {
	visited := make(map[string]NodeVisitedType)
	for _, task := range job.Tasks {
		if visited[task.Name] == NodeVisited {
			continue
		}

		if loop := validateComponentIsAcyclic(&task, &visited); loop != nil {
			return loop
		}
	}
	return nil
}

func validateComponentIsAcyclic(task *Task, visited *map[string]NodeVisitedType) (loop *[]string) {
	log.WithField("task", task.Name).Infoln("DFS in")
	nodeType, found := (*visited)[task.Name]
	if !found {
		(*visited)[task.Name] = NodeInCurrentPath
	} else {
		switch nodeType {
		case NodeVisited:
			return nil
		case NodeInCurrentPath:
			loop := make([]string, 1)
			loop[0] = task.Name
			return &loop
		}
	}

	for _, dep := range task.PendingDependencies {
		loop := validateComponentIsAcyclic(dep, visited)
		if loop != nil {
			log.WithField("task", task.Name).Infoln("DFS found loop")
			if len(*loop) < 2 || (*loop)[0] != (*loop)[len(*loop)-1] {
				*loop = append(*loop, task.Name)
			}
			return loop
		}
	}

	(*visited)[task.Name] = NodeVisited
	log.WithField("task", task.Name).Infoln("DFS out")
	return nil
}

func (r *Repo) GetJob(id JobId) (*Job, error) {
	job := new(Job)
	res := r.db.Preload(clause.Associations).Take(job, "id = ", id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return job, nil
}

func (r *Repo) ProcessReadyToRunTasks(cb func([]Task) error) error {
	batchSize := 1024
	batch := make([]Task, batchSize)

	res := r.db.Model(&Task{}).Preload(clause.Associations).Joins("LEFT JOIN deps ON tasks.id = deps.task_id").Where("tasks.status = ? AND deps.task_id IS NULL", JobCreated).FindInBatches(&batch, len(batch), func(tx *gorm.DB, batchId int) error {
		err := cb(batch)
		if err != nil {
			log.WithError(err).Error("Batch task process failed")
			return err
		}
		return nil
	})

	return res.Error
}

func (r *Repo) GetAvailableAgents() (*[]Agent, error) {
	agents := make([]Agent, 0)

	res := r.db.Model(&Agent{}).Preload(clause.Associations).Joins("LEFT JOIN task_executions ON agents.id = task_executions.agent_id").Where("task_executions.agent_id IS NULL", JobCreated).Find(&agents)
	if res.Error != nil {
		return nil, res.Error
	}

	return &agents, nil
}
