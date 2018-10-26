package main
import "fmt"

func main(){
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] =2

	fmt.Println(m)
	fmt.Println(len(m))

	delete(m,"k1")

	fmt.Println(m)

	x,y := m["k2"]

	fmt.Println("x:",x)
	fmt.Println("y:",y)

	_, z := m["k2"]

	fmt.Println(z)
}
