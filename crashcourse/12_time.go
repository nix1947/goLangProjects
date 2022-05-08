// time is represented as the number of second since jan 1 1970 in UTC
// In 2038, int32 capacity exceeded so int64 is okay.

package main

import (
	"fmt"
	"time"
)

func main() {
	future := time.Unix(126227800800, 0)
	fmt.Println(future)
}


