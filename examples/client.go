package main

import (
	"errors"
	"fmt"
	"strconv"

	iperf "github.com/lthomasmp/iperf_go"
)

func main() {
	anIperftest := iperf.NewIperfTest()
	defer iperf.IperfFreeTest(anIperftest)
	if anIperftest.Ptr == nil {
		fmt.Println(errors.New("test object is nil"))
	}
	iperf.IperfDefaults(anIperftest)
	iperf.IperfSetVerbose(anIperftest, 1)
	if err := iperf.IperfSetTestRole(anIperftest, "c"); err != nil {
		fmt.Println(errors.New("error while setting role"))
	}
	iperf.IperfSetTestServerHostname(anIperftest, "127.0.0.1")
	iperf.IperfSetTestServerPort(anIperftest, 30001)
	// Setting json output to 1 is necessary to get the json of the result
	iperf.IperfSetTestJsonOutput(anIperftest, 1)
	if error := iperf.IperfRunClient(anIperftest); error < 0 {
		fmt.Println(errors.New("error while calling run_client function : " + strconv.Itoa(error)))
	}
}
