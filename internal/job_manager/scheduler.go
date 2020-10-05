package job_manager

import (
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type scheduler struct {
	repo       *Repo
	dispatcher *dispatcher

	startedJobs  chan JobId
	updatedTasks chan TaskId
}

func NewScheduler(repo *Repo, dispatcher *dispatcher) (*scheduler, error) {
	return &scheduler{
		repo:         repo,
		dispatcher:   dispatcher,
		startedJobs:  make(chan JobId),
		updatedTasks: make(chan TaskId),
	}, nil
}

func (s *scheduler) UpdateTask(task TaskId, timeout time.Duration) error {
	timer := time.After(timeout)
	select {
	case <-timer:
		log.WithField("timeout", timeout).Errorln("Failed to handle task update: is main scheduler goroutine failed?")
		return errors.New("Failed to handle task update")
	case s.updatedTasks <- task:
		return nil
	}
}

func (s *scheduler) StartJob(job JobId, timeout time.Duration) error {
	timer := time.After(timeout)
	select {
	case <-timer:
		log.WithField("timeout", timeout).Errorln("Failed to schedule job start: is main scheduler goroutine failed?")
		return errors.New("Failed to schedule job start")
	case s.startedJobs <- job:
		return nil
	}
}

func (s *scheduler) Run(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()
	err := s.run()
	log.WithError(err).Errorf("scheduler failed")
	return err
}

func (s *scheduler) run() error {
	ticker := time.Tick(time.Second * 100000) // FIXME(BigRedEye): Read interval from config
	for {
		select {
		case task := <-s.updatedTasks:
			if err := s.onUpdatedTask(task); err != nil {
				log.WithError(err).Error("scheduler updated task handler failed")
				return fmt.Errorf("scheduler updated task handler failed: %w", err)
			}
		case job := <-s.startedJobs:
			if err := s.onStartedJob(job); err != nil {
				log.WithError(err).Error("scheduler start job handler failed")
				return fmt.Errorf("scheduler start job handler failed: %w", err)
			}
		case <-ticker:
			if err := s.onTick(); err != nil {
				log.WithError(err).Error("scheduler iteration failed")
				return fmt.Errorf("scheduler iteration failed: %w", err)
			}
		}
	}
}

// FIXME(BigRedEye): Do we need some custom logic here?
func (s *scheduler) onUpdatedTask(_ TaskId) error {
	return s.scheduleTasks()
}

func (s *scheduler) onStartedJob(_ JobId) error {
	return s.scheduleTasks()
}

func (s *scheduler) onTick() error {
	return s.scheduleTasks()
}

func (s *scheduler) scheduleTasks() error {
	return s.repo.Transaction(s.scheduleTasksImpl)
}

// TODO(BigRedEye): Implement fair share
func (s *scheduler) scheduleTasksImpl(tx *Repo) error {
	// TODO(BigRedEye): Filter agents by task requirements
	agents, err := tx.GetAvailableAgents()
	if err != nil {
		log.WithError(err).Error("Failed to select available agents")
		return err
	}

	tasksToRun := make([]Task, 0)
	err = tx.ProcessReadyToRunTasks(func(batch []Task) error {
		for i := range batch {
			log.WithField("task_name", batch[i].Name).WithField("task_id", batch[i].ID.String()).Infoln("Task is ready to run")
			tasksToRun = append(tasksToRun, batch...)
		}
		return nil
	})
	if err != nil {
		log.WithError(err).Error("Failed to select ready to run tasks")
		return err
	}

	var nextAvailableAgent = 0
	for i := range tasksToRun {
		if nextAvailableAgent >= len(*agents) {
			break
		}

		s.dispatcher.PushTask(assignedTask{
			resource: (*agents)[nextAvailableAgent].ID,
			task:     &tasksToRun[i],
		})

		nextAvailableAgent++
	}

	return nil
}
