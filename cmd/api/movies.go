package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.standvpmnt.github.com/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    movie := data.Movie{
        ID: id,
        CreatedAt: time.Now(),
        Title: "Casablanca",
        Runtime: 102,
        Genres: []string{"drama", "romance", "war"},
        Version: 1,
    }

    // fmt.Fprintf(w, "show the details of movie %d\n", id)
    err = app.writeJSON(w, http.StatusOK, movie, nil)
    if err != nil {
        app.logger.Print(err)
        http.Error(
            w,
            "The server encountered a problem and could not process your request",
            http.StatusInternalServerError,
        )
    }
}
