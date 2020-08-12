# part1

För att du ska kunna kompilera detta program så behöver du installera go (https://golang.org/dl/) och git (https://git-scm.com/downloads) kommer cokså att behövas då vi kommer att använda det mycket, har du inget konto på github så skaffa ett gärna men kommer inte att behövas. Men det kommer att vara bra för att lära sig att göra en fork och dela tillbaka sina ändringar uppstream.

Efter du har installerat både go och git så öppna en terminal och kör
```
go get github.com/johansundell/small-golang-course
```
Du kommer att få ett litet fel efter det då jag har varit elak coh delat in detta i underkataloger för att ha olika delar i det som vi ha kommer i slutet (Egen byggd webserver som kommer att köra som en tjänst och hämta och lämna data till FileMaker).

Men detta kommando kommer att skapa en katalog i din hem-katalog som heter go och i den en underkatalog src/github.com/johansundell/small-golang-course/part1 som inhehåller källkoden för denna startup del för att se att allt som behövs är installerat. Har skrivit lite dokumentarer som beskriver de olika delarna i detta projekt.

För att kompilera denna main.go fil så navigera till den katalogen med din terminal och kör 
```
go build
```
Det kommer att skapa ett program som du kan köra (det programmet kommer att heta part1 eller part1.exe beroende på plattform) ;)

Alla fel/problem som du stöter på, sänd mig en chatt på slack så ska jag hjälpa till. Har vi detta fungerande så kan vi gå vidare :) 

Lycka till/Johan