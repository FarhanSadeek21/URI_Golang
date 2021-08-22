package main
import ("fmt")
func main() {
	var number float32
	count := 0
	for i := 1; i <= 6; i++ {
		fmt.Scanf("%f", &number)
		if number >= 0 {
			count += 1
		}
	}
	fmt.Printf("%d valores positivos\n", count)
}