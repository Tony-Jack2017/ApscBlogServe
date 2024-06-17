package model

import (
	"ApscBlog/common/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func AddUser(user *User) error {
	_, err := userConn.InsertOne(context.TODO(), user)
	return err
}
func SearchUser(user *User) (error, *User) {
	var res User
	data, err := bson.Marshal(user)
	if err != nil {
		return err, nil
	}
	err = userConn.FindOne(context.TODO(), data).Decode(&res)
	if err != nil {
		return err, nil
	}
	return nil, &res
}
func UpdateUser(user *User) error {
	filter := bson.M{"user_id": user.UserID}
	_, err := userConn.UpdateOne(context.TODO(), filter, user)
	return err
}
func GetUserList(user *User) (error, *[]User) {
	var res *[]User
	data, err := bson.Marshal(user)
	if err != nil {
		return err, nil
	}
	list, err := userConn.Find(context.TODO(), data)
	if err != nil {
		return err, nil
	}
	err = list.Decode(&res)
	if err != nil {
		return err, nil
	}
	return nil, res
}
