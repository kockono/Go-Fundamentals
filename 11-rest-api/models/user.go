package models

import (
	"errors"

	"res-api.com/apis/db"
	"res-api.com/apis/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	// Statement for prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashpassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashpassword

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) FindByEmail() error {
	query := `SELECT id, email, password FROM users WHERE email = ?`
	// Statement for prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.Email)
	err = row.Scan(&u.ID, &u.Email, &u.Password)
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	// pasamos como parametro el email u.Email = ?
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return err
	}

	// Comparamos la contraseña que nos pasan con la contraseña hasheada
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("Crede")
	}

	return nil
}
