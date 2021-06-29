// Package output contains all functions which generate an output to files or the console.
// This includes the everything about logging, the report generation,
// the saving of artefact files and the zipping feature.
//
package output

import (
	"io"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	"Auditheia/memory"
	"Auditheia/memory/constants"
)

// Log Logger for different verbosity levels as specified in InitLog
// Levels:
//  0 - Error Level     (Only Errors, Panics and Fatal are logged)
//  1 - Warning Level   (Only Warnings and below are logged)
//  2 - Info Level      (Only Infos and below are logged)
//  3 - Debug Level     (Only Debug Messages and below are logged)
//  4 - Trace Level     (Highest log level, everything will be logged if possible, including calling functions and their files)
//
var Log = logrus.New()

// InitLog initiates the logger. Utilizes lumberjack for rotated file handling.
// Sets the loggers output to the logfile and the stdout via io.MultiWriter.
// Logging only to stdout if the file rotation fails.
//
// Returns an error if the rotating file handler fails.
//
func InitLog() error {
	// create a log file rotator with lumberjack
	fileRotator := &lumberjack.Logger{
		Filename:  path.Join(memory.OptionsList.BaseFolder, memory.OptionsList.LogFileName),
		MaxSize:   20,
		MaxAge:    30,
		LocalTime: true,
	}

	// setting log level for logrus depending on passed verbosity level (default: logrus.WarnLevel)
	switch memory.OptionsList.Verbosity {
	case constants.VERBOSITY_ERROR:
		Log.SetLevel(logrus.ErrorLevel)
	case constants.VERBOSITY_WARNING:
		Log.SetLevel(logrus.WarnLevel)
	case constants.VERBOSITY_INFO:
		Log.SetLevel(logrus.InfoLevel)
	case constants.VERBOSITY_DEBUG:
		Log.SetLevel(logrus.DebugLevel)
	case constants.VERBOSITY_TRACE:
		Log.SetLevel(logrus.TraceLevel)
	default:
		Log.SetLevel(logrus.WarnLevel)
	}

	// adds the calling function to the logs at trace lvl. high performance impact!
	if memory.OptionsList.Verbosity >= constants.VERBOSITY_TRACE {
		Log.SetReportCaller(true)
	}

	err := fileRotator.Rotate()
	if err == nil {
		// set output of logrus to the file rotator
		Log.SetOutput(io.MultiWriter(os.Stdout, fileRotator))

		// setting formatter for log entries (default)
		Log.SetFormatter(&logrus.TextFormatter{})

		Log.Infoln("Logging initialized...")
	} else {
		Log.Warnln("Logging only to console")
	}

	return err
}
