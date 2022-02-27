package token

import "time"

// Maker is an interface for manage token
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	Verify(token string) (*Payload, error)
}
