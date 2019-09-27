package main

import (
	"errors"
	"strconv"

	iperf "github.com/lthomasmp/iperf_go"
)

func main() {
	anIperftest := NewIperfTest()
	defer IperfFreeTest(anIperftest)
	if anIperftest.Ptr == nil {
		return anIperftest, errors.New("test object is nil")
	}
	IperfDefaults(anIperftest)
	IperfSetVerbose(anIperftest, 1)
	if err := IperfSetTestRole(anIperftest, "c"); err != nil {
		return anIperftest, errors.New("error while setting role")
	}
	IperfSetTestDomain(anIperftest, 4)
	IperfSetTestConnectionTimeout(anIperftest, 2000)
	IperfSetTestServerHostname(anIperftest, address)
	IperfSetTestServerPort(anIperftest, port)
	// Setting json output to 1 is necessary to get the json of the result
	IperfSetTestJsonOutput(anIperftest, 1)
	if err := IperfSetTestOmit(anIperftest, omit); err != nil {
		return anIperftest, errors.New("error while setting omit")
	}
	if err := IperfSetTestDuration(anIperftest, duration); err != nil {
		return anIperftest, errors.New("error while setting duration")
	}
	if err := IperfSetTestReporterInterval(anIperftest, reporter); err != nil {
		return anIperftest, errors.New("error while setting reporter interval")
	}
	if err := IperfSetTestStatsInterval(anIperftest, stats); err != nil {
		return anIperftest, errors.New("error while setting stats interval")
	}
	IperfSetTestRate(anIperftest, rate)
	IperfSetTestBlksize(anIperftest, blksize)
	IperfSetTestNumStreams(anIperftest, streamNumber)
	IperfSetTestProtocol(anIperftest, protocol)
	if error := IperfRunClient(anIperftest); error < 0 {
		return anIperftest, errors.New("error while calling run_client function : " + strconv.Itoa(error))
	}
	anIperftest.JsonString = IperfGetTestJsonOutputString(anIperftest)
	return anIperftest, nil
}
