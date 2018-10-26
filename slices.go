package main
import "fmt"

func main(){
	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "3")
	s = append(s, "d", "e")

	fmt.Println(s)
	fmt.Println("Size of slice is", len(s))

	l := s[2:5]

	fmt.Println(l)

	c := make([]string, 2)
	copy(c,s)
	fmt.Println(c)
}
