package main
import "fmt"
func main(){
	var start int32
	fmt.Scanf("%d", &start)
	if start % 2 == 0 {
		start += 1
	}
	for i := start; i <= (start + 11); i += 2 {
		fmt.Println(i)
	}
}