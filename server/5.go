package main

import (
	"fmt"
	"net"
)

func main() {

	// get available network interfaces for
	// this machine
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, i := range interfaces {

		fmt.Printf("Name : %v \n", i.Name)

		byNameInterface, err := net.InterfaceByName(i.Name)

		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println("Interface by Name : ", byNameInterface)

		addresses, err := byNameInterface.Addrs()

		for k, v := range addresses {

			fmt.Printf("Interface Address #%v : %v\n", k, v.String())
		}
		fmt.Println("------------------------------------")

	}
}