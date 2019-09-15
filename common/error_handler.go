package common

import (
	"fmt"
	"os"
)

func ThrowError(err error) {
	fmt.Printf(err.Error())
	os.Exit(0)
}
