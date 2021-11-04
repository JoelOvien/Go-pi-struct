package app

import (
	"postapi/app/database"

	"github.com/gorilla/mux"
)

//This App struct represents the structure of our app, it would contain two fields
//having a DB and a Router..the router is the gorilla/mux router
type App struct{
	Router *mux.Router
	//this field identifies in our app struct that we are making use of the 
	//PostDB in this app

	//This specifies what we said ahead about using two fields for our App struct...a DB and a Router field
	DB      database.PostDB
}

//Now we are gonna create a New() func that will be in charge of returning the actual application based on our struct
func New() *App{
	//define a to be of value...App which is our struct
	a:=  &App{
		Router: mux.NewRouter(),
	}
	
	a.initRoutes()
	return a

}

func (a *App) initRoutes(){
	a.Router.HandleFunc("/",a.indexHandler()).Methods("GET")
}