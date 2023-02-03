package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)

	data := map[string]string {
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(
			w, 
			"The server encountered a problem and could not process your request",
			http.StatusInternalServerError,
		)
	}
	// js, err := json.Marshal(data)
	// if err != nil {
	// 	app.logger.Print(err)
	// 	http.Error(
	// 		w, 
	// 		"The server encountered a problem and could not process your request", 
	// 		http.StatusInternalServerError,
	// 	)
	// }
	
	// js = append(js, '\n')

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)
}
