package db

import (
	"context"
	"github.com/google/uuid"
	"log"
	"math/rand"
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

func (resource *Resource) Create(userId string) (string, bool) {
	if message, ok := resource.Validate(); !ok {
		return message, false
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	defer conn.Release()

	resource.Id = uuid.New().String()
	resource.Owner = userId
	resource.Permissions = Permissions{}
	resource.State = State{Status: "offline"}
	resource.Token = randString(40)
	_, err = conn.Exec(context.Background(), "INSERT INTO resources(id, owner_id, name, type, token) VALUES ($1, $2, $3, $4, $5);", resource.Id, resource.Owner, resource.Name, resource.Type, resource.Token)
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
	return "", resource
}

func DeleteResource(userId, resourceId string) (string, bool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM resources WHERE id=$1 AND owner_id=$2", resourceId, userId)
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	if ct.RowsAffected() == 0 {
		log.Println("Resource not found")
		return "Resource not found", false
	}
	return "", true
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
		resources = append(resources, resource)
	}
	return "", resources
}