package job_manager

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type JobStatus int32
type JobId = uuid.UUID
type TaskId = uuid.UUID

const (
	JobCreated JobStatus = iota
	JobScheduled
	JobRunning
	JobFinished
	JobFailed
	JobCancelled
)

type Job struct {
	ID        JobId `gorm:"primarykey,type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Status JobStatus

	Tasks []Task
	Name  string
}

type Task struct {
	ID        TaskId `gorm:"primarykey,type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Status              JobStatus
	JobID               JobId   `gorm:"type:uuid"`
	PendingDependencies []*Task `gorm:"many2many:task_pending_dependencies"`

	Name        string
	Command     pq.StringArray `gorm:"type:text[]"`
	Files       pq.StringArray `gorm:"type:text[]"`
	Dependecies pq.StringArray `gorm:"type:text[]"`
}
