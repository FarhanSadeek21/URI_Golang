package main
import ("fmt")
func main() {
	var hour, speed, consumption float64
	fmt.Scanf("%f", &hour)
	fmt.Scanf("%f", &speed)
	consumption = speed * hour
	fmt.Printf("%.3f\n") 
}