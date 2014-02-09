package main

import (
	libRig "github.com/jasonmm/go-rig/librig"
	"fmt"
)

func main() {
	libRig.NameGender = libRig.MALE
	libRig.PhoneHasX = false
	ident := libRig.GetIdentity()
	fmt.Println(ident)
}
