package db

import (
	"context"
	"math/rand"

	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/resource_manager/app"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Permissions struct {
	Users  []string `json:"users"`
	Groups []string `json:"groups"`
}

type State struct {
	Status string `json:"status"`
}

type Resource struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Owner       string      `json:"owner"`
	Token       string      `json:"token"`
	Permissions Permissions `json:"permissions"`
	State       State       `json:"state"`
}

func GetStringTypeByEnum(enumType pb.EResourceType) string {
	switch enumType {
	case pb.EResourceType_SINGLE:
		return "single"
	case pb.EResourceType_POOL:
		return "pool"
	case pb.EResourceType_CLOUD_INSTANCE:
		return "cloud_instance"
	case pb.EResourceType_CLOUD_TASK:
		return "cloud_task"
	default:
		return ""
	}
}

func GetPbTypeByString(enumType string) pb.EResourceType {
	switch enumType {
	case "single":
		return pb.EResourceType_SINGLE
	case "pool":
		return pb.EResourceType_POOL
	case "cloud_instance":
		return pb.EResourceType_CLOUD_INSTANCE
	case "cloud_task":
		return pb.EResourceType_CLOUD_TASK
	default:
		return pb.EResourceType_RESOURCE_UNKNOWN
	}
}

func (resource *Resource) Validate() (string, bool) {
	if resource.Name == "" {
		return "Name is empty", false
	}
	if resource.Type != "single" && resource.Type != "pool" {
		return "Wrong type value", false
	}
	return "", true
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func AddUsersToResource(resourceId string, users []string) error {
	if users == nil {
		return nil
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, user := range users {
		_, err = conn.Exec(context.Background(), "INSERT INTO ur_permissions (user_id, resource_id) VALUES ($1, $2)  ON CONFLICT DO NOTHING;", user, resourceId)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveUsersFromResource(resourceId string, users []string) error {
	if users == nil {
		return nil
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, user := range users {
		_, err = conn.Exec(context.Background(), "DELETE FROM ur_permissions WHERE user_id = $1 AND resource_id = $2;", user, resourceId)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddGroupsToResource(resourceId string, groups []string) error {
	if groups == nil {
		return nil
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, group := range groups {
		_, err = conn.Exec(context.Background(), "INSERT INTO gr_permissions (group_id, resource_id) VALUES ($1, $2)  ON CONFLICT DO NOTHING;", group, resourceId)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveGroupsFromResource(resourceId string, groups []string) error {
	if groups == nil {
		return nil
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, group := range groups {
		_, err = conn.Exec(context.Background(), "DELETE FROM gr_permissions WHERE group_id = $1 AND resource_id = $2;", group, resourceId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (resource *Resource) Modify(usersToAdd, usersToRemove, groupsToAdd, groupsToRemove []string) error {
	if err := AddUsersToResource(resource.Id, usersToAdd); err != nil {
		return err
	}
	if err := RemoveUsersFromResource(resource.Id, usersToRemove); err != nil {
		return err
	}
	if err := AddGroupsToResource(resource.Id, groupsToAdd); err != nil {
		return err
	}
	if err := RemoveGroupsFromResource(resource.Id, groupsToRemove); err != nil {
		return err
	}
	return nil
}

func (resource *Resource) CreateMe() error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "INSERT INTO resources(id, owner_id, name, type, token) VALUES ($1, $2, $3, $4, $5);", resource.Id, resource.Owner, resource.Name, resource.Type, resource.Token)
	if err != nil {
		return err
	}
	return nil
}

func (resource *Resource) Create(userId string) (string, bool) {
	if message, ok := resource.Validate(); !ok {
		return message, false
	}

	resource.Id = uuid.New().String()
	resource.Owner = userId
	resource.Permissions = Permissions{}
	resource.State = State{Status: "offline"}
	resource.Token = randString(40)
	err := resource.CreateMe()
	if err != nil {
		log.Println(err)
		return "Something went wrong with db", false
	}
	return "", true
}

func GetResource(userId, resourceId string) (string, *Resource) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return "Internal error", nil
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, owner_id, name, type, token FROM resources WHERE id=$1 AND owner_id=$2", resourceId, userId)
	if err != nil {
		log.Println(err)
		return "Internal error", nil
	}
	if !rows.Next() {
		log.Println("Resource not found")
		return "No resource with such id", nil
	}
	var resource = &Resource{}
	err = rows.Scan(&resource.Id, &resource.Owner, &resource.Name, &resource.Type, &resource.Token)
	if err != nil {
		log.Println(err)
		return "Internal error", nil
	}
	rows.Close()
	resource.State = State{Status: ConnectedAgents.GetResourceStatus(resourceId)}

	rows, err = conn.Query(context.Background(), "SELECT user_id FROM ur_permissions WHERE resource_id = $1", resourceId)
	if err != nil {
		log.Println(err)
		return "Internal error", nil
	}
	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			log.Println(err)
			return "Internal error", nil
		}
		resource.Permissions.Users = append(resource.Permissions.Users, user)
	}

	rows, err = conn.Query(context.Background(), "SELECT group_id FROM gr_permissions WHERE resource_id = $1", resourceId)
	if err != nil {
		log.Println(err)
		return "Internal error", nil
	}
	for rows.Next() {
		var group string
		err = rows.Scan(&group)
		if err != nil {
			log.Println(err)
			return "Internal error", nil
		}
		resource.Permissions.Groups = append(resource.Permissions.Groups, group)
	}

	return "", resource
}

func GetResources() (*[]Resource, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.WithError(err).Error("Internal error")
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, owner_id, name, type, token FROM resources")
	if err != nil {
		log.WithError(err).Error("Internal error")
		return nil, err
	}

	res := make([]Resource, 0)

	for rows.Next() {
		var resource = Resource{}
		err = rows.Scan(&resource.Id, &resource.Owner, &resource.Name, &resource.Type, &resource.Token)
		if err != nil {
			log.WithError(err).Println("Internal error")
			return nil, err
		}

		resource.State = State{Status: ConnectedAgents.GetResourceStatus(resource.Id)}
		res = append(res, resource)
	}

	rows.Close()

	return &res, nil
}

func DeleteResource(userId, resourceId string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM resources WHERE id=$1 AND owner_id=$2", resourceId, userId)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return &app.ResourceNotFound{}
	}
	return nil
}

func GetUserResources(userId string) (string, []Resource) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return "Internal error", nil
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, owner_id, name, type, token FROM resources WHERE owner_id=$1", userId)
	if err != nil {
		log.Println(err)
		return "Internal error", nil
	}

	var resources = []Resource{}
	for rows.Next() {
		resource := Resource{}
		err = rows.Scan(&resource.Id, &resource.Owner, &resource.Name, &resource.Type, &resource.Token)
		if err != nil {
			log.Println(err)
			return "Internal error", nil
		}
		resource.State = State{Status: ConnectedAgents.GetResourceStatus(resource.Id)}
		resources = append(resources, resource)
	}
	return "", resources
}

func GetResourceIdByToken(token string) (string, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return "", err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id FROM resources WHERE token=$1", token)
	if err != nil {
		return "", err
	}
	if !rows.Next() {
		return "", &app.ResourceNotFound{}
	}

	var id string
	err = rows.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
