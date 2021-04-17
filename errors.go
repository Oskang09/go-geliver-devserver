package devserver

// Errors
const (
	// Dev Server
	ErrBodyReadFail       = "E0007: fail to read body, %v"
	ErrInvalidJsonRequest = "E0008: invalid json request, %v"
	ErrInvalidPassword    = "E0009: invalid password for devserver"
	ErrFailToListen       = "E0002: failed to listen %v"
)
