package memory

// Meta struct provides all meta data fields which may be set inside the config file
//
type Meta struct {
	CustomerName string `json:"customer_name" yaml:"customer_name"`
	InitialDate  string `json:"initial_date" yaml:"initial_date"`
	LastChanged  string `json:"last_changed" yaml:"last_changed"`
	Version      string `json:"version" yaml:"version"`
	ConfOS       string `json:"conf_os" yaml:"conf_os"`
	RootRequired bool   `json:"root_required" yaml:"root_required"`
}
