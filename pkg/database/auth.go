package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) LoginAuthorization(loginRequest *models.LoginRequest) (int, string, error) {
	request := fmt.Sprintf("SELECT id, status FROM public.users WHERE login='%s' AND password='%s';", loginRequest.Login, loginRequest.Password)

	row := d.dbDriver.QueryRow(request)
	var id int
	var status string
	err := row.Scan(&id, &status)

	return id, status, err
}

func (d *Database) RegisterUser(registerRequest *models.RegisterRequest) (int, error) {

	request := fmt.Sprintf("SELECT id FROM public.users WHERE login ='%s' OR email = '%s';", registerRequest.Login, registerRequest.Email)
	log.Println(request)
	row := d.dbDriver.QueryRow(request)
	var id int
	if err := row.Scan(&id); err == nil {
		log.Println(err)
		return -1, errors.New("The user already exists")
	}

	ctx := context.Background()
	tx, err := d.dbDriver.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	status := "observer" // обычный пользователь
	request_UserRegister := fmt.Sprintf("INSERT INTO public.users(reg_time, login, email, password, status) VALUES('%s', '%s', '%s', '%s', '%s');", registerRequest.Time, registerRequest.Login, registerRequest.Email, registerRequest.Password, status)

	_, err = tx.ExecContext(ctx, request_UserRegister)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	newUserIDrequest := fmt.Sprintf("SELECT id FROM public.users WHERE login ='%s';", registerRequest.Login)

	newUserRow := d.dbDriver.QueryRow(newUserIDrequest)
	var newUserID int
	err = newUserRow.Scan(&newUserID)
	return newUserID, err

}

func (d *Database) generateRandomString(size int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
