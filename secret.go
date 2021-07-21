package mfa

import "github.com/xlzd/gotp"

// NewSecret returns a secret of given length.
func NewSecret(l int) string {
	return gotp.RandomSecret(l)
}
