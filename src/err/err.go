package err

import "fmt"

//ThrowInvalidHexLitteralError can be thrown when lexing a hexlitteral
func ThrowInvalidHexLitteralError() {
	fmt.Printf("Invalid hex litteral character used")
}
