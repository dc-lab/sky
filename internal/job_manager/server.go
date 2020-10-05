package job_manager

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/dc-lab/sky/api/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	repo      *Repo
	config    *Config
	scheduler *scheduler
}

func CreateServer(config *Config, repo *Repo) (*Server, error) {
	scheduler, err := NewScheduler(repo)
	if err != nil {
		return nil, err
	}
	return &Server{
		repo:      repo,
		config:    config,
		scheduler: scheduler,
	}, nil
}

func (s *Server) StartJob(ctx context.Context, req *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	tasks := make([]Task, len(req.Tasks))
	for i := range req.Tasks {
		tasks[i] = Task{
			Name:        req.Tasks[i].GetName(),
			Command:     req.Tasks[i].GetCommand(),
			Files:       req.Tasks[i].GetFiles(),
			Dependecies: req.Tasks[i].GetDependencies(),
		}
	}

	job := &Job{
		Name:  req.GetName(),
		Tasks: tasks,
	}

	err := s.repo.AddJob(job)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add job")
	}

	return &pb.StartJobResponse{
		Id: job.ID.String(),
	}, nil
}

func (s *Server) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}

func (s *Server) Run() error {
	go s.scheduler.Run()

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
