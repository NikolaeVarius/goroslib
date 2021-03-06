
# goroslib

[![GoDoc](https://godoc.org/github.com/aler9/goroslib?status.svg)](https://godoc.org/github.com/aler9/goroslib)
[![Go Report Card](https://goreportcard.com/badge/github.com/aler9/goroslib)](https://goreportcard.com/report/github.com/aler9/goroslib)
[![Build Status](https://travis-ci.org/aler9/goroslib.svg?branch=master)](https://travis-ci.org/aler9/goroslib)

goroslib is a library in pure Go that allows to build clients (nodes) for the Robot Operating System (ROS).

The Robot Operating System (ROS) is a project that provides a protocol specification to make multiple programs communicate with each other over time, exchanging structured data through topics, services and parameters. It was conceived for linking sensors, algorithms and actuators in unmanned ground vehicles (UGVs) and robots, but it is not bounded to the robot world and can be used anywhere there's the need of building streams of data (for example in video processing).

The official project provides libraries for writing nodes in C++ and Python, but they require the download of over 1GB of data and work only through a cmake-based buildchain, that is computationally intensive and difficult to customize. This library allows to write lightweight nodes that can be built with the standard Go compiler, do not need any runtime library and have a size of some megabytes.

Features:
* Subscribe and publish to topics
* Call and provide services
* Get and set parameters
* Get infos about other nodes, topics, services
* Compilation of `.msg` files is not necessary, message definitions are deducted from code
* Standard messages are available in folder `msgs/`
* Examples provided for every feature, comprehensive test suite, continuous integration

The library provides its features by implementing in pure Go all the ROS protocols (xml-rpc, TCPROS) and APIs (Master API, Parameter Server API, Slave API).

## Installation

Go &ge; 1.12 is required, and modules must be enabled (there must be a `go.mod` file in your project folder, that can be created with the command `go mod init main`). To install the library, it is enough to write its name in the import section of the source files that will use it. Go will take care of downloading the needed files:
```go
import (
    "github.com/aler9/goroslib"
)
```

## Examples

* [subscriber](examples/subscriber.go)
* [subscriber-custom](examples/subscriber-custom.go)
* [publisher](examples/publisher.go)
* [publisher-custom](examples/publisher-custom.go)
* [serviceclient](examples/serviceclient.go)
* [serviceprovider](examples/serviceprovider.go)
* [param-set-get](examples/param-set-get.go)
* [cluster-info](examples/cluster-info.go)

## Message definitions

To define custom messages, the standard ROS C++/Python libraries require `.msg` files in this format:
```
bool field1
int32 field2
```

This library doesn't require any `.msg` file, it is enough to write Go structures in this format:
```go
import (
    "github.com/aler9/goroslib/msgs"
)

type MessageName struct {
    msgs.Package `ros:"my_package"`
    Field1 msgs.Bool
    Field2 msgs.Int32
}
```

Primitive field types (bool, int, etc) and standard messages are [listed in the documentation](https://godoc.org/github.com/aler9/goroslib/msgs).

A command-line utility is provided to convert existing `.msg` files into their equivalent Go structures:
```
go get github.com/aler9/goroslib/commands/msg-import
msg-import --rospackage=my_package mymessage.msg
```

## Documentation

https://godoc.org/github.com/aler9/goroslib

## Testing

If you want to hack the library and test the results, unit tests can be launched with:

```
make test
```

## Links

Protocol documentation (v1)
* https://wiki.ros.org/ROS/Technical%20Overview
* https://wiki.ros.org/Implementing%20Client%20Libraries
* https://wiki.ros.org/ROS/Master_API
* https://wiki.ros.org/ROS/Parameter%20Server%20API
* https://wiki.ros.org/ROS/Slave_API
* https://wiki.ros.org/ROS/Connection%20Header
* https://wiki.ros.org/ROS/TCPROS
* https://fossies.org/linux/wireshark/epan/dissectors/packet-tcpros.c

Messages
* https://github.com/ros/std_msgs
* https://github.com/ros/common_msgs

Other Go libraries
* (v1) https://github.com/akio/rosgo
* (v2) https://github.com/juaruipav/rclgo

Other non-Go libraries
* (v1) [cpp] https://github.com/ros/ros_comm/tree/melodic-devel/clients/roscpp/src/libros (https://docs.ros.org/melodic/api/roscpp/html/classros_1_1NodeHandle.html)
* (v1) [python] https://docs.ros.org/melodic/api/rosnode/html/
* (v1) [c] https://github.com/ros-industrial/cros
* (v2) [misc] https://fkromer.github.io/awesome-ros2/
