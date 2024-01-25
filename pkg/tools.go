package pkg

import (
	"fmt"
	"os"
)

func HandleError(message string, err error) {
	if err != nil {
		fmt.Println("ERROR: invalid data format, ", message, " ", err)
		os.Exit(1)
	}

}
