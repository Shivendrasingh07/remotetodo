package helper

import (
	"database/sql"
	"github.com/remotetodo/database"
	"github.com/remotetodo/models"
	"time"
)

func Newtodo(userid int, task, detail string, date time.Time) (string, error) {

	SQL := `INSERT INTO todos(userid,task,detail,date)VALUES($1,$2,$3,$4) RETURNING task`
	var todo string

	err1 := database.Data.Get(&todo, SQL, userid, task, detail, date)
	if err1 != nil {
		return "newtodoerr", err1
	}

	return todo, nil

}

func Updatetodo(userid int, task, detail string, date time.Time) error {

	SQL := `update todos set task=$2,detail=$3,date=$4 where userid=$1 and archived_at isnull `

	_, err := database.Data.Exec(SQL, userid, task, detail, date)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func Show(userid int) ([]models.ShowTodo, error) {

	SQL := `SELECT * from todos WHERE userid=$1 and archived_at isnull `
	user := make([]models.ShowTodo, 0)
	err := database.Data.Select(&user, SQL, userid)
	if err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows {
		return nil, err
	}
	return user, nil
}

func Up(userid int) ([]models.UpTodo, error) {

	date := time.Now()

	SQL := `SELECT * from todos WHERE userid=$1 and date::date>=$2::date and archived_at isnull `
	todo := make([]models.UpTodo, 0)
	err := database.Data.Select(&todo, SQL, userid, date)
	if err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows {
		return nil, err
	}
	return todo, nil
}

func Ex(userid int) ([]models.UpTodo, error) {
	date := time.Now()

	SQL := `SELECT * from todos WHERE userid=$1 and date::date<$2::date and archived_at isnull `
	todo := make([]models.UpTodo, 0)
	err := database.Data.Select(&todo, SQL, userid, date)
	if err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows {
		return nil, err
	}
	return todo, nil
}

func Delete(userid, id int) error {
	date := time.Now()
	SQL := `update todos set archived_at=date WHERE userid = $1 and id = $2 and archived_at isnull `

	_, err := database.Data.Exec(SQL, date, userid, id)

	if err != nil {
		return err
	}

	return nil
}

func Complete(userid int, status bool) ([]models.Complete, error) {

	SQL := `SELECT * from todos WHERE userid=$1 and task_completed=$2 and archived_at isnull `
	todo := make([]models.Complete, 0)
	err := database.Data.Select(&todo, SQL, userid, status)
	if err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows {
		return nil, err
	}
	return todo, nil
}
