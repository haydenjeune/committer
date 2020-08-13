package errors

import (
	"fmt"
	"os"
)

// PrintAndExit prints the error message contained in err and exits with code 1.
func PrintAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
