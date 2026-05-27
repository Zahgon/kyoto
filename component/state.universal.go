package component

// Universal is a default universal component state implementation.
// It uses combination of JSON, base64 and URI encoding
// to marshal and unmarshal the state.
type Universal struct {
	Name
}

func (*Universal) Marshal(src any) string {
	_ = "STUB: not implemented"
	// Marshal into json
	return ""
}

// Encode to URI representation to avoid html breaking

// Encode to base64

// Return

func (*Universal) Unmarshal(dst any, str string) {
	_ = "STUB: not implemented"
	// Decode from base64
	return
}

// Decode from URI representation

// Unmarshal from json
