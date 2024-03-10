package main

import "fmt"

/*func main() {
	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])
}*/

/*func main() {
	x := make(map[int]int)
	x[1] = 20
	x[2] = 30
	fmt.Println(x[1], x[2])
}*/

func main() {
	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Litio"
	fmt.Println(elemento["Li"])
}
