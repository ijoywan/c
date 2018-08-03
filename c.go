package c

import (
	"fmt"
	"github.com/ijoywan/d"
)

func CallC() {
	fmt.Println("call C: v1.0.0")
	fmt.Println("   --> call D:")
	fmt.Printf("\t")
	d.Hello()
	fmt.Println("   --> call D end")
}
