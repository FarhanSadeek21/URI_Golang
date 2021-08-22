package main
import "fmt"
func main() {
	var a, b int32
	fmt.Scanf("%d %d", &b, &a)
	var time int32 = a - b
	if time <= 0 {
		time += 24
	}
	fmt.Printf("O JOGO DUROU %d HORA(S)\n", time)

}