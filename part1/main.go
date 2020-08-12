/*
Då detta är en körbar fil vi vill ha så är det main vi
måste ha här, annars så är det biblioteket vi vill
döpa vårt projekt till
*/
package main

/*
Här lägger vi till de bibliotek vi använder i vår
fil, i detta fall fmt som nu vi har för att kunna
skriva till konsolen
*/
import "fmt"

/*
main funktion som är start punkten i ett program,
(go är lite konstig då man måste ha måsvingar
efter funktioner och if, for osv; men vi slipper
ha ; efter varje rad så gilla läget ;). Så detta
är ett språk där man inte får ha dem som man vill
typ
func main()
{
}
funkar inte alls, men testa ;)
*/
func main() {
	// Vi skriver ut en rad till konsolen
	fmt.Println("Hello World!! ;)")
}
