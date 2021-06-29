package main

import (
	"fmt"
	"os"
	"path/filepath"

	"Auditheia/configParser"
	sanityChecker2 "Auditheia/configParser/sanityChecker"
	"Auditheia/executor"
	"Auditheia/memory"
	"Auditheia/output"

	"github.com/sirupsen/logrus"
)

func main() {

	// reset working directory to executable exPath
	err := resetFilePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// parse CLI flags (initializes options in the process)
	// must be the first call in main to ensure all options are set
	configParser.ParseCLIFlags()

	// initialize the logger
	err = output.InitLog()
	if err != nil {
		output.Log.Errorln(err)
	}

	// parse configuration from file
	if err := configParser.ParseConfigFile(); err != nil {
		output.Log.Errorln(err)
		output.Log.Fatal("failed to correctly parse configuration, exiting...")
	}

	if memory.OptionsList.CheckOnly {
		output.Log.Infoln("check-only option selected, skipping audit execution...")
		auditCount, taskCount := memory.CountAuditsAndTasks()
		output.Log.WithFields(logrus.Fields{
			"audit_count": auditCount,
			"task_count":  taskCount,
		}).Infof("Parsed %v Audits with a total of %v Tasks.", auditCount, taskCount)
		fmt.Printf("Parsed %v Audits with a total of %v Tasks.\n", auditCount, taskCount)
		additionalFilesSane := sanityChecker2.CheckAdditionalFiles(memory.AdditionalFiles)
		auditsSane := sanityChecker2.CheckAuditList()
		if !additionalFilesSane || !auditsSane {
			fmt.Printf("config has some sanity issues - see logs for details\n")
		} else {
			fmt.Printf("no problems in config found\n")
		}
	} else {

		// execute audits
		for n := range memory.AuditList {
			output.Log.Infof("Executing Audit %v...", n)
			err := executor.ExecuteAudit(&memory.AuditList[n])
			if err != nil {
				output.Log.Errorln(err)
			}
		}

		// save artefacts
		err = output.SaveArtefacts()
		if err != nil {
			output.Log.Errorln(err)
		}

		// generate report
		report := output.GenerateReport()

		if memory.OptionsList.ReportYaml {
			_, err := report.ToYAML()
			if err != nil {
				output.Log.Errorln(err)
			}
		} else {
			_, err := report.ToJSON(true)
			if err != nil {
				output.Log.Errorln(err)
			}
		}

		// write report to file
		err = report.WriteToFile()
		if err != nil {
			output.Log.Errorln(err)
		}
		err = output.ZipFolder(memory.OptionsList.BaseFolder)
		if err != nil {
			output.Log.Errorln(err)
		}
		fmt.Println("program execution finished, exiting...")
	}

}

func resetFilePath() error {
	exPath, err := os.Executable()
	if err != nil {
		return err
	}
	if err = os.Chdir(filepath.Dir(exPath)); err != nil {
		return err
	}
	return nil
}
