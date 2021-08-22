package main
import ("fmt")
func main() {
	var a, b, c int32
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a > b {
		if a > c {
			fmt.Printf("%d eh o maior\n", a)
		} else {
			fmt.Printf("%d eh o maior\n", c)
		}
	} else {
		if b > c {
			fmt.Printf("%d eh o maior\n", b)
		} else {
			fmt.Printf("%d eh o maior\n", c)
		}
	}
}