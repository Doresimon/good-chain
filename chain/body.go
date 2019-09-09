package chain

// Body ...
type Body struct {
	Type string `json:"type"`
	Action string  `json:"action"`
	Content *Content `json:"content"`
}
// Content ...
type Content struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Index string `json:"index"`
	Pk string`json:"pk"`
	Extra map[string]string`json:"extra"`
}
