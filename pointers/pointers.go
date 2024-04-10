package pointers

import "fmt"

func Pointers() {
	v := "caminar"
	p := &v
	fmt.Println("address -->", p)
	fmt.Println("value -->", *p)
	*p = "soñar"
	fmt.Println("value -->", *p)
}
