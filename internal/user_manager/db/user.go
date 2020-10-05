package db

import (
	"context"
	"github.com/dc-lab/sky/internal/user_manager/app"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (user *User) Validate() (string, bool) {
	if user.Login == "" {
		return "Name is empty", false
	}
	if user.Password == "" {
		return "Password is empty", false
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id FROM users WHERE login=$1;", user.Login)
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	defer rows.Close()
	if rows.Next() {
		return "Login is busy", false
	}
	return "", true
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func generateHashFromPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func checkPassword(hash, raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return !(err != nil && err == bcrypt.ErrMismatchedHashAndPassword)
}

func generateNewToken() string {
	return randString(40)
}

func (user *User) Create() (string, bool) {
	if resp, ok := user.Validate(); !ok {
		return resp, ok
	}

	user.Password = generateHashFromPassword(user.Password)
	user.Id = uuid.New().String()
	user.Token = generateNewToken()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO users(id, login, password, token) VALUES ($1, $2, $3, $4)", user.Id, user.Login, user.Password, user.Token)
	if err != nil {
		log.Println(err)
		return "Something went wrong with db", false
	}
	user.Password = ""
	return "", true
}

func (user *User) Get() (string, bool) {
	login, password := user.Login, user.Password

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, login, password, token FROM users WHERE login=$1", login)
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}
	if !rows.Next() {
		log.Println("User not found")
		return "No such user", false
	}
	err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Token)
	if err != nil {
		log.Println(err)
		return "Internal error", false
	}

	if !checkPassword(user.Password, password) {
		log.Println("Wrong password")
		return "Invalid login credentials. Please try again", false
	}

	// Ok, that should mean that we have found user
	user.Password = ""
	return "", true
}

func ChangePassword(userId, oldPassword, newPassword string) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT password FROM users WHERE id=$1", userId)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return &app.UserNotFound{}
	}

	var hashedOldPassword string
	err = rows.Scan(&hashedOldPassword)
	if err != nil {
		return err
	}
	rows.Close()

	if !checkPassword(hashedOldPassword, oldPassword) {
		return &app.WrongPassword{}
	}

	hashedNewPassword := generateHashFromPassword(newPassword)
	newToken := generateNewToken()
	_, err = conn.Exec(context.Background(), "UPDATE users SET password = $1, token = $2 WHERE id = $3", hashedNewPassword, newToken, userId)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
