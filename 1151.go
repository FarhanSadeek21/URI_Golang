package main
import "fmt"
func main(){
    var limit, n3 int
	fmt.Scanf("%d", &limit)
    var n1, n2 int
	n1 = 0
	n2 = 0
    if limit == 1 {
		fmt.Println(n)
    } else if limit == 2 {
        //cout << n1 << " " << n2 << endl
		fmt.Printf("%d %d", n1, n2)
    }
    else {
        fmt.Printf("%d %d", n1, n2)
        limit -= 2
        for i := 1; i <= limit; i++ {
            n3 = n1 + n2
            //cout << n3
			fmt.Printf("%d", n3)
            if i == limit {
                //cout << endl
				fmt.Println()
            }  else {
                //cout << " "
				fmt.Printf(" ")
            }
            n1 = n2
            n2 = n3
        }
    }
    return 0
}