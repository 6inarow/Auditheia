package output

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"Auditheia/memory"
	"Auditheia/memory/constants"
)

// ZipFolder uses readDirectory to receive a folder structure and zipFiles to zip the folder structure
// the path of the zip file is based on memory.OptionsList.BaseFolder with the fix name "auditReport.zip"
//
//  ZipFolder("folder")
//
// returns an error if either reading the target directory or zipping the files fails
//
func ZipFolder(zipfolder string) error {

	Log.Traceln("reading log directory")

	files, err := readDirectory(zipfolder, constants.REGEX_DEFAULT_ZIP_NAME, constants.REGEX_ZIP_FILE_EXTENSION)
	if err != nil {
		Log.Errorf("readDirectory error: %s", err.Error())
		return err
	}
	files = append(files, memory.OptionsList.ConfFilePath)
	Log.Traceln("zipping given files")
	zipPath := path.Join(zipfolder, constants.DEFAULT_ZIP_NAME)
	if err = zipFiles(zipPath, files); err != nil {
		Log.Errorf("zipFiles error: %s", err.Error())
		return err
	}
	Log.Traceln("Zipping completed")

	return nil
}

// readDirectory returns the folder structure of directory in a string slice.
// Matches each path with the pattern via regex. If a match is found,
// the path is excluded from directoryStruct. Excludes files or entire directories matched.
//
//  readDirectory("folder")
//
// returns a slice containing the paths of the files contained in the directory and an error if either reading the
// root directory or traversing the directory fail
//
func readDirectory(directory string, excludePatterns ...string) (directoryStruct []string, err error) {

	info, err := os.Stat(directory)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, errors.New("no folder provided")
	}

	// go through the whole directory
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, pattern := range excludePatterns {
			if memory.CheckRegex(pattern, path) {
				if info.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
		}

		if !info.IsDir() {
			directoryStruct = append(directoryStruct, path)
		}
		return nil

	})
	if err != nil {
		return nil, err
	}
	return
}

// ZipFiles compresses one or many files into a single zip archive file.
//
// Param 1: filename is the output zip file's name.
//
// Param 2: files is a list of files to add to the zip.
//
// returns an error if either os.Create or addFileToZip fail
//
func zipFiles(filename string, files []string) error {

	Log.Traceln("Creating zip")
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(newZipFile *os.File) {
		err := newZipFile.Close()
		if err != nil {
			Log.Errorln(err)
		}
	}(newZipFile)

	// create writer for zip file
	zipWriter := zip.NewWriter(newZipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			Log.Errorln(err)
		}
	}(zipWriter)

	Log.Traceln("Adding files to zip")
	// Add files to zip
	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			Log.Warningf("File (%s) could not be found err: %s", file, err.Error())
			continue
		}
	}
	return nil
}

// addFileTopZip adds a file to a zip
//
//  addFileToZip(zipwriter, "audit.json")
//
// returns an error if os.Stat, zip.FileInfoHeader, zipWriter.CreateHeader, os.Open or io.Copy fail
//
func addFileToZip(zipWriter *zip.Writer, filename string) error {

	Log.Traceln("Add files to zip")

	// Get the file information
	Log.Traceln("Get file informations")
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(filename)
	}

	header.Name = filepath.Join(baseDir, strings.TrimPrefix(filename, memory.OptionsList.BaseFolder))

	header.Name = strings.Trim(header.Name, "/\\")

	// If file isn't a directory change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	if !info.IsDir() {
		header.Method = zip.Deflate
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}

	Log.Traceln("Copy file into zip")
	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		Log.Errorf("File copy didn't work error: %s", err.Error())
	}
	return err
}
