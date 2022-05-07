package domain

import "reflect"

type UserId int64

func (userId UserId) ToInt64() int64 {
	ref := reflect.ValueOf(userId)
	if ref.Kind() != reflect.Int64 {
		return 0
	}
	return ref.Int()
}

type User struct {
	Id        UserId
	UserName  string
	FirstName string
	LastName  string
}

type UserRetrieve struct {
	User  *User
	Error error
}
