package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		// app.logger.Print(err)
		// http.Error(
		// 	w,
		// 	"The server encountered a problem and could not process your request",
		// 	http.StatusInternalServerError,
		// )
	}
}
