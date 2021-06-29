// Package memory contains all public structs and multi use functions used by the program
//
package memory

import (
	"regexp"
	"time"

	"Auditheia/memory/constants"
)

// AuditList contains all Audits. Audits parsed by configParser.ParseConfigFile can be accessed here
//
var AuditList []Audit

// OptionsList ia a struct containing all Options set via flags or default
//
var OptionsList Options

// AdditionalFiles are files to be saved as artefacts in any case.
//
// The file paths are stored in a map, where the file name is the key, to prevent duplicate entries.
//
var AdditionalFiles map[string]bool

// MetaData contains the passed metadata of the configuration file
//
var MetaData Meta

// StartTime only used to keep track of elapsed time at report creation so far
// set start time of program for report
//
var StartTime = time.Now()

// CountAuditsAndTasks counts the amount of audits and tasks provided by the config file saved in memory.AuditList
//
func CountAuditsAndTasks() (int, int) {
	var auditCount = 0
	var taskCount = 0
	for n := range AuditList {
		auditCount++
		taskCount += countTasks(AuditList[n].Tasks...)
	}
	return auditCount, taskCount
}

// countTasks counts all tasks and subtasks
//
//  countTasks(memory.Tasks, memory.Tasks)
//
func countTasks(tasks ...Task) int {
	var taskCount = 0
	for n := range tasks {
		taskCount++
		if tasks[n].OnSuccess != nil {
			taskCount += countTasks(*tasks[n].OnSuccess)
		}
		if tasks[n].OnFail != nil {
			taskCount += countTasks(*tasks[n].OnFail)
		}
		if tasks[n].OnError != nil {
			taskCount += countTasks(*tasks[n].OnError)
		}
	}
	return taskCount
}

// CheckRegex parses a regular expression in a string and returns true if there is at least one match
//
//	CheckRegex(source, '(regular_expression)')
// 	CheckRegex(source, `(?i)(powershell\b)`)
//
func CheckRegex(regex string, input string) bool {
	var re = regexp.MustCompile(regex)
	match := re.MatchString(input)
	return match
}

// SeparateAtChar separates the given string by the given rune as long as the rune isn't inside single or double quotes
//
// if preserveQuotes is set to true the return string slice contains quotes
//
//  SeparateAtChar("command | 'command|command'", '|', true)
//
func SeparateAtChar(input string, c rune, preserveQuotes bool) []string {
	in := []rune(input)
	var out []string
	var currentQuote rune
	var insideQuotes bool
	var escaped bool
	var currentPart []rune
	for _, char := range in {

		if char == constants.QUOTE_DOUBLE && !escaped {
			if insideQuotes {
				if char == currentQuote {
					currentQuote = 0
					insideQuotes = false
					if preserveQuotes {
						currentPart = append(currentPart, char)
					}
				} else {
					currentPart = append(currentPart, char)
				}
			} else {
				currentQuote = constants.QUOTE_DOUBLE
				insideQuotes = true
				if preserveQuotes {
					currentPart = append(currentPart, char)
				}
			}
		} else if char == constants.QUOTE_SINGLE && !escaped {
			if insideQuotes {
				if char == currentQuote {
					currentQuote = 0
					insideQuotes = false
					if preserveQuotes {
						currentPart = append(currentPart, char)
					}
				} else {
					currentPart = append(currentPart, char)
				}
			} else {
				currentQuote = constants.QUOTE_SINGLE
				insideQuotes = true
				if preserveQuotes {
					currentPart = append(currentPart, char)
				}
			}

		} else if char == c && !insideQuotes {
			out = append(out, string(currentPart))
			currentPart = []rune{}
		} else {
			currentPart = append(currentPart, char)
		}
		if char == constants.BACKSLASH && !escaped {
			escaped = true
			continue
		}
		escaped = false

	}

	out = append(out, string(currentPart))

	return out
}

// CleanString replaces regular expressions inside a string
//
//	CleanString(source, '(regular_expression)')
// 	CleanString(source, `(?i)(powershell\b)`)
//
func CleanString(input string, expression string) string {
	re := regexp.MustCompile(expression)
	return re.ReplaceAllString(input, "")
}
