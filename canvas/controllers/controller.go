package controllers

import (
	"../models"
)

/*
Status - strucutre or server status
*/
type Status struct {
	Status  string
	Name    string
	Version string
}

/*
Controller - structure to make multiple controllers if needed
*/
type Controller struct {
	Name    string
	Logger  *models.Logger
	Session *models.Connection
}

/*
Login - structure to hold temp logins
*/
type Login struct {
	Username string
	Password string
}
