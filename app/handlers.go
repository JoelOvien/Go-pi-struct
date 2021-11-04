package app

import (
	"fmt"
	"net/http"
)

func (a *App) indexHandler() http.HandlerFunc	{
	return func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Welcome to my API")
	}
}