package osInfo

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"Auditheia/memory/constants"
)

// windowsInfo struct provides a map of all infos received received by getWindowsInfo()
//
type windowsInfo struct {
	infos map[string]string
}

const (
	IDENTIFIER_NAME_WINDOWS    = "Caption"
	IDENTIFIER_VERSION_WINDOWS = "Version"
)

// name provides the windows os name
//
// returns the system os name or UNKNOWN NAME
//
func (w *windowsInfo) name() (name string, err error) {
	if w.infos == nil {
		w.infos, err = getWindowsInfo()
		if err != nil {
			return UNKNOWN_NAME, err
		}
	}
	if value, ok := w.infos[IDENTIFIER_NAME_WINDOWS]; ok {
		return value, nil
	} else {
		return UNKNOWN_NAME, fmt.Errorf("host os version: identifier '%s' not found", IDENTIFIER_VERSION_WINDOWS)
	}
}

// version provides the windows os version
//
// returns the systems os version or UNKNOWN VERSION
//
func (w *windowsInfo) version() (version string, err error) {
	if w.infos == nil {
		w.infos, err = getWindowsInfo()
		if err != nil {
			return UNKNOWN_VERSION, err
		}
	}
	if value, ok := w.infos[IDENTIFIER_VERSION_WINDOWS]; ok {
		return value, nil
	} else {
		return UNKNOWN_VERSION, fmt.Errorf("host os version: identifier '%s' not found", IDENTIFIER_VERSION_WINDOWS)
	}
}

// fullInfo provides the name and version of the windows os
//
// returns the systems name and version or UNKNOWN DETAILS
//
func (w *windowsInfo) fullInfo() (string, error) {
	name, err1 := w.name()
	version, err2 := w.version()

	if err1 != nil && err2 != nil {
		return strings.Join([]string{name, version}, " "), nil
	} else {
		var err1message string
		if err1 != nil {
			err1message = err1.Error()
		}
		var err2message string
		if err2 != nil {
			err2message = err2.Error()
		}
		return UNKNOWN_COMBINATION, fmt.Errorf("%s", strings.Join([]string{err1message, err2message}, "; "))
	}
}

// getWindowsInfo gets all windows os infos provided by wmic
//
// returns a map with all parsed key value pairs of the commands output and an error if the command failed
//
func getWindowsInfo() (infos map[string]string, err error) {
	// this only works with english locale... who thought this was a bright idea?
	data, err := exec.Command("wmic", "OS", "get", "Version,Caption", "/format:list").CombinedOutput()
	if err != nil {
		return nil, err
	}
	infos, err = keyValueParser(string(data), "\r\n", "=", func(s string) string {
		return s
	}, func(s string) string {
		return regexp.MustCompile(constants.REGEX_NOT_ALPHANUMERICAL).ReplaceAllString(s, " ")
	})
	return
}
