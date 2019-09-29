package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	iperf "github.com/lthomasmp/iperf_go"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Wrong argument. Usage: ./ClientServerExample <your_port>")
	}

	Port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Conversion failed")
	}

	go func() {
		err := RunServer(Port)
		if err != nil {
			log.Fatalf("Failed to run server")
		}

	}()
	time.Sleep(3 * time.Second)
	anIperfTest, err := RunClient(Port)
	if err != nil {
		log.Fatalf("Failed to run the client")
	}
	fmt.Println(anIperfTest.JsonString)
}

func RunClient(Port int) (*iperf.IperfTest, error) {
	anIperftest := iperf.NewIperfTest()
	defer iperf.IperfFreeTest(anIperftest)

	if anIperftest.Ptr == nil {
		return anIperftest, errors.New("test object is nil")
	}
	iperf.IperfDefaults(anIperftest)
	iperf.IperfSetVerbose(anIperftest, 1)

	if err := iperf.IperfSetTestRole(anIperftest, "c"); err != nil {
		return anIperftest, errors.New("error while setting role")
	}
	iperf.IperfSetTestServerHostname(anIperftest, "127.0.0.1")
	iperf.IperfSetTestServerPort(anIperftest, Port)

	iperf.IperfSetTestOmit(anIperftest, 3)
	iperf.IperfSetTestDuration(anIperftest, 5)
	iperf.IperfSetTestReporterInterval(anIperftest, 1)

	// Setting json output to 1 is necessary to get the json of the result
	iperf.IperfSetTestJsonOutput(anIperftest, 1)

	if error := iperf.IperfRunClient(anIperftest); error < 0 {
		return anIperftest, errors.New("error while calling run_client function")
	}
	return anIperftest, nil
}

func RunServer(Port int) error {
	anIperftest := iperf.NewIperfTest()
	defer iperf.IperfFreeTest(anIperftest)

	if anIperftest.Ptr == nil {
		return errors.New("test object is nil")
	}
	iperf.IperfDefaults(anIperftest)
	iperf.IperfSetVerbose(anIperftest, 1)

	if err := iperf.IperfSetTestRole(anIperftest, "s"); err != nil {
		return errors.New("error while setting role")
	}
	iperf.IperfSetTestServerPort(anIperftest, Port)

	if error := iperf.IperfRunServer(anIperftest); error < 0 {
		return errors.New("error while calling run_client function")
	}

	return nil
}
