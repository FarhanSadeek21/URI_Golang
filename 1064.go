package main
import ("fmt")
func main() {
	var number, total, average float32
	count := 0
	for i := 1; i <= 6; i++ {
		fmt.Scanf("%f", &number)
		if number >= 0 {
			count += 1
			total += number
		}
	}
	average = total / float32 (count)
	fmt.Printf("%d valores positivos\n", count)
	fmt.Printf("%.1f\n", average)
}