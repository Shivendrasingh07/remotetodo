package helper

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/remotetodo/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//HashPassword helps to encrypt the password
func HashPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(pass)
}

//CheckHashPassword  verify the password of user
func CheckHashPassword(password, hashpass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

//NewUser save the data of new user
func NewUser(name, email, password string) (int, error) {
	SQL := `INSERT INTO users(name,email,password)VALUES($1,$2,$3) RETURNING userid`
	var userid int
	pass := HashPassword(password)
	err := database.Data.Get(&userid, SQL, name, email, pass)
	if err != nil {
		return 404, err
	}
	return userid, nil

}

//Login function takes email and password, returns userID
func Login(email, password string) (string, error) {
	SQL := `SELECT userid,password FROM users WHERE email=$1`
	var userid, Hashpass string
	err := database.Data.QueryRowx(SQL, email).Scan(&userid, &Hashpass)
	if err != nil {
		return "", err
	}
	pass, passErr := CheckHashPassword(password, Hashpass)
	if pass != true && passErr != nil {
		return "", passErr
	}
	//fmt.Println(userid)
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		Issuer:    userid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("token error : +v", err)
	}

	return ss, nil

}

func ForgetPass(userid int, email, password string) error {
	SQL := `SELECT userid FROM users WHERE email=$1`
	var id int
	err := database.Data.QueryRowx(SQL, email).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	if userid == id {
		newSQL := `UPDATE users SET password=$1 where userid=$2`
		pass := HashPassword(password)
		_, newErr := database.Data.Exec(newSQL, pass, userid)
		if newErr != nil {
			fmt.Println(newErr)
		}
	} else {
		return errors.New("WRONG CREDENTIALS")
	}

	return nil
}

/*
func GetUser(token string) (string, error) {
	SQL := `SELECT userid from session where token=$1`
	var user string
	err := database.Data.Get(&user, SQL, token)
	if err != nil {
		return "", err
	}
	return user, nil
}
*/
