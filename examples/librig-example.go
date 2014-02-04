package main

import (
	libRig "github.com/jasonmm/go-rig/librig"
	"fmt"
)

func main() {
//	libRig.NameGender = libRig.MALE
	ident := libRig.GetIdentity()
	fmt.Println(ident)
}
