package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"time"
)

type BaseTime struct {
	CreatedAt LocalTime `json:"created_at" bson:"created_at"`
	UpdatedAt LocalTime `json:"updated_at" bson:"updated_at"`
	DeletedAt LocalTime `json:"deleted_at" bson:"deleted_at"`
}

type LocalTime time.Time

// override MarshalJSON

func (t LocalTime) MarshalJSON() ([]byte, error) {
	localTime := time.Time(t)
	return []byte(fmt.Sprintf("\"%v\"", localTime.Format("2006-01-02 15:04:05"))), nil
}

// when storage the timestamp convert time string to sql timestamp

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	localTime := time.Time(t)
	if localTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return localTime, nil
}

// convert sql timestamp to string which is yyyy-mm-dd HH:MM:SS format

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can't convert %v to timestamp", v)
}

func (t LocalTime) ParseTime(format string, isUnix bool) string {
	localTime := time.Time(t)
	if isUnix {
		return fmt.Sprintf("%s", localTime.UnixNano())
	} else {
		return localTime.Format(judgeTimeFormat(format))
	}
}

func judgeTimeFormat(format string) string {
	return "2006-01-02 15:04:05"
}

// MarshalBSONValue mongodb是存储bson格式，因此需要实现序列化bsonvalue(这里不能实现MarshalBSON，MarshalBSON是处理Document的)，将时间转换成mongodb能识别的primitive.DateTime

func (t *LocalTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	targetTime := t.ParseTime("", false)
	return bson.MarshalValue(targetTime)
}

// UnmarshalBSONValue 实现bson反序列化，从mongodb中读取数据转换成time.Time格式，这里用到了bsoncore中的方法读取数据转换成datetime然后再转换成time.Time
func (t *LocalTime) UnmarshalBSONValue(t2 bsontype.Type, data []byte) error {
	v, _, valid := bsoncore.ReadValue(data, t2)
	if valid == false {
		return errors.New(fmt.Sprintf("%s, %s, %s", "读取数据失败:", t2, data))
	}
	if v.Type == bsontype.DateTime {
		*t = LocalTime(v.Time())
	}
	return nil
}
