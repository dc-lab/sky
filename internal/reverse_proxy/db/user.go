package db

import "context"

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

func GetIdByToken(token string) (string, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return "", err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id FROM users WHERE token=$1", token)
	if err != nil {
		return "", err
	}
	if !rows.Next() {
		return "", &UserNotFoundError{}
	}

	var id string
	err = rows.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
