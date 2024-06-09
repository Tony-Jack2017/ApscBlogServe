package model

import (
	"ApscBlog/common/model"
	"context"
)

type User struct {
	UserID   string `json:"user_id" bson:"user_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Words    string `json:"words" bson:"words"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Cover    string `json:"cover" bson:"cover"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
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
