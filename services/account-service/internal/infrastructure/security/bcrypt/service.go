package bcrypt

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
}

// HashPassword hashes the given password using bcrypt. If the password exceeds
// bcrypt's 72-byte limit, it is first pre-hashed with SHA-256 to ensure security.
func (s *Service) HashPassword(password string) (string, error) {
	// bcrypt has a 72-byte input limit. To support arbitrary-length passwords
	// safely, pre-hash long passwords with SHA-256 before passing to bcrypt.
	pass := []byte(password)
	if len(pass) > 72 {
		h := sha256.Sum256(pass)
		pass = h[:]
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares the provided password with the given bcrypt hash. It first
// attempts a direct bcrypt comparison (for legacy hashes of short passwords). If that fails,
// it pre-hashes the candidate password with SHA-256 and compares again (to support hashes of long passwords).
func (s *Service) ComparePassword(hash, password string) bool {
	// Try direct bcrypt comparison first (legacy: bcrypt(password)).
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err == nil {
		return true
	}

	// If that fails, try pre-hashing the candidate with SHA-256 and compare
	// (handles passwords that were pre-hashed before bcrypt to avoid the 72-byte limit).
	h := sha256.Sum256([]byte(password))
	if err := bcrypt.CompareHashAndPassword([]byte(hash), h[:]); err == nil {
		return true
	}

	return false
}