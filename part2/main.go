package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Knyter upp alla anrop till vår funktion "defaultHandler"
	http.HandleFunc("/", defaultHandler)
	/*
		Logga fel och stoppa om vi inte kan starta vår webserver
		på port 8080, testa med att logga in på
		http://localhost:8080 för att se din webserver
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	/*
		w variablen vi får i detta anrop är vad vi kan skicka tillbaka
		r variablen vi får i detta anrop är vad som skickas till oss
		fmt.Fprint funktionen kan skriva till en io.Writer som variabeln
		w implementerar med samma funktioner och då kan användas på samma sätt;
		men mer av detta senare :)
	*/
	fmt.Fprint(w, "<html><body>Hello World ;)</body></html>")
}
