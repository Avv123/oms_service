package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	OrderCollection *mongo.Collection
)

type User struct {
	ID       string   `json:"id" bson:"_id"`
	Username string   `json:"username" bson:"username"`
	Email    string   `json:"email" bson:"email"`
	Password string   `json:"password" bson:"password"`
	Orders   []string `json:"orders" bson:"orders"`
}

func GeneratefromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil

}

//{
//"id":"1",
//"username":"aryaman",
//"email":"arya2003@gmail.com",
//"password":"arya2003"
//
//}
