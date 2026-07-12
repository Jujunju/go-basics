package main

import "fmt"

type IPAddress [4]uint8

func main() {

	hosts := map[string]IPAddress{
		"loopBack":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v :%v\n", name, ip)
	}

}

func (i IPAddress) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}
