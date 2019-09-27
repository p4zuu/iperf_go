package iperf_go

/*
#cgo LDFLAGS: -liperf
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <stdint.h>
#include <iperf_api.h>
#include <sys/socket.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// IperfTest is a little trick
// The C. functions returns a private variable
// So we created a new struct with public members so that the object are usable outside
type IperfTest struct {
	Ptr           *C.struct_iperf_test
	JsonString    string
	JsonStart     []byte `json:"start"`
	JsonIntervals []byte `json:"intervals"`
	JsonEnd       []byte `json:"end"`
}

type Protocol int

const (
	UDP = iota
	TCP
)

type IperfStream struct {
	Ptr *C.struct_iperf_stream
}

func NewIperfTest() *IperfTest {
	test := IperfTest{}
	test.Ptr = C.iperf_new_test()
	return &test
}

/*
func NewIperfStream(test *IperfTest, socket int, sender int) {
	C.iperf_new_stream(test.Ptr, C.int(socket), C.int(sender))
}
*/
func IperfDefaults(test *IperfTest) int {
	return int(C.iperf_defaults(test.Ptr))
}

func IperfInitStream(stream *IperfStream, test *IperfTest) int {
	return int(C.iperf_init_stream(stream.Ptr, test.Ptr))
}
func IperfAddStream(test *IperfTest, stream *IperfStream) {
	C.iperf_add_stream(test.Ptr, stream.Ptr)
}

func IperfSetVerbose(test *IperfTest, verbose int) {
	C.iperf_set_verbose(test.Ptr, C.int(verbose))
}

func IperfSetTestRole(test *IperfTest, role string) error {
	if role != "c" && role != "s" {
		return errors.New("wrong role value")
	}
	Crole := C.CString(role)
	defer C.free(unsafe.Pointer(Crole))
	C.iperf_set_test_role(test.Ptr, *Crole)
	return nil
}

func IperfSetTestServerHostname(test *IperfTest, hostname string) {
	Chostname := C.CString(hostname)
	C.iperf_set_test_server_hostname(test.Ptr, Chostname)
	C.free(unsafe.Pointer(Chostname))
}

func IperfSetTestServerPort(test *IperfTest, port int) {
	C.iperf_set_test_server_port(test.Ptr, C.int(port))
}

func IperfSetTestOmit(test *IperfTest, omit int) error {
	if omit < 1 {
		return errors.New("wrong omit value")
	}
	C.iperf_set_test_omit(test.Ptr, C.int(omit))
	return nil
}

// SetDomain et SetConnect_timeout n'existent pas dans l'API de Iperf. C'est probablement ces deux paramètres qui sont non initialisés
/*func IperfSetTestDomain(test *IperfTest, domain int) error {
	if domain == 4 || domain == 6 {
		test.Ptr.settings.domain = C.int(domain)
		return nil
	} else {
		return errors.New("Wrong IP domain : 4 or 6 supported")
	}
}

// Socket timeout in ms
func IperfSetTestConnectionTimeout(test *IperfTest, timeout int) error {
	if timeout < 0 {
		return errors.New("Timeout can't be negative")
	} else {
		test.Ptr.settings.connect_timeout = C.int(timeout)
		return nil
	}
}*/

func IperfSetTestDuration(test *IperfTest, duration int) error {
	if duration <= 0 {
		return errors.New("wrong duration")
	}
	C.iperf_set_test_duration(test.Ptr, C.int(duration))
	return nil
}

//Set num streams : parallel connections to server
func IperfSetTestNumStreams(test *IperfTest, n int) {
	C.iperf_set_test_num_streams(test.Ptr, C.int(n))
}

// Set bitrate in bit/sec
func IperfSetTestRate(test *IperfTest, rate uint) {
	C.iperf_set_test_rate(test.Ptr, C.ulonglong(rate))
}

// Set length of buffer for read/write in KB
func IperfSetTestBlksize(test *IperfTest, blksize int) {
	C.iperf_set_test_blksize(test.Ptr, C.int(blksize))
}

//Set protocol : C.Pudp or C.Ptcp or
func IperfSetTestProtocol(test *IperfTest, protocol Protocol) error {
	var proto C.int
	switch protocol {
	case UDP:
		proto = C.Pudp
	case TCP:
		proto = C.Ptcp
	default:
		return errors.New("Wrong protocol")
	}
	if C.set_protocol(test.Ptr, proto) < 0 {
		return errors.New("Unknown error while setting protocol")
	}
	return nil
}

func IperfSetTestJsonOutput(test *IperfTest, value int) {
	C.iperf_set_test_json_output(test.Ptr, C.int(value))
}

func IperfSetTestReporterInterval(test *IperfTest, interval float32) error {
	if interval < 1 {
		return errors.New("wrong reporter interval")

	}
	C.iperf_set_test_reporter_interval(test.Ptr, C.double(interval))
	return nil
}

func IperfSetTestStatsInterval(test *IperfTest, interval int) error {
	if interval < 1 {
		return errors.New("wrong stats interval")

	}
	C.iperf_set_test_stats_interval(test.Ptr, C.double(interval))
	return nil
}

// TODO: return the error value with the integer value
func IperfRunClient(test *IperfTest) int {
	return int(C.iperf_run_client(test.Ptr))
}

func IperfStrError(errno int) string {
	return C.GoString(C.iperf_strerror(C.int(errno)))
}

func IperfGetTestOutfile(test *IperfTest) {
	C.iperf_get_test_outfile(test.Ptr)
}

func IperfGetTestJsonOutputString(test *IperfTest) string {
	return C.GoString(C.iperf_get_test_json_output_string(test.Ptr))
}

func IperfGetTestJsonOutput(test *IperfTest) int {
	return int(C.iperf_get_test_json_output(test.Ptr))
}

func IperfFreeTest(test *IperfTest) {
	C.iperf_free_test(test.Ptr)
}

func IperfFreeStream(stream *IperfStream) {
	C.iperf_free_stream(stream.Ptr)
}

func SetProtocol(test *IperfTest, protID int) (result int) {
	return int(C.set_protocol(test.Ptr, C.int(protID)))
}
