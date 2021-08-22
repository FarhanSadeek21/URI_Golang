package main
import "fmt"
func main() {
	var hour, minute, second int
	fmt.Scanf("%d", &second)
	hour = second / 3600
	minute = (second % 3600) / 60
    second = (second % 3600) % 60
	fmt.Printf("%d:%d:%d\n", hour, minute, second)
}