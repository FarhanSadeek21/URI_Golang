package main
import ("fmt")
func main() {
	var number float32
	fmt.Scanf("%f", &number)
	if number >= 0 && number <= 25 {
	fmt.Println("Intervalo [0,25]")
	} else if number > 25 && number <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if number > 50 && number <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else if number > 75 && number <= 100 {
		fmt.Println("Intervalo (75,100]")
	} else {
		fmt.Println("Fora de intervalo")
	}
}