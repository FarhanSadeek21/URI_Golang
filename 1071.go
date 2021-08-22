package main
import "fmt"
func main() {
	var n1, n2, higher, lower int
	var sum int
	sum = 0
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	if n1 > n2 {
		higher = n1
		lower = n2
	} else if n1 < n2 {
		lower = n1
		higher = n2
	}
	for i := (lower + 1) ; i < higher; i++ {
		if i % 2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
