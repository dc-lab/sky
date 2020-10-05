package server

import (
	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/data_manager/master/modeldb"
)

func FileProtoToDb(file *pb.File) *modeldb.File {
	return &modeldb.File{
		Id:         file.Id,
		Name:       file.Name,
		TaskId:     file.TaskId,
		Executable: file.IsExecutable,
		Tags:       file.Tags,
	}
}

func FileDbToProto(file *modeldb.File) *pb.File {
	return &pb.File{
		Id:           file.Id,
		Name:         file.Name,
		TaskId:       file.TaskId,
		IsExecutable: file.Executable,
		Tags:         file.Tags,
		UploadUrls:   file.UploadUrls,
	}
}
