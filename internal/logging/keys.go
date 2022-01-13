package logging

// Global logging keys attached to all requests.
//
// Practitioners or tooling reading logs may be depending on these keys, so be
// conscious of that when changing them.
const (
	// Go type of the provider selected by mux.
	KeyTfMuxProvider = "tf_mux_provider"
)
