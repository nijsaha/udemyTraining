package main

import (
	"fmt"
	"log"
)

//CustomMsgError Error Struct Type
type CustomMsgError struct {
	lat, long string
	err       error
}

//implements the above struct.  Therefore implements the error interface. i.e. overwrites (Polymorphism)
func (c *CustomMsgError) Error() string {
	return fmt.Sprintf("custom error has occured: %v %v %v", c.lat, c.long, c.err)
}

func main() {
	x, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)

	} else {
		fmt.Println(x)
	}

}
func sqrt(n int) (float64, error) {
	if n < 0 {
		nme := fmt.Errorf("SQRT cannot be derived for negative numbers %v", n) //little bit more context)
		return 0, &CustomMsgError{"50.2289 N", "78.8727 W", nme}               //extra more context
	}
	return 42, nil
}
