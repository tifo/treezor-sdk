package event

import "github.com/tifo/treezor-sdk"

// DocumentEvent represents a document event
type DocumentEvent struct {
	Documents []*treezor.Document `json:"documents"`
}
