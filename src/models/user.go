package models

import "go.uber.org/zap/zapcore"

type User struct {
	UUID     string `json:"id"`
	UserName string `json:"user_name"`
}

//MarshalLogObject maps the User struct to be used in logging messages
func (u User) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddString("uuid", u.UUID)
	e.AddString("user_name", u.UserName)

	return nil
}

type CreateUserPayload struct {
	UserName        string `json:"user_name"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
