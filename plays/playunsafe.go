package plays

import (
	"fmt"
	unsafeimport "unsafe"
)

// PlayUnSafe - Just to print out the Sizeof a pointer to an int64
// This varies from machine to machine or do I need to say "Arch" here?
func PlayUnSafe() {
	x := int64(5)
	fmt.Println(unsafeimport.Sizeof(x))
}
