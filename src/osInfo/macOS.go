package osInfo

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"Auditheia/memory/constants"
)

// macOsInfo struct provides a map of all information received by getMacOSInfo()
//
type macOsInfo struct {
	infos map[string]string
}

const (
	IDENTIFIER_NAME_MACOS    = "ProductName"
	IDENTIFIER_VERSION_MACOS = "ProductVersion"
)

// name provides the macOS name
//
// returns the system os name or UNKNOWN NAME
//
func (m *macOsInfo) name() (name string, err error) {
	if m.infos == nil {
		m.infos, err = getMacOSInfo()
		if err != nil {
			return UNKNOWN_NAME, err
		}
	}
	if value, ok := m.infos[IDENTIFIER_NAME_MACOS]; ok {
		return value, nil
	} else {
		return UNKNOWN_NAME, fmt.Errorf("host os name: identifier '%s' not found", IDENTIFIER_NAME_MACOS)
	}
}

// version provides the macOS version
//
// returns the system os version or UNKNOWN VERSION
//
func (m *macOsInfo) version() (name string, err error) {
	if m.infos == nil {
		m.infos, err = getMacOSInfo()
		if err != nil {
			return UNKNOWN_NAME, err
		}
	}
	if value, ok := m.infos[IDENTIFIER_VERSION_MACOS]; ok {
		return value, nil
	} else {
		return UNKNOWN_NAME, fmt.Errorf("host os name: identifier '%s' not found", IDENTIFIER_VERSION_MACOS)
	}
}

// fullInfo provides the name and version of the macOS
//
// returns the systems name and version or UNKNOWN DETAILS
//
func (m *macOsInfo) fullInfo() (combined string, err error) {
	name, err1 := m.name()
	version, err2 := m.version()

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

// getMacOSInfo parses all macOS infos provided by sw_vers
//
// returns a map with all parsed key value pairs of the read file and an error if the command to read the file fails
//
func getMacOSInfo() (infos map[string]string, err error) {
	data, err := exec.Command("sw_vers").CombinedOutput()
	if err != nil {
		return nil, err
	}
	infos, err = keyValueParser(string(data), "\n", ":", func(s string) string {
		return s
	}, func(s string) string {
		return strings.TrimSpace(regexp.MustCompile(constants.REGEX_NOT_ALPHANUMERICAL).ReplaceAllString(s, " "))
	})
	return
}
