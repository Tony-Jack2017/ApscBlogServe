package base

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID         primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	Account        string             `json:"account" bson:"account,select(login|check)"`
	Password       string             `json:"password" bson:"password,select(login)"`
	FullName       string             `json:"fullName" bson:"fullName"`
	Username       string             `json:"username" bson:"username"`
	Words          string             `json:"words" bson:"words"`
	Avatar         string             `json:"avatar" bson:"avatar"`
	Cover          string             `json:"cover" bson:"cover"`
	Email          string             `json:"email" bson:"email"`
	Phone          string             `json:"phone" bson:"phone"`
	model.BaseTime `bson:"inline"`
}

func AddUser(user *User) error {
	user.UserID = primitive.NewObjectID()
	_, err := model2.UserConn.InsertOne(context.TODO(), user)
	return err
}
func SearchUser(filter interface{}) (error, *User) {
	var res User
	err := model2.UserConn.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return err, nil
	}
	return nil, &res
}
func UpdateUser(user *User) error {
	filter := bson.M{"user_id": user.UserID}
	_, err := model2.UserConn.UpdateOne(context.TODO(), filter, user)
	return err
}
func GetUserList(user *User) (error, *[]User) {
	var res *[]User
	data, err := bson.Marshal(user)
	if err != nil {
		return err, nil
	}
	list, err := model2.UserConn.Find(context.TODO(), data)
	if err != nil {
		return err, nil
	}
	err = list.Decode(&res)
	if err != nil {
		return err, nil
	}
	return nil, res
}
