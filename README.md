# iperf_go: A Golang wrapper for iperf3 API

"
iperf is a tool for active measurements of the maximum achievable bandwidth on IP networks. It supports tuning of various parameters related to timing, protocols, and buffers. For each test it reports the measured throughput / bitrate, loss, and other parameters.
"

I do not develop Iperf.
For more information see the : https://software.es.net/iperf

## Requirements

You first need to install the iperf3 library available on the [ENSET Github](https://github.com/esnet/iperf). I advise you to build it as it is described in the repo. If you have not installed the library, the Go compiler will not find `#cgo LDFLAGS: -liperf`

## Installation

Just run

```console
$ go get github.com/lthomasmp/iperf_go
```

## Running 

To test the wrapper, just run the server and the client code in two different terminals.

```console
$ cd examples
$ go run server.go <your_port>
```

```console
$ cd examples
$ go run client.go 127.0.0.1 <your_port>
```

The wrapper and the examples are not finished yet, I will have finished them soon.
