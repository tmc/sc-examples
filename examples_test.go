package examples

import (
	"fmt"

	"github.com/tmc/sc"
)

func ExampleBasics() {
	m := sc.Machine{}
	fmt.Println(m.State)
	// output:
	// MACHINE_STATE_UNSPECIFIED
}
