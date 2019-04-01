package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
Index - return if server is live
*/
func (c *Controller) Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := c.Session.Connect()
	if err != nil {
		error := models.RespError{Error: "Failed to connect, cannot reach database"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		error := models.RespError{Error: "Failed to connect, cannot reach database"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(Status{Status: "OK", Name: "Canvas.API.Status", Version: "1.0.0"})
	return
}
