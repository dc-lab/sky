package job_manager

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/dc-lab/sky/api/proto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	repo       *Repo
	config     *Config
	dispatcher *dispatcher
	scheduler  *scheduler
}

func CreateServer(config *Config, repo *Repo) (*Server, error) {
	dispatcher, err := NewDispatcher(config)
	if err != nil {
		return nil, err
	}

	scheduler, err := NewScheduler(repo, dispatcher)
	if err != nil {
		return nil, err
	}

	return &Server{
		repo:       repo,
		config:     config,
		dispatcher: dispatcher,
		scheduler:  scheduler,
	}, nil
}

func (s *Server) StartJob(ctx context.Context, req *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	tasks := make([]Task, len(req.Tasks))
	for i := range req.Tasks {
		tasks[i] = Task{
			Name:         req.Tasks[i].GetName(),
			Command:      req.Tasks[i].GetCommand(),
			Files:        req.Tasks[i].GetFiles(),
			Dependencies: req.Tasks[i].GetDependencies(),
		}
		log.Println(tasks[i])
	}

	job := &Job{
		Name:  req.GetName(),
		Tasks: tasks,
	}

	err := s.repo.AddJob(job)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create job: %w", err)
	}

	err = s.scheduler.StartJob(job.ID, time.Duration(time.Second))
	if err != nil {
		log.WithError(err).Warn("Failed to explicitly push task to the scheduler")
	}

	return &pb.StartJobResponse{
		Id: job.ID.String(),
	}, nil
}

func (s *Server) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse job id")
	}

	job, err := s.repo.GetJob(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unknown job")
	}

	tasks := make([]*pb.TaskStatus, len(job.Tasks))
	for i, task := range job.Tasks {
		deps := make([]string, len(task.PendingDependencies))
		for j, dep := range task.PendingDependencies {
			deps[j] = dep.Name
		}

		tasks[i] = &pb.TaskStatus{
			Id: task.ID.String(),
			Spec: &pb.TaskSpec{
				Name:         task.Name,
				Command:      task.Command,
				Dependencies: task.Dependencies,
			},
			Status:              pb.JobStatus(task.Status),
			PendingDependencies: deps,
		}
	}

	res := &pb.GetJobResponse{
		Id:     id.String(),
		Status: pb.JobStatus(job.Status),
		Tasks:  tasks,
	}

	return res, nil
}

func (s *Server) Run() error {
	errs := make(chan error, 3)
	var wg sync.WaitGroup
	go func() {
		errs <- s.dispatcher.Run(&wg)
	}()

	go func() {
		errs <- s.scheduler.Run(&wg)
	}()

	go func() {
		errs <- s.run(&wg)
	}()

	defer wg.Wait()

	errBuilder := strings.Builder{}
	numErrors := 0
	for i := 0; i < 3; i++ {
		err := <-errs
		if err != nil {
			numErrors++
			errBuilder.WriteString("")
			errBuilder.WriteString(err.Error())
			errBuilder.WriteString(",")
		}
	}

	if numErrors > 0 {
		return errors.New(errBuilder.String())
	}

	return nil
}

func (s *Server) run(wg *sync.WaitGroup) error {
	srv := grpc.NewServer()
	pb.RegisterJobManagerServer(srv, s)
	reflection.Register(srv)

	log.Println("Starting grpc server at", s.config.GrpcBindAddress)
	lis, err := net.Listen("tcp", s.config.GrpcBindAddress)
	if err != nil {
		log.WithError(err).Fatalln("Failed to listen socket for grpc server")
	}
	err = srv.Serve(lis)

	if err != nil {
		log.WithError(err).Errorf("Grpc server failed")
	}

	return err
}
