package common

import "fmt"

const (
	HTTP_SUCCESS        = 200
	HTTP_INTERNAL_ERROR = 500
)

func IsError(err error) bool {
	if err != nil {
		fmt.Println("Error : ", err)
		return true
	}

	return false
}
