package ccv3

// APILink represents a generic link from a response object.
type APILink struct {
	// HREF is the fully qualified URL for the link.
	HREF   string `json:"href"`
	Method string `json:"method"`

	// Meta contains additional metadata about the API.
	Meta struct {
		// Version of the API
		Version string `json:"version"`
	} `json:"meta"`
}

// APILinks is a directory of follow-up urls for the resource.
type APILinks map[string]APILink
