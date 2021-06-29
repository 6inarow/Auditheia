// Package configParser provides the functions necessary for parsing inputs from the commandline and the
// configuration file.
//
package configParser

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"

	"Auditheia/memory"
	"Auditheia/memory/constants"
	"Auditheia/output"

	"gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"
)

// ParseCLIFlags processes the arguments passed upon commandline call and writes their values to memory.OptionsList
// utilizing the flag package. Flags are assigned a default value.
//
// IMPORTANT: ParseCLIFlags must be called before using any Options from memory.OptionsList
//
func ParseCLIFlags() {
	// defining flag and default value for ConfFilePath
	flag.StringVar(&memory.OptionsList.BaseFolder, "output", constants.DEFAULT_BASEFOLDER, "define output folder name/path")
	// defining flag and default value for ConfFilePath
	flag.StringVar(&memory.OptionsList.ConfFilePath, "conf", constants.DEFAULT_CONF, "define configuration file name/path")
	// define whether the output format shall be JSON
	flag.BoolVar(&memory.OptionsList.ReportJSON, "json", false, "set the report format to JSON")
	// define whether the output format shall be YAML
	flag.BoolVar(&memory.OptionsList.ReportYaml, "yaml", false,
		"set the report format to YAML. If both JSON and YAML are selected, the report format will be JSON")
	// defining flag and default value for CheckOnly
	flag.BoolVar(&memory.OptionsList.CheckOnly, "check-only", false, "if set, the program will not perform the tasks of the audit but will only perform a sanity check")
	// defining flag and default value for Verbosity
	flag.IntVar(&memory.OptionsList.Verbosity, "verbosity", constants.VERBOSITY_WARNING,
		"define the level of logging:\n\t0 - Errors\n\t1 - Warning\n\t2 - Info\n\t3 - Debug\n\t4 - Trace\n")

	// parsing defined flags and assigning their values to correspondingly defined Options fields
	flag.Parse()

	memory.OptionsList.LogFileName = constants.DEFAULT_LOG
	memory.OptionsList.ReportFileName = constants.DEFAULT_REPORT

	cleanFlags()

}

// ParseConfigFile reads the data from the file specified in memory.OptionsList,
// unmarshalls it and maps the contents correspondingly to memory.MetaData,
// memory.AuditList and memory.AdditionalFiles.
//
// returns an error if os.ReadFile, json.Unmarshal or respectively yaml.Unmarshal fail
//
func ParseConfigFile() error {
	output.Log.Infof("Parsing Configuration from file: %s", memory.OptionsList.ConfFilePath)

	// reading file content
	output.Log.Infoln("Reading configuration file...")
	data, err := os.ReadFile(memory.OptionsList.ConfFilePath)
	if err != nil {
		return err // is that fine? -> fixed double error log if file not exists (read file and unmarshal)
	}

	// create temporary save for unmarshalled data
	tmpData := struct {
		// import memory.Meta fields
		memory.Meta
		AdditionalFiles *[]string       `json:"additional_files" yaml:"additional_files"`
		AuditList       *[]memory.Audit `json:"audit_list" yaml:"audit_list"`
	}{}

	if memory.CheckRegex(constants.REGEX_JSON_FILE_EXTENSION, path.Ext(memory.OptionsList.ConfFilePath)) { // extension json?
		output.Log.Infoln("Unmarshalling JSON...")
		// unmarshal JSON, write data into tmpData struct
		if err = json.Unmarshal(data, &tmpData); err != nil {
			return err
		}
	} else if memory.CheckRegex(constants.REGEX_YAML_FILE_EXTENSION, path.Ext(memory.OptionsList.ConfFilePath)) { // extension yaml?
		output.Log.Infoln("unmarshalling YAML")
		if err = yaml.Unmarshal(data, &tmpData); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unknown config format: '%s'", path.Ext(memory.OptionsList.ConfFilePath))
	}

	// assign Meta and AuditList to target fields in memory
	memory.MetaData = tmpData.Meta
	if tmpData.AuditList == nil {
		return fmt.Errorf("no audits have been found")
	}
	memory.AuditList = *tmpData.AuditList

	// map AdditionalFiles to memory
	memory.AdditionalFiles = map[string]bool{}
	if tmpData.AdditionalFiles != nil {
		for _, fileName := range *tmpData.AdditionalFiles {
			memory.AdditionalFiles[fileName] = true
		}
	}

	// check if configured os and runtime os match
	_ = checkCorrectRuntimeOS()

	// check if root permissions are granted if needed in conf
	_, err = checkCorrectRootPermissions()
	if err != nil {
		output.Log.WithFields(logrus.Fields{"error": err}).Errorln("Checking for correct privileges failed.")
	}

	return nil
}

// checkCorrectRootPermissions checks whether the program was called with admin privileges if the configuration file
// states it required. The permission check is dependent of the operating system,
// as unix systems have the root user with id '0' and windows using security identifiers.
// For windows, read permissions on the file '\\.\PHYSICALDRIVE0' are checked by opening the file.
// If said operation fails, the program assumes being called with insufficient rights.
//
// returns true/false if admin permissions are set as defined and an error if determining the current user fails
//
func checkCorrectRootPermissions() (bool, error) {
	// get userInfo info
	userInfo, err := user.Current()
	if err != nil {
		return false, err
	}

	// check whether privileged rights are required
	if memory.MetaData.RootRequired {
		switch runtime.GOOS {
		case "linux", "darwin":
			// on unix systems the standard root userInfo id is 0
			if userInfo.Uid != "0" {
				// log warning if non-root userInfo is detected
				output.Log.WithFields(logrus.Fields{
					"username": userInfo.Username,
					"user_id":  userInfo.Uid,
				}).Warningf("Root previleges required, but non-root User %s (ID %s) detected.", userInfo.Username, userInfo.Uid)
				return false, nil
			}
		case "windows":
			// file exists on nearly all windows systems
			file, err := os.Open("\\\\.\\PHYSICALDRIVE0")
			defer func() {
				if file != nil {
					_ = file.Close()
				}
			}()
			// check for file open error
			if err != nil {
				// log warning if opening of file failed, as this indicates insufficient privileges
				output.Log.WithFields(logrus.Fields{
					"username": userInfo.Username,
					"user_sid": userInfo.Uid,
				}).Warningf("Admin previleges required, but non-admin User %s (SID %s) detected.", userInfo.Username,
					userInfo.Uid)
				return false, nil
			}
		default:
			// print a warning if the detected OS is not supported for permission check, as no generally correct permission check behavior exists
			output.Log.Warningf("Cannot check userInfo privileges: os '%s' not supported", runtime.GOOS)
			return false, nil
		}
	}
	return true, nil
}

// checkCorrectRuntimeOS checks whether the current operating system is the same as defined in the configuration file.
//
// returns true if runtime.GOOS and 'conf_os' (configuration file) are identical
//
func checkCorrectRuntimeOS() bool {
	if runtime.GOOS != memory.MetaData.ConfOS {
		// log warning in case of os mismatch
		output.Log.WithFields(logrus.Fields{
			"detected_os": runtime.GOOS,
			"config_os":   memory.MetaData.ConfOS,
		}).Warningf("OS mismatch: '%s' configured, but '%s' detected", memory.MetaData.ConfOS, runtime.GOOS)
		return false
	}
	return true
}

// cleanFlags cleans the flag input of memory.OptionsList
// BaseFolder needs a trailing slash
// Verbosity has to be between 0 and 4 (inclusive)
//
func cleanFlags() {

	if memory.OptionsList.BaseFolder != constants.DEFAULT_BASEFOLDER {
		memory.OptionsList.BaseFolder = strings.Trim(memory.OptionsList.BaseFolder, "\"'")
	}

	if memory.OptionsList.Verbosity > 4 || memory.OptionsList.Verbosity < 0 {
		output.Log.Warningf("no such verbosity level: '%d'. falling back to default '1' (Warning)",
			memory.OptionsList.Verbosity)
		memory.OptionsList.Verbosity = 1
	}
}
