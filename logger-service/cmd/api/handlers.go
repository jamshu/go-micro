package main

import (
	"logger/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func(app *Config) WriteLog(w http.ResponseWriter, r *http.Request){
	var requestPayload JSONPayload
	err := app.readJSON(w,r,&requestPayload)
	if err != nil {
		app.errorJSON(w,err)
		return
	}
	//insert into mongo database
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w,err)
		return
	}
	resp := jsonResponse{
		Error: false,
		Message: "entry logged successfully",
	}
	app.writeJSON(w,http.StatusAccepted,resp)
}