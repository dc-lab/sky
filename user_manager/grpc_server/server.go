package grpc_server

import (
	pb "github.com/dc-lab/sky/api/proto/user_manager"
	"github.com/dc-lab/sky/user_manager/db"
)

type Server struct {}

func (s Server) GetUserGroups(user *pb.TUser, stream pb.UserManager_GetUserGroupsServer) error {
	userId := user.GetId()
	groups, err := db.GetGroups(userId)
	if err != nil {
		return err
	}
	for _, group := range groups {
		protoGroup := pb.TGroup{
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
