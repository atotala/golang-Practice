package main
import "fmt"

func main(){
	nums := []int{4,5,6}

	for _, num := range nums{
		fmt.Println(num)
	}

	kvs := map[string]string{"a":"apple", "b":"battle"}

	for k,v := range kvs{
		fmt.Printf("%s -> %s \n",k,v)
	}

	for i,j := range "ANAND"{
		fmt.Println(i,j)
	}
}
