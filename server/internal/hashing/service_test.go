package hashing_test

import (
	"testing"

	"github.com/limonanthony/portfolio/internal/hashing"
)

func TestHashingService(t *testing.T) {
	service := hashing.NewService()

	t.Run("Hash method", func(t *testing.T) {
		t.Run("should not return the same string", func(t *testing.T) {
			input := "test"
			hashed, err := service.Hash(input)
			if err != nil {
				t.Errorf("Hash failed: %v", err)
			}

			if hashed == input {
				t.Errorf("Hash should not equal input: %s", input)
			}
		})

		t.Run("should not return an empty string", func(t *testing.T) {
			input := "test"
			hashed, err := service.Hash(input)
			if err != nil {
				t.Errorf("Hash failed: %v", err)
			}

			if hashed == "" {
				t.Error("Hash should not be empty")
			}
		})

		t.Run("should hash 71 characters passwords", func(t *testing.T) {
			input := "very_long_password_123_@#$%^&_until_72_characters_because_bcrypt_______)"
			_, err := service.Hash(input)

			if err != nil {
				t.Errorf("Hash failed for long password: %v", err)
			}
		})
	})

	t.Run("Verify method", func(t *testing.T) {
		input := "test"
		hashed, err := service.Hash(input)

		if err != nil {
			t.Errorf("Hash failed: %v", err)
		}

		t.Run("should return true when", func(t *testing.T) {
			t.Run("input is the same as hashed", func(t *testing.T) {
				if service.Verify(input, hashed) == false {
					t.Error("Verify should return true for correct input")
				}
			})
		})

		t.Run("should return false when", func(t *testing.T) {
			t.Run("input is not the same as hashed", func(t *testing.T) {
				if service.Verify("test2", hashed) == true {
					t.Error("Verify should return false for incorrect input")
				}
			})

			t.Run("input is empty", func(t *testing.T) {
				if service.Verify("", hashed) == true {
					t.Error("Verify should return false for empty input")
				}
			})
		})
	})
}
