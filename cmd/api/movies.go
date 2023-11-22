package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "create a new movie")
	if err != nil {
		return
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}
	_, err = fmt.Fprintf(w, "show the details of movie %d\n", id)
	if err != nil {
		return
	}
}
