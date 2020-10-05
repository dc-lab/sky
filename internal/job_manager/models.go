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
	JobUnspecified JobStatus = iota
	JobCreated
	JobWaiting
	JobScheduled
	JobRunning
	JobFinished
	JobFailed
	JobCancelled
)

type Job struct {
	ID        JobId `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Status JobStatus

	Tasks []Task
	Name  string
}

type Task struct {
	ID        TaskId `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Status              JobStatus
	JobID               JobId    `gorm:"type:uuid"`
	PendingDependencies []*Task  `gorm:"many2many:deps"`
	Executors           []*Agent `gorm:"many2many:task_executions"`

	Name         string
	Command      pq.StringArray `gorm:"type:text[]"`
	Files        pq.StringArray `gorm:"type:text[]"`
	FilePaths    pq.StringArray `gorm:"type:text[]"`
	Dependencies pq.StringArray `gorm:"type:text[]"`
}

type Agent struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AssignedTasks []*Task `gorm:"many2many:task_executions"`
}

type TaskExecution struct {
	AgentID string
	TaskID  TaskId `gorm:"type:uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Status JobStatus
}
