package main

// Vi behöver lite flera bibliotek nu ;)
import (
	"fmt"
	"log"
	"net/http"
)

// Fortfarande där vi startar allt
func main() {
	// Knyter upp alla anrop till vår funktion "defaultHandler"
	http.HandleFunc("/", defaultHandler)
	// lägger till en till funktion som kommer att svara på test
	http.HandleFunc("/test", justAnExampleJSONHandler)

	/*
		Så länge man ger en funktion som har samma signatur så
		kan man också göra det inline, läs mer i funktionen
		defaultHandler) vad de olika sakerna är
	*/
	// och en annan på exempel
	http.HandleFunc("/exempel", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "bara ett exempel")
	})
	/*
		Logga fel och stoppa om vi inte kan starta vår webserver
		på port 8080 på alla dina ipadresser, testa med att besöka
		http://localhost:8080 för att se din webserver
		http://localhost:8080/test ger en annan sida
		http://localhost:8080/exempel har en till
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	/*
		w variablen vi får i detta anrop är vad vi kan skriva till
		r variablen vi får i detta anrop är vad som skickas till oss
		fmt.Fprint funktionen kan skriva till en io.Writer som variabeln
		w implementerar med samma funktioner och då kan användas på samma sätt;
		men mer av detta senare :)
	*/
	fmt.Fprint(w, "<html><body>Hello World ;)</body></html>")
}

/*
	Vi skapar en till webhandler som ger ett svar tillbaka
*/
func justAnExampleJSONHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>The worlds worst web server ;)</body></html>")
}
