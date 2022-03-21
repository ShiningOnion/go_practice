package main

import (
	"fmt"
	"go_practice/p0007_errors/2.custom-error/errorSample"
)

func checkUserNameExist(username string) (bool, error) {
	if username == "Leo" {
		return true, errorSample.ErrUserNameExist{UserName: username}
	}

	return false, nil
}

func main() {
	if _, err := checkUserNameExist("foo"); err == nil {
		fmt.Println("foo not exist")
	}

	if _, err := checkUserNameExist("Leo"); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := checkUserNameExist("Leo"); err != nil {
		if ok := errorSample.IsErrUserNameExist(err); ok {
			fmt.Println("user Leo already exist.")
		}
	}
}
