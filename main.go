package main

import (
	"fmt"

	. "amber/foundation"
)

func main() {	
	lst := NewListOf(1, 2, 3)

	if three := Find(lst, 3); three != nil {
		
	}

	fmt.Println("Hello")
}
