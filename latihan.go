package main
import "fmt"

func digitTerbesar(n int)int{
	var num, banding int
	banding = n%10
	for n > 0{
		num = n%10
		if banding < num {
			banding = n
		}
		n = n/10
	}
	return banding
}
func main(){
	var masukan int
	fmt.Scan(&masukan)
	fmt.Print(digitTerbesar(masukan))
}