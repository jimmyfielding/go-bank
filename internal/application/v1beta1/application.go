package v1beta1

import (
	"time"
)

type Application struct {
	FirstName   string `validate:"required,alpha"`
	Surname     string
	DateOfBirth time.Time
	Address     Address
}

type Address struct {
	FirstLine  string
	SecondLine string
	City       string
	County     string
	PostCode   string
}

type ApplicationWrapper struct {
	Application *Application
}
