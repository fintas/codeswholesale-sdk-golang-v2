package codeswholesalesdkgolangv2

// Links provides REST links
type Links struct {
	Rel         string      `json:"rel"`
	Href        string      `json:"href"`
	Hreflang    interface{} `json:"hreflang"`
	Media       interface{} `json:"media"`
	Title       interface{} `json:"title"`
	Type        interface{} `json:"type"`
	Deprecation interface{} `json:"deprecation"`
}
