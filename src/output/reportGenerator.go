package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"Auditheia/memory"
	"Auditheia/memory/constants"
	"Auditheia/osInfo"

	"gopkg.in/yaml.v2"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// report contains all information relevant for the report. Thus it contains a list of memory.AuditDTO objects, which
// only contain specific fields of memory.Audit objects. See memory.AuditDTO for details. The field reportText contains
// a marshalled version of report, which currently is generated as json.
//
type report struct {
	CustomerName    string            `json:"customer_name" yaml:"customer_name"`
	ConfigVersion   string            `json:"config_version" yaml:"config_version"`
	OperatingSystem string            `json:"operating_system" yaml:"operating_system"`
	HostName        string            `json:"host_name" yaml:"host_name"`
	StartTime       time.Time         `json:"start_time" yaml:"start_time"`
	GenerationTime  time.Time         `json:"generation_time" yaml:"generation_time"`
	ElapsedTime     string            `json:"elapsed_time" yaml:"elapsed_time"`
	ConfigFile      string            `json:"config_file" yaml:"config_file"`
	ReportBody      []memory.AuditDTO `json:"report_body" yaml:"report_body"`
	reportText      string
}

// GenerateReport returns a report-Object (Constructor). Elapsed time is set right before creating the object.
//
// returns a report pointer and an error (currently wont ever return errors,
// as all possible errors are handled inside the function)
//
func GenerateReport() *report {
	Log.Infoln("Generating report")
	// collect some basic info

	customerName := memory.MetaData.CustomerName
	confVersion := memory.MetaData.Version
	reportOS, err := osInfo.FullInfo()
	if err != nil {
		Log.Errorln(err)
	}
	reportHostName, err := os.Hostname()
	if err != nil {
		Log.Errorln(err)
	}
	reportStartTime := memory.StartTime
	reportEndTime := time.Now()
	reportConfFile := filepath.Base(memory.OptionsList.ConfFilePath)
	var reportBody []memory.AuditDTO

	// copy Audits and contained Tasks
	for _, audit := range memory.AuditList {
		var dAudit = memory.AuditDTO{}
		dAudit.GetFrom(audit)
		reportBody = append(reportBody, dAudit)
	}

	var elapsedTime = time.Now().Sub(memory.StartTime).String()

	return &report{
		CustomerName:    customerName,
		ConfigVersion:   confVersion,
		OperatingSystem: reportOS,
		HostName:        reportHostName,
		StartTime:       reportStartTime,
		GenerationTime:  reportEndTime,
		ElapsedTime:     elapsedTime,
		ConfigFile:      reportConfFile,
		ReportBody:      reportBody,
	}
}

// updateElapsedTime sets the duration between program start and calling of this function as r.ElapsedTime
//
func (r *report) updateElapsedTime() {
	r.ElapsedTime = time.Now().Sub(memory.StartTime).String()
}

// ToJSON serializes the content of report in JSON and saves it to r.reportText. The parameter indent allows to specify
// whether the JSON shall be in one line (indent = false) or with multiple lines and indentation (indent = true).
// ToJSON calls json.Marshal or json.MarshalIndent respectively.
//
// returns errors from json.Marshal or json.MarshalIndent
//
func (r *report) ToJSON(indent bool) (string, error) {
	Log.Infoln("Marshalling report as JSON...")
	var data []byte
	var err error
	if indent {
		data, err = json.MarshalIndent(r, "", "	")
	} else {
		data, err = json.Marshal(r)
	}
	if err != nil {
		return string(data), err
	}
	r.reportText = string(data)
	return r.reportText, nil
}

// ToYAML serializes the content of report in JSON and saves it to r.reportText.
// ToYAML calls yaml.Marshal respectively.
//
// returns errors from yaml.Marshal
//
func (r *report) ToYAML() (string, error) {
	Log.Infoln("Marshalling report as JSON...")
	data, err := yaml.Marshal(r)

	if err != nil {
		return string(data), err
	}
	r.reportText = string(data)
	return r.reportText, nil
}

// WriteToFile writes the content of r.reportText to the file specified in memory.OptionsList
// inside the programs base directory (memory.OptionsList.BaseFolder). If r.reportText
// is empty, ToJSON is called in an attempt to create a report text before writing it to the file system.
//
// Returns errors from ToJSON, os.Create or os.File.WriteString
//
func (r report) WriteToFile() error {
	Log.Infoln("Writing report to file...")
	// create report text if empty
	if r.reportEmpty() {
		Log.Warningln("report text is empty!")
		Log.Warningln("Trying to create report...")
		if memory.OptionsList.ReportYaml {
			_, err := r.ToYAML()
			if err != nil {
				return err
			}
		} else {
			_, err := r.ToJSON(true)
			if err != nil {
				return err
			}
		}

		if r.reportEmpty() {
			Log.Warningln("report text still empty, missing content")
			return errors.New("empty report body")
		}
	}
	// create output filename for report
	Log.Infoln("Creating report file...")
	var filename string
	if memory.OptionsList.ReportYaml {
		filename = path.Join(memory.OptionsList.BaseFolder,
			strings.TrimSuffix(memory.OptionsList.ReportFileName,
				path.Ext(memory.OptionsList.ReportFileName))+constants.SUFFIX_YAML)
	} else {
		filename = path.Join(memory.OptionsList.BaseFolder,
			strings.TrimSuffix(memory.OptionsList.ReportFileName, path.Ext(memory.OptionsList.ReportFileName))+constants.SUFFIX_JSON)
	}

	// create a file rotator with lumberjack
	fileRotator := &lumberjack.Logger{
		Filename:  filename,
		MaxAge:    30,
		LocalTime: true,
	}
	err := fileRotator.Rotate()
	if err != nil {
		return err
	}

	// write report text to file
	Log.Infoln("Writing content to file...")
	_, err = fileRotator.Write([]byte(r.reportText))
	if err != nil {
		return err
	}
	return nil
}

func (r *report) reportEmpty() bool {
	return r.reportText == "" || len(r.reportText) == 0
}

// SaveArtefacts is used to save generated Artefacts (from command outputs) and additional Files set in the config file.
// Saves additional files first. see saveAdditionalFiles for more info.
// searches memory.AuditList for Artefacts and attempts to save them as files in the programs output directory
// (memory.OptionsList.BaseFolder), which have the following Path:
//
//      artefacts/audit_<AUDITNUMBER>_<AUDITNAME>/task_<TASKNUMBER>/<ARTEFACTNAME>.
//
// E.g. an artefact with the Name "lsmod output" from the second memory.Task of the memory.Audit "cramfs filesystem
// check" will have the following path:
//
//      artefacts/audit_0_cramfs_filesystem_check/task_0/lsmod_output
//
// Errors returned are from os.MkdirAll, os.Create, WriteString or Close.
//
func SaveArtefacts() error {
	Log.Infoln("Saving Artefacts...")
	artefactDir := path.Join(memory.OptionsList.BaseFolder, constants.DEFAULT_ARTEFACTS_DIR)
	err := saveAdditionalFiles(artefactDir)
	if err != nil {
		Log.Errorln(err)
	}
	for n := range memory.AuditList {
		auditDir := generateName("audit", strconv.Itoa(n), memory.AuditList[n].Name)
		var taskList []*memory.Task
		for m := range memory.AuditList[n].Tasks {
			taskList = addTasks(taskList, &memory.AuditList[n].Tasks[m])
		}
		if taskList == nil || len(taskList) == 0 {
			Log.Warningf("Audit %d: No Tasks with Artefacts found!", n)
		}
		for m := range taskList {
			taskDir := generateName("task", strconv.Itoa(m))
			err := writeArtefactsToDirectory(taskList[m].GetArtefacts(), artefactDir, auditDir, taskDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// saveAdditionalFiles creates a directory for additional files and then iterates over the specified list of
// filenames in memory.AdditionalFiles, calling copyFile for each filename.
//
// returns an error if os.MkdirAll or copyFile fail
func saveAdditionalFiles(parentDir string) error {
	Log.Infoln("Saving additional files...")
	additionalDir := "additional_files"
	if memory.AdditionalFiles == nil || len(memory.AdditionalFiles) == 0 {
		return errors.New("no additional files to save")
	}
	err := os.MkdirAll(path.Join(parentDir, additionalDir), os.ModePerm)
	if err != nil {
		return err
	}
	for name := range memory.AdditionalFiles {
		err = copyFile(name, path.Join(parentDir, additionalDir, generateName(name)))
		if err != nil {
			Log.Errorf("Saving file '%s' as artefact failed: %s", name, err)
		}
	}
	return nil
}

// copyFile copies the specified source to the destination by opening the source file with os.Open and calling io.Copy
//
// returns an error if either os.Open, os.Create or io.Copy fail
//
func copyFile(source string, destination string) error {
	var sourceFields = logrus.Fields{
		"file_name": source,
	}
	// log messages make this messy!
	Log.WithFields(sourceFields).Infof("Saving file: %s", source)

	// Check whether the source file is a regular file. Prevents accidental copying of entire devices.
	info, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("could not determine file entry: %s", err)
	} else if !info.Mode().IsRegular() {
		return fmt.Errorf("file '%s' is not a regular file; type: '%s'", source, info.Mode().Type())
	}

	Log.WithFields(sourceFields).Debugf("Opening file: %s", source)
	// opening source file in read only mode
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	// deferring src.CLose in case of abrupt function return
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			Log.Debugf("Could not close source file: %s", src.Name())
		}
	}(src)
	Log.WithFields(sourceFields).Debugf("Creating destination file: %s", source)
	// creating destination file in output directory
	dest, err := os.Create(destination)
	if err != nil {
		return err
	}
	// deferring dest.Close in case of abrupt function return
	defer func(dest *os.File) {
		err := dest.Close()
		if err != nil {
			Log.Debugf("Could not close source file: %s", dest.Name())
		}
	}(dest)
	Log.WithFields(sourceFields).Debugf("Copying file content from '%s' to '%s'", source, destination)
	// copying content of source file to destination
	_, err = io.Copy(dest, src)
	if err != nil {
		// removing destination file, as contents are deemed useless since copying failed
		_ = os.Remove(destination)
		return err
	}
	return nil
}

// writeArtefactsToDirectory creates the necessary specified directories for the artifacts and creates the files in
// which the content of the Artefacts is written.
//
// returns an Error if os.MkdirAll, os.Create, WriteString or Close fail.
//
func writeArtefactsToDirectory(artefacts []memory.Artefact, parentDirs ...string) error {
	for n := range artefacts {
		err := os.MkdirAll(path.Join(parentDirs...), os.ModePerm)
		if err != nil {
			return err
		}
		artefactFileName := path.Join(append(parentDirs, generateName(artefacts[n].Name))...)
		file, err := os.Create(artefactFileName)
		if err != nil {
			return err
		}
		// write artefact content to file
		_, err = file.WriteString(artefacts[n].Content)
		if err != nil {
			return err
		}

		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// addTasks checks whether the specified task exists and contains at least one memory.Artefact,
// in which case it will append the task to the specified list and search the task for potential subsequent tasks and
// their possible Artefacts.
//
// returns a list of Tasks.
//
func addTasks(list []*memory.Task, task *memory.Task) []*memory.Task {
	if task != nil && len(task.GetArtefacts()) != 0 {
		list = append(list, task)
		if task.OnSuccess != nil {
			list = addTasks(list, task.OnSuccess)
		} else if task.OnFail != nil {
			list = addTasks(list, task.OnFail)
		} else if task.OnError != nil {
			list = addTasks(list, task.OnError)
		}
	}
	return list
}

// generateName replaces all instances of blanks with underscores for every string provided and concatenates theses
// strings separated by underscores, making the strings suitable as file names.
//
// returns a file system naming compliant string
//
func generateName(parts ...string) string {
	for n := range parts {
		re1 := regexp.MustCompile(`[^a-zA-Z0-9_.-]`)
		re2 := regexp.MustCompile(`_{2,}`)
		parts[n] = re1.ReplaceAllString(parts[n], constants.SEPARATOR_UNDERSCORE)
		parts[n] = re2.ReplaceAllString(parts[n], constants.SEPARATOR_UNDERSCORE)
	}
	return strings.Join(parts, constants.SEPARATOR_UNDERSCORE)
}
