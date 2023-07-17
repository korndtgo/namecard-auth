package model

import "time"

type ResetPassword struct {
	Id int `json:"id"`

	Uuid string `json:"uuid"`

	RefCode string `json:"refCode"`

	Otp string `json:"otp"`

	IsActive bool `json:"isActive"`

	ExpiredAt time.Time `json:"expiredAt"`

	CreatedAt time.Time `json:"createdAt"`
}
