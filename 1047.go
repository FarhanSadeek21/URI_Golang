package main
import "fmt"
func main(){
	var init_h, init_m, fin_h, fin_m, hour, minute int
	fmt.Scanf("%d %d %d %d", &init_h, &init_m, &fin_h, &fin_m)
	time := (fin_h * 60 + fin_m) - (init_h * 60 + init_m)
	if time <= 0{
		time += 1440
	}
	hour = time / 60
	minute = time % 60
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)", hour, minute)
}