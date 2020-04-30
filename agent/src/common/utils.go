package common

import "fmt"

func DealWithError(err error) {
	if err != nil {
		fmt.Println(err)
		// TODO(glebx777): add logging
	}
}
