package lemon

type (
	ResourceObject struct {
		Type          string `json:"type"`
		ID            string `json:"id,omitempty"`
		Attributes    any    `json:"attributes,omitempty"`
		Relationships any    `json:"relationships,omitempty"`
	}

	GeneralWrapper struct {
		Data any `json:"data"`
	}
)
