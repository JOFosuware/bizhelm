package bcrypt_test

import (
	"crypto/sha256"
	"strings"
	"testing"

	bcryptsvc "github.com/jofosuware/bizhelm/account-service/internal/infrastructure/security/bcrypt"
	bgcrypt "golang.org/x/crypto/bcrypt"
)

func TestHashPasswordAndCompare(t *testing.T) {
	s := &bcryptsvc.Service{}

	cases := []struct {
		name string
		pwd  string
	}{
		{name: "short", pwd: "secret"},
		{name: "empty", pwd: ""},
		{name: "very_long", pwd: strings.Repeat("a", 5000)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := s.HashPassword(tc.pwd)
			if err != nil {
				t.Fatalf("HashPassword error: %v", err)
			}
			if hash == "" {
				t.Fatalf("expected non-empty hash")
			}
			if !s.ComparePassword(hash, tc.pwd) {
				t.Fatalf("ComparePassword failed for case %s", tc.name)
			}
			// wrong password should not validate
			if s.ComparePassword(hash, tc.pwd+"x") {
				t.Fatalf("ComparePassword incorrectly validated wrong password for %s", tc.name)
			}
		})
	}
}

func TestCompareLegacyAndPrehashVariants(t *testing.T) {
	s := &bcryptsvc.Service{}

	// legacy: bcrypt of plain short password
	plain := "legacy-secret"
	legacyHash, err := bgcrypt.GenerateFromPassword([]byte(plain), bgcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate legacy bcrypt sample: %v", err)
	}
	if !s.ComparePassword(string(legacyHash), plain) {
		t.Fatalf("ComparePassword failed for legacy bcrypt hash")
	}

	// pre-hash variant: bcrypt(sha256(password)) for long password
	long := strings.Repeat("b", 400)
	h := sha256.Sum256([]byte(long))
	preHashHash, err := bgcrypt.GenerateFromPassword(h[:], bgcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate pre-hash bcrypt sample: %v", err)
	}
	if !s.ComparePassword(string(preHashHash), long) {
		t.Fatalf("ComparePassword failed for pre-hash bcrypt variant")
	}
}
