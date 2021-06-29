// Package constants provides all default constants the program uses
package constants

// status messages for tasks
const (
	STATUS_SUCCESS string = "Success"
	STATUS_FAIL    string = "Fail"
	STATUS_ERROR   string = "Error"
)

// operators
const (
	OPERATOR_EQUALS       string = "equals"
	OPERATOR_NOT_EQUALS   string = "not equals"
	OPERATOR_CONTAINS     string = "contains"
	OPERATOR_NOT_CONTAINS string = "not contains"
	OPERATOR_LESSER       string = "lesser"
	OPERATOR_GREATER      string = "greater"
)

// default task types
const (
	TASK_TYPE_SCRIPT  string = "script"
	TASK_TYPE_COMMAND string = "command"
)

// default paths
const (
	DEFAULT_BASEFOLDER    string = "auditFiles"
	DEFAULT_CONF          string = "audit.json"
	DEFAULT_LOG           string = "audit.log"
	DEFAULT_REPORT        string = "report.json"
	DEFAULT_SCRIPTS_DIR   string = "scripts"
	DEFAULT_ARTEFACTS_DIR string = "artefacts"
	DEFAULT_ZIP_NAME      string = "auditReport.zip"
)

// default suffixes
const (
	SUFFIX_JSON string = ".json"
	SUFFIX_YAML string = ".yaml"
)

// default separators
const (
	SEPARATOR_UNDERSCORE = "_"
)

// verbosity levels
const (
	VERBOSITY_ERROR   int = 0
	VERBOSITY_WARNING int = 1
	VERBOSITY_INFO    int = 2
	VERBOSITY_DEBUG   int = 3
	VERBOSITY_TRACE   int = 4
)

// IO rerouting symbols
const (
	SHELL_PIPE        rune = '|'
	SHELL_INPUT_PIPE  rune = '<'
	SHELL_OUTPUT_PIPE rune = '>'
)

// quotes
const (
	QUOTE_SINGLE rune = '\''
	QUOTE_DOUBLE rune = '"'
)

// backslash
const BACKSLASH rune = '\\'

// default regexes
const (
	REGEX_ZIP_FILE_EXTENSION                   string = `.+\.zip\b`
	REGEX_DEFAULT_ZIP_NAME                     string = `auditReport\.zip\b`
	REGEX_POWERSHELL                           string = `(?i)(powershell\b)`
	REGEX_WINDOWS_OPTION_SLASH_C               string = `(?i)(/c\b)`
	REGEX_UNIX_OPTION_HYPHEN_C                 string = `(?i)(-c\b)`
	REGEX_UNIX_SHELL_BASH                      string = `(?i)(bash\b|shell\b)`
	REGEX_WINDOWS_INVALID_FILE_NAME_CHARACTERS string = `[\\/'"?%*|<> ]`
	REGEX_UNIX_INVALID_FILE_NAME_CHARACTERS    string = `/`
	REGEX_JSON_FILE_EXTENSION                  string = `(?i)(^\.json$)`
	REGEX_YAML_FILE_EXTENSION                  string = `(?i)(^\.(yaml|yml)$)`
	REGEX_NOT_ALPHANUMERICAL                   string = `[^a-zA-Z0-9-.]+`
)
