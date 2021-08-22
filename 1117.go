package main
import "fmt"
func main(){
	var num, total, average float64
	count := 0
	for count < 2{
		fmt.Scanf("%f", &num)
		if num > 10 || num < 0 {
			fmt.Println("nota invalida")
		} else {
			total += num
			count += 1
		}
	}
	average = total / 2
	fmt.Printf("media = %.2f\n", average)
}