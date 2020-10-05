package job_manager

import (
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

	err = db.AutoMigrate(&Job{}, &Task{})
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
	log.WithField("job_name", job.Name).Infoln("Add job")

	job.Status = JobCreated
	taskByName := make(map[string]*Task)
	for _, task := range job.Tasks {
		taskByName[task.Name] = &task
		task.Status = JobCreated
	}
	for _, task := range job.Tasks {
		task.Status = JobCreated
	}

	res := r.db.Create(job)
	if err := res.Error; err != nil {
		return err
	}

	log.WithField("job_id", job.ID).Infoln("created job")
	for _, task := range job.Tasks {
		log.WithFields(log.Fields{
			"task_id": task.ID,
			"job_id":  task.JobID,
		}).Infoln("created task")
	}

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

	res := r.db.Preload("task_pending_dependencies").Joins("LEFT JOIN task_pending_dependencies ON tasks.id = task_pending_dependencies.task_id").Where("tasks.status = ? AND task_pending_dependencies.task_id IS NULL", JobCreated).FindInBatches(&batch, len(batch), func(tx *gorm.DB, batchId int) error {
		err := cb(batch)
		if err != nil {
			log.WithError(err).Error("Batch task process failed")
			return err
		}
		return nil
	})

	return res.Error
}
