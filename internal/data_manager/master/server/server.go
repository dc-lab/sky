package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/data_manager/master/service"
)

type Server struct {
	Files *service.FilesService
}

func getUserId(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	userIds, ok := md["user-id"]
	if !ok || len(userIds) != 1 {
		return "", false
	}
	return userIds[0], true
}

func (s *Server) CreateFile(ctx context.Context, req *pb.CreateFileRequest) (*pb.CreateFileResponse, error) {
	userId, ok := getUserId(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to retrieve user id")
	}

	// FIXME
	if userId != "test-user" {
		return nil, status.Errorf(codes.Unauthenticated, "Unknown user")
	}

	if req.File == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to parse file info")
	}

	fileReq := FileProtoToDb(req.File)
	/// FIXME
	fileReq.Owner = userId

	file, err := s.Files.CreateFile(fileReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to create new file: %s", err)
	}

	return &pb.CreateFileResponse{File: FileDbToProto(file)}, nil
}

func (s *Server) GetFileInfo(ctx context.Context, req *pb.GetFileInfoRequest) (*pb.GetFileInfoResponse, error) {
	file, err := s.Files.GetFile(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to find file: %s", err)
	}

	return &pb.GetFileInfoResponse{File: FileDbToProto(file)}, nil
}

func (s *Server) GetFileLocation(ctx context.Context, req *pb.GetFileLocationRequest) (*pb.GetFileLocationResponse, error) {
	urls, err := s.Files.GetFileLocation(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to find file: %s", err)
	}

	return &pb.GetFileLocationResponse{DownloadUrls: urls}, nil
}

func (s *Server) GetTaskResults(ctx context.Context, req *pb.GetTaskResultsRequest) (*pb.GetTaskResultsResponse, error) {
	files, err := s.Files.ListTaskResults(req.TaskId, req.Path)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to list results: %s", err)
	}

	var res pb.GetTaskResultsResponse
	for _, file := range files {
		res.Files = append(res.Files, FileDbToProto(&file))
	}

	return &res, nil
}

func (s *Server) UpdateTaskResults(ctx context.Context, req *pb.UpdateTaskResultsRequest) (*pb.UpdateTaskResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskResults not implemented")
}

func (s *Server) ValidateUpload(ctx context.Context, req *pb.ValidateUploadRequest) (*pb.ValidateUploadResponse, error) {
	// FIXME
	if req.UserId != "test-user" {
		return nil, status.Errorf(codes.Unauthenticated, "Unknown user")
	}

	file, err := s.Files.GetFile(req.FileId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to find file: %s", err)
	}

	allow := false
	// FIXME Check ACL
	if req.UserId == file.Owner {
		allow = true
	}

	return &pb.ValidateUploadResponse{Allow: allow}, nil
}

func (s *Server) SubmitFileHash(ctx context.Context, req *pb.SubmitFileHashRequest) (*pb.SubmitFileHashResponse, error) {
	// FIXME
	if req.UserId != "test-user" {
		return nil, status.Errorf(codes.Unauthenticated, "Unknown user")
	}

	allow, err := s.Files.TrySetFileHash(req.UserId, req.NodeId, req.FileId, req.Hash)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unknown file")
	}

	return &pb.SubmitFileHashResponse{Allow: allow}, nil
}

func (s *Server) GetFileHash(ctx context.Context, req *pb.GetFileHashRequest) (*pb.GetFileHashResponse, error) {
	// FIXME
	if req.UserId != "test-user" {
		return nil, status.Errorf(codes.Unauthenticated, "Unknown user")
	}

	file, err := s.Files.GetFile(req.FileId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to find file: %s", err)
	}

	res := &pb.GetFileHashResponse{}
	res.Allow = (file.Owner == req.UserId)

	if res.Allow {
		res.Hash = file.Hash
	}

	return res, nil
}

func (s *Server) ResolveBlobReplicas(context.Context, *pb.ResolveBlobReplicasRequest) (*pb.ResolveBlobReplicasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveBlobReplicas not implemented")
}

func (s *Server) Loop(ctx context.Context, report *pb.NodeStatus) (*pb.NodeTarget, error) {
	targetBlobs, err := s.Files.ProcessNodeReport(report.NodeId, report.FreeSpace, report.BlobHashes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to list results: %s", err)
	}

	target := &pb.NodeTarget{BlobHashes: targetBlobs}

	return target, nil
}
