package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"postapi/app"
)

func main(){
	fmt.Println("Our go")
	app:=	app.New()

	http.HandleFunc("/",app.Router.ServeHTTP)
	err := http.ListenAndServe(":9000",nil)
	check(err)

}


func check(e error){
	if e !=nil{
		log.Println(e)
		os.Exit(1)
	}
}
