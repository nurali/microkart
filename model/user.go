package model

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID     *uuid.UUID
	Name   *string
	Passwd *string
}

type Respository interface {
	Create(user *User) (*User, error)
	Load(username string) (*User, error)
}

type DbRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Respository {
	return &DbRepository{
		db: db,
	}
}

func (r *DbRepository) Create(user *User) (*User, error) {
	sql := "insert into users(name,passwd) VALUES($1,$2);"
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(user.Name, user.Passwd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *DbRepository) Load(username string) (*User, error) {
	sql := "select * from users where name=$1;"
	row := r.db.QueryRow(sql, username)

	var id *uuid.UUID
	var name, passwd *string
	err := row.Scan(&id, &name, &passwd)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:     id,
		Name:   name,
		Passwd: passwd,
	}, nil
}
