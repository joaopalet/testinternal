package function

import (
	"fmt"

	"testinternal/random"
)

func MyFunction() {
	random.Random()
	fmt.Printf("My function was executed")
}
