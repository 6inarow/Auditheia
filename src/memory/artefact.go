package memory

// Artefact contains the name of the artefact and its content
type Artefact struct {
	Name    string `json:"name" yaml:"name"`
	Content string `json:"content" yaml:"content"`
}

// NewArtefact creates a new artefact object
func NewArtefact(name string, content string) *Artefact {
	return &Artefact{Name: name, Content: content}
}
