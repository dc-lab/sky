package job_manager

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type scheduler struct {
	repo *Repo

	startedJobs  chan JobId
	updatedTasks chan TaskId
}

func NewScheduler(repo *Repo) (*scheduler, error) {
	return &scheduler{
		repo:         repo,
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

func (s *scheduler) Run() {
	err := s.run()
	log.WithError(err).Fatalln("scheduler failed")
}

func (s *scheduler) run() error {
	ticker := time.Tick(time.Second) // FIXME(BigRedEye): Read interval from config
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
	return s.repo.Transaction(func(tx *Repo) error {
		err := tx.ProcessReadyToRunTasks(func(batch []Task) error {
			return nil
		})
		if err != nil {
			log.WithError(err).Error("Failed to schedule tasks")
			return err
		}
		return nil
	})
}

/*
func (s *scheduler) scheduleJob(jobId JobId) error {
	log.WithField("job_id", jobId).Infoln("Start schedule job")

	return s.repo.Transaction(func(tx *Repo) error {
		job, err := tx.GetJob(jobId)
		if err != nil {
			return err
		}

		return nil
	})
}
*/
