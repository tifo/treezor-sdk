package treezor

type ConnectCursor struct {
	Prev    *string `json:"prev,omitempty"`
	Current *string `json:"current,omitempty"`
	Next    *string `json:"next,omitempty"`
}
