package handler

import (
	"encoding/json"
	"fmt"
	"github.com/remotetodo/database/helper"
	"github.com/remotetodo/middleware"
	"github.com/remotetodo/models"
	"net/http"
)

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var todo models.Todo

	userID := middleware.GetUserFromContext(request)
	fmt.Println(userID)
	err := json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(todo.Date)
	usertask, newerr := helper.Newtodo(userID, todo.Task, todo.Detail, todo.Date)

	if newerr != nil {
		fmt.Println("error is : +v", newerr)
	}
	errjson := json.NewEncoder(writer).Encode(usertask)
	if errjson != nil {
		fmt.Println("error is :=v", errjson)
	}

}

func Update(w http.ResponseWriter, r *http.Request) {

	var updatetodo models.Todo
	userID := middleware.GetUserFromContext(r)
	fmt.Println(userID)
	err := json.NewDecoder(r.Body).Decode(&updatetodo)
	if err != nil {
		panic(err)
	}
	newerr := helper.Updatetodo(userID, updatetodo.Task, updatetodo.Detail, updatetodo.Date)
	if newerr != nil {
		fmt.Println("error is : +v", newerr)
	}
	w.Write([]byte(fmt.Sprintf("updated successfully")))

}

func Showalltodo(w http.ResponseWriter, r *http.Request) {

	userID := middleware.GetUserFromContext(r)
	fmt.Println(userID)

	todolist, err := helper.Show(userID)
	if err != nil {
		fmt.Println("todolisterror")
		panic(err)
		return
	}
	err = json.NewEncoder(w).Encode(todolist)
	if err != nil {
		fmt.Println("json error")
		panic(err)
	}

}

func Upcoming(w http.ResponseWriter, r *http.Request) {

	userID := middleware.GetUserFromContext(r)
	fmt.Println(userID)

	show, helpererr := helper.Up(userID)

	if helpererr != nil {
		panic(helpererr)
		return
	}

	errr := json.NewEncoder(w).Encode(show)
	if errr != nil {
		panic(errr)
		return
	}
}

func Expired(w http.ResponseWriter, r *http.Request) {

	userID := middleware.GetUserFromContext(r)
	fmt.Println(userID)

	showexpired, helpererr := helper.Ex(userID)

	if helpererr != nil {
		panic(helpererr)
		return
	}

	errr := json.NewEncoder(w).Encode(showexpired)
	if errr != nil {
		panic(errr)
		return
	}
}

func Deletetodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Deletetodo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		panic(err)
	}

	userID := middleware.GetUserFromContext(r)
	fmt.Println(userID)

	err = helper.Delete(userID, todo.ID)

	if err != nil {
		panic(err)
		return
	}
	w.Write([]byte(fmt.Sprintf("successfully deleted")))

}

func Completed(w http.ResponseWriter, r *http.Request) {

	var status = true

	userID := middleware.GetUserFromContext(r)
	fmt.Println(userID)
	com, helpererr := helper.Complete(userID, status)

	if helpererr != nil {
		panic(helpererr)
		return
	}

	errr := json.NewEncoder(w).Encode(com)
	if errr != nil {
		panic(errr)
		return
	}

}
