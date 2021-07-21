package mfa

import (
	"time"

	"github.com/xlzd/gotp"
)

// TOTPDevice creates a new TOTP device.
type TOTPDevice struct {
	totp *gotp.TOTP
}

// NewTOTPDevice returns a new TOTPDevice with the given secret.
//
// The secret should be a base32 encoded value.
func NewTOTPDevice(secret string) *TOTPDevice {
	return &TOTPDevice{
		totp: gotp.NewDefaultTOTP(secret),
	}
}

// Generate return the current OTP.
func (d *TOTPDevice) Generate() (*string, error) {
	// get an otp for the secret
	return d.GenerateAt(time.Now().Unix())
}

// GenerateAt returns the OTP for the given timestamp.
func (d *TOTPDevice) GenerateAt(ts int64) (*string, error) {
	// get an otp for the secret
	otp := d.totp.At(int(ts))

	return &otp, nil
}
