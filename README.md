# iperf_go: A Golang wrapper for iperf3 API

"
iperf is a tool for active measurements of the maximum achievable bandwidth on IP networks. It supports tuning of various parameters related to timing, protocols, and buffers. For each test it reports the measured throughput / bitrate, loss, and other parameters.
"

I do not develop Iperf.
For more information see the : https://software.es.net/iperf

## MAINTAINING

This is a poc I did a few years ago, I don't think I will improve it. The best for you is probably to fork. But you 
can still suggest improvements or fixes, I will review it :^)


## Requirements

You first need to install the iperf3 library available on the [ENSET Github](https://github.com/esnet/iperf). I advise you to build it as it is described in the repo. If you have not installed the library, the Go compiler will not find `#cgo LDFLAGS: -liperf`

## Installation

Just run

```console
go get github.com/lthomasmp/iperf_go
```

## Running 

To test the wrapper, just run the example. This example use your local IP address.

```console
cd examples
go build ClientServerExample.go
./ClientServerExample <your_port>
```
