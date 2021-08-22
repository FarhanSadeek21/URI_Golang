package main
import ("fmt")
func main() {
	var year, month, day int
	fmt.Scanf("%d", &day)
	year = day / 365
	month = (day % 365) / 30
	day = (day % 365) % 30
	fmt.Printf("%d ano(s)\n", year)
	fmt.Printf("%d mes(es)\n", month)
	fmt.Printf("%d dia(s)\n", day)
}