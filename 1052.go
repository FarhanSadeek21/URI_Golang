package main
import (
	"fmt"
)
func main(){
	var month int
	fmt.Scanf("%d", &month)
	monthArray := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(monthArray[month - 1])
}