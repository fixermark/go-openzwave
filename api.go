package openzwave

//
// Provides a facade for the C++ API that exposes just enough of the underlying C++
// API to be useful to implementing the Ninja Zwave driver
//
// The functions in this module are responsible for marshalling to and from the C functions
// declared in api.hpp and api.cpp.
//

//
// The following #cgo directives assumed that 'go' is a symbolic link that references the gopath that contains the current directory, e.g. ../../../..
//
// All 'go' packages that have this package as a dependency should include such a go link and will then inherit the library built in this package.
//

// #cgo LDFLAGS: -lopenzwave -Lgo/src/github.com/ninjasphere/go-openzwave/openzwave
// #cgo CPPFLAGS: -Iopenzwave/cpp/src/platform -Iopenzwave/cpp/src -Iopenzwave/cpp/src/value_classes
// #include "api.h"
import "C"

// A value-less type that is used to represent signals generated by the API, particularly quit signals used to
// ask an EventLoop to quit.

type Signal struct{}

type api struct {
	loop          EventLoop
	callback      Callback
	device        string
	quitEventLoop chan Signal
	logger        Logger
	networks      map[uint32]*network
}

//
// The API interface is available to implementors of the EventLoop type when the
// Configurator.Run() method is called.
//
type API interface {
	// The EventLoop should return from the function when a signal is received on this channel
	QuitSignal() chan Signal

	// the API logger
	Logger() Logger
}

func (self *api) QuitSignal() chan Signal {
	return self.quitEventLoop
}

func (self *api) Logger() Logger {
	return self.logger
}

func (self *api) getNetwork(homeId uint32) *network {
	net, ok := self.networks[homeId]
	if !ok {
		net = newNetwork(homeId)
		self.networks[homeId] = net
	}
	return net
}
