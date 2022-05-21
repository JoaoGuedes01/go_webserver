package main

import (
	"fmt"
	"log"
	"net/http"
)


func postHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!= nil{
		fmt.Fprintf(w,"Parseform() err: %v",err)
		return
	}

	fmt.Fprintf(w, "POST request successful \n")
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Name: %s \n", name)
	fmt.Fprintf(w, "Password: %s \n", password)
}


func getHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET"{
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)

	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)

	fmt.Print("Starting the Web Server at port 3000")
	if err := http.ListenAndServe(":3000",nil); err != nil{
		log.Fatal(err)
	}
}
