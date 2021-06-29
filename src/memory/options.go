package memory

// Options struct contains all fields which may and can be altered via cli flags.
//
type Options struct {
	BaseFolder     string
	ConfFilePath   string
	LogFileName    string
	ReportFileName string
	CheckOnly      bool
	ReportJSON     bool
	ReportYaml     bool
	Verbosity      int
}
