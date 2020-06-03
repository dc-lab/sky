package db

import (
	"context"
	"github.com/dc-lab/sky/user_manager/app"
	"github.com/google/uuid"
)

type Group struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users,omitempty"`
}

func (group *Group) Validate() error {
	if group.Name == "" {
		return &app.EmptyField{}
	}
	return nil
}

func (group *Group) HasUser(userId string) bool {
	for _, user := range group.Users {
		if user == userId {
			return true
		}
	}
	return false
}

func AddUsersToGroup(groupId string, users []string) error {
	if users == nil {
		return nil
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, user := range users {
		_, err = conn.Exec(context.Background(), "INSERT INTO user_group_relations (group_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;", groupId, user)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveUsersFromGroup(groupId string, users []string) error {
	if users == nil {
		return nil
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, user := range users {
		_, err = conn.Exec(context.Background(), "DELETE FROM user_group_relations WHERE group_id = $1 AND user_id = $2;", groupId, user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (group *Group) Create(userId string) error {
	if err := group.Validate(); err != nil {
		return err
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	group.Id = uuid.New().String()
	if !group.HasUser(userId) {
		group.Users = append(group.Users, userId)
	}
	_, err = conn.Exec(context.Background(), "INSERT INTO groups(id, name) VALUES ($1, $2);", group.Id, group.Name)
	if err != nil {
		return err
	}
	if err := AddUsersToGroup(group.Id, group.Users); err != nil {
		return err
	}
	return nil
}

func GetGroups(userId string) ([]Group, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, name FROM groups JOIN user_group_relations ON id = group_id WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	var groups = []Group{}
	for rows.Next() {
		group := Group{}
		err = rows.Scan(&group.Id, &group.Name)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	rows.Close()

	for _, group := range groups {
		rows, err = conn.Query(context.Background(), "SELECT user_id FROM user_group_relations WHERE group_id = $1", group.Id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var user string
			err = rows.Scan(&user)
			if err != nil {
				return nil, err
			}
			group.Users = append(group.Users, user)
		}
		rows.Close()
	}

	return groups, nil
}

func GetGroup(userId, groupId string) (*Group, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, name FROM groups WHERE id = $1", groupId)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, &app.GroupNotFound{}
	}
	var group Group
	err = rows.Scan(&group.Id, &group.Name)
	if err != nil {
		return nil, err
	}
	rows.Close()

	rows, err = conn.Query(context.Background(), "SELECT user_id FROM user_group_relations WHERE group_id = $1", groupId)
	if err != nil {
		return nil, err
	}
	var userBelongs bool
	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		if user == userId {
			userBelongs = true
		}
		group.Users = append(group.Users, user)
	}
	if !userBelongs {
		return nil, &app.PermissionDenied{}
	}
	return &group, nil
}

func (group *Group) Delete() error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(context.Background(), "DELETE FROM groups WHERE id=$1", group.Id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return &app.GroupNotFound{}
	}
	return nil
}

func (group *Group) Modify(usersToAdd, usersToRemove []string) error {
	if err := AddUsersToGroup(group.Id, usersToAdd); err != nil {
		return err
	}
	if err := RemoveUsersFromGroup(group.Id, usersToRemove); err != nil {
		return err
	}
	return nil
}