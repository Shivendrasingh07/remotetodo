package handler

import (
	"encoding/json"
	"fmt"
	"github.com/remotetodo/database/helper"
	"github.com/remotetodo/models"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("err")
	}

	userid, newerr := helper.NewUser(user.Name, user.Email, user.Password)

	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(userid)
	if err2 != nil {
		fmt.Println("err")
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {
	var Cred models.LoginUser
	err := json.NewDecoder(request.Body).Decode(&Cred)
	if err != nil {
		fmt.Println("err")
	}
	loginUser, loginErr := helper.Login(Cred.Email, Cred.Password)
	if loginErr != nil {
		fmt.Println("err")
	}

	err2 := json.NewEncoder(writer).Encode(loginUser)
	if err2 != nil {
		fmt.Println("err")
	}

}

func ResetPassword(w http.ResponseWriter, request *http.Request) {
	var cred models.FogetPassword
	err := json.NewDecoder(request.Body).Decode(&cred)
	if err != nil {
		fmt.Println(err)
	}
	NewErr := helper.ForgetPass(cred.Userid, cred.Email, cred.Password)
	if NewErr != nil {
		fmt.Println(NewErr)
	}
}
