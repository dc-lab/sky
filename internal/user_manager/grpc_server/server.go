package grpc_server

import (
	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/user_manager/db"
)

type Server struct{}

func (s Server) GetUserGroups(user *pb.User, stream pb.UserManager_GetUserGroupsServer) error {
	userId := user.GetId()
	groups, err := db.GetGroups(userId)
	if err != nil {
		return err
	}
	for _, group := range groups {
		protoGroup := pb.Group{
			Id:    group.Id,
			Name:  group.Name,
			Users: group.Users,
		}
		if err = stream.Send(&protoGroup); err != nil {
			return err
		}
	}
	return nil
}
