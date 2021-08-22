package main
import ("fmt")
func main(){
	var a, b, higher, lower int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	sum := 0
	if a > b {
		higher = a
		lower = b
	} else {
		higher = b
		lower = a
	}
	for i := lower; i <= higher; i++ {
		if i % 13 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}