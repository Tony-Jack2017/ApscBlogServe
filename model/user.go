package model

import (
	"ApscBlog/common/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id;unique"`
	Account  string             `json:"account" bson:"account,select(login|check)"`
	Password string             `json:"password" bson:"password,select(login)"`
	FullName string             `json:"fullName" bson:"fullName"`
	Username string             `json:"username" bson:"username"`
	Words    string             `json:"words" bson:"words"`
	Avatar   string             `json:"avatar" bson:"avatar"`
	Cover    string             `json:"cover" bson:"cover"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone"`
	model.BaseTime
}

func AddUser(user *User) (bool, error) {
	_, err := userConn.InsertOne(context.TODO(), user)
	if err != nil {
		return false, err
	}
	return true, nil
}
func UpdateUser() {
}
func SearchUser() (error, *User) {
	return nil, nil
}
