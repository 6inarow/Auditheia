package osInfo

import (
	"fmt"
	"os"
	"strings"
)

// linuxInfo struct provides a map of all information received by getLinuxInfo()
//
type linuxInfo struct {
	infos map[string]string
}

const (
	RELEASE_FILE_LINUX           = "/etc/lsb-release"
	IDENTIFIER_NAME_LINUX        = "DISTRIB_ID"
	IDENTIFIER_VERSION_LINUX     = "DISTRIB_RELEASE"
	IDENTIFIER_CODENAME_LINUX    = "DISTRIB_CODENAME"
	IDENTIFIER_DESCRIPTION_LINUX = "DISTRIB_DESCRIPTION"
)

// name provides the linux os name
//
// returns the system os name or UNKNOWN NAME
//
func (l *linuxInfo) name() (name string, err error) {
	if l.infos == nil {
		l.infos, err = getLinuxInfo()
		if err != nil {
			return UNKNOWN_NAME, err
		}
	}
	if value, ok := l.infos[IDENTIFIER_NAME_LINUX]; ok {
		return value, nil
	} else {
		return UNKNOWN_NAME, fmt.Errorf("host os name: identifier '%s' not found", IDENTIFIER_NAME_LINUX)
	}
}

// version provides the linux os version
//
// returns the systems os version or UNKNOWN VERSION
//
func (l *linuxInfo) version() (version string, err error) {
	if l.infos == nil {
		l.infos, err = getLinuxInfo()
		if err != nil {
			return UNKNOWN_VERSION, err
		}
	}
	if value, ok := l.infos[IDENTIFIER_VERSION_LINUX]; ok {
		return value, nil
	} else {
		return UNKNOWN_VERSION, fmt.Errorf("host os version: identifier '%s' not found", IDENTIFIER_VERSION_LINUX)
	}
}

// fullInfo provides the name and version of the linux os and some additional info
//
// returns the systems name, version and additional info or UNKNOWN DETAILS
//
func (l *linuxInfo) fullInfo() (combined string, err error) {
	if l.infos == nil {
		l.infos, err = getLinuxInfo()
		if err != nil {
			return UNKNOWN_COMBINATION, err
		}
	}
	if value, ok := l.infos[IDENTIFIER_DESCRIPTION_LINUX]; ok {
		return value, nil
	} else {
		return UNKNOWN_COMBINATION, fmt.Errorf("host os version: identifier '%s' not found", IDENTIFIER_DESCRIPTION_LINUX)
	}
}

// getLinuxInfo gets all linux os infos located in /etc/os-release
//
// returns a map with all parsed key value pairs of the read file and an error if the command to read the file fails
//
func getLinuxInfo() (infos map[string]string, err error) {
	data, err := os.ReadFile(RELEASE_FILE_LINUX)
	if err != nil {
		return nil, err
	}
	infos, err = keyValueParser(string(data), "\n", "=", func(s string) string {
		return s
	}, func(s string) string {
		return strings.ReplaceAll(s, "\"", "")
	})
	return
}
