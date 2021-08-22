package main
import ("fmt")
func main(){
	var distance, fuel, consumption float64
	fmt.Scanf("%f", &distance)
	fmt.Scanf("%f", &fuel)
	consumption = distance / fuel
	fmt.Printf("%.3f km/l\n", consumption)
}