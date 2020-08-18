# part2

Kör
```
go get -u github.com/johansundell/small-golang-course
```
För att få ner de nya uppateringarna i detta projekt


Denna del så kommer du att skapa en egen webserver som kommer att köra på port 8080 på din dator. Navigera med terminalen till part2 av detta projekt och kör
```
go build
```
detta kommer att kompilera en binär fil som heter part2 (part2.exe på win), starta den och besök http://localhost:8080 för att se vår skapelse ;)
Om din port 8080 är upptagen så byt ut raden i main.go som sätter vilken port den kommer att jobba på, de två ställen ni kan besöka kan ni se i main.go filen.

Lek gärna vidare med fler slutpunkter för websidor, nästa del så kommer vi att börja med templates för dessa websidor, och hur vi kan göra detta mera dynamiskt.

Som ni kan se så har golang ett enkelt sätt att sätta upp en webserver.
