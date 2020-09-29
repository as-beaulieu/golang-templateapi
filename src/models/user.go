package models

import "go.uber.org/zap/zapcore"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//MarshalLogObject maps the User struct to be used in logging messages
func (u User) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddString("id", u.ID)
	e.AddString("name", u.Name)

	return nil
}
