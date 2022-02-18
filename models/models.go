package models

import (
	"time"
)

type User struct {
	//	UserId   int    `db:"userid" json:"userId"`
	Name     string `db:"name" json:"Name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	//CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type LoginUser struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type FogetPassword struct {
	Email    string `db:"email" json:"email"`
	Userid   int    `db:"userid" json:"userid"`
	Password string `db:"password" json:"password"`
}

type Todo struct {
	Task   string    `db:"task" json:"task"`
	Detail string    `db:"detail" json:"detail"`
	Date   time.Time `db:"date" json:"date"`
}

type ShowTodo struct {
	ID            int       `db:"id" json:"ID"`
	UserId        int       `db:"userid" json:"userId"`
	Task          string    `db:"task" json:"task"`
	Detail        string    `db:"detail" json:"detail"`
	Date          time.Time `db:"date" json:"date"`
	CreatedAt     time.Time `db:"createdat" json:"createdAt"`
	TaskCompleted string    `db:"task_completed" json:"task_completed"`
	//ArchivedAt    time.Time `db:"archived_at" json:"archived_at"`
}

type UpTodo struct {
	ID        int       `db:"id" json:"ID"`
	UserId    int       `db:"userid" json:"userId"`
	Task      string    `db:"task" json:"task"`
	Detail    string    `db:"detail" json:"detail"`
	Date      time.Time `db:"date" json:"date"`
	CreatedAt time.Time `db:"createdat" json:"createdAt"`
}

type Deletetodo struct {
	ID int `db:"id" json:"ID"`
	//UserId    int       `db:"userid" json:"userId"`

}

type Complete struct {
	ID            int       `db:"id" json:"ID"`
	UserId        int       `db:"userid" json:"userId"`
	Task          string    `db:"task" json:"task"`
	Detail        string    `db:"detail" json:"detail"`
	Date          time.Time `db:"date" json:"date"`
	CreatedAt     time.Time `db:"createdat" json:"createdAt"`
	TaskCompleted string    `db:"task_completed" json:"task_completed"`
}
