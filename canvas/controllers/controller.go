package controllers

import (
	"../models"
)

/*
Status - structure for server status
*/
type Status struct {
	Status  string
	Name    string
	Version string
}

/*
Controller - structure for controller
*/
type Controller struct {
	Name    string
	Logger  *models.Logger
	Session *models.Connection
}
