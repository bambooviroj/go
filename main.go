package main

import (
	"fmt"
)

/*ID defination
 */
func ID(v bool) bool {
	return v
}

func main() {
	fmt.Printf("%v\n", ID(true))
}
