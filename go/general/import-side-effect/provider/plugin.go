package cloudprovider

import (
	"fmt"
)

var test string

func RegisterProvider() {
	test = "anchnet"
}

func Print(){
	fmt.Printf("%v",test)
}
