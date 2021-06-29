// Package osInfo provides OS specific information (currently only it's name and version)
package osInfo

import (
	"fmt"
	"runtime"
	"strings"
)

// info is an interface defining which information to provide
//
type info interface {
	name() (string, error)
	version() (string, error)
	fullInfo() (string, error)
}

// used if the systems os name, version, etc are not known
//
const (
	UNKNOWN_NAME        = "UNKNOWN NAME"
	UNKNOWN_VERSION     = "UNKNOWN VERSION"
	UNKNOWN_COMBINATION = "UNKNOWN DETAILS"
)

//
var osInfo = newInfo()

// Name returns the os' name
//
func Name() (string, error) {
	return osInfo.name()
}

// Version returns the os' version
//
func Version() (string, error) {
	return osInfo.version()
}

// FullInfo returns the os' name and version combined and may some additional information about the os
//
func FullInfo() (string, error) {
	return osInfo.fullInfo()
}

// newInfo retrieves information based on the detected operating system
//
// returns an object compliant with the info interface
//
func newInfo() info {
	switch runtime.GOOS {
	case "darwin":
		return &macOsInfo{}
	case "linux":
		return &linuxInfo{}
	case "windows":
		return &windowsInfo{}
	default:
		return &unsupported{}
	}
}

// unsupported struct is a placeholder in case the detected os is not supported
//
type unsupported struct {
}

// fullInfo
//
func (d unsupported) fullInfo() (string, error) {
	return "UNSUPPORTED OS", nil
}

// name returns UNSUPPORTED OS and no error
//
func (d unsupported) name() (string, error) {
	return "UNSUPPORTED OS", nil
}

// version returns UNSUPPORTED OS and no error
//
func (d unsupported) version() (string, error) {
	return "UNSUPPORTED OS", nil
}

// stringCleanerFunc defines a function signature for changing a strings content
//
type stringCleanerFunc func(s string) string

// keyValueParser separates the passed data based on the pairSeparator and the attempts to parse the keys and their
// corresponding values using the keyValueSeparator.
// cf1 and cf2 are called on the parsed key and value respectively to remove any unwanted characters from the key or
// value strings.
//
// returns a map containing the key value pairs and an error if no key value pairs where found
//
func keyValueParser(
	data string, pairSeparator string, keyValueSeparator string, cf1 stringCleanerFunc, cf2 stringCleanerFunc,
) (
	out map[string]string, err error,
) {
	if strings.Contains(data, pairSeparator) {
		out = make(map[string]string)
		pairs := strings.Split(data, pairSeparator)
		for _, pair := range pairs {
			if strings.Contains(pair, keyValueSeparator) {
				parts := strings.Split(pair, keyValueSeparator)
				if len(parts) == 2 {
					key := cf1(parts[0])
					value := cf2(parts[1])
					out[key] = value
				}
			}
		}
	}
	if len(out) == 0 {
		err = fmt.Errorf("no key value pairs found")
	}

	return
}
