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
				t.Fail()
			}

			if hashed == input {
				t.Fail()
			}
		})

		t.Run("should not return an empty string", func(t *testing.T) {
			input := "test"
			hashed, err := service.Hash(input)
			if err != nil {
				t.Fail()
			}

			if hashed == "" {
				t.Fail()
			}
		})

		t.Run("should hash 71 characters passwords", func(t *testing.T) {
			input := "very_long_password_123_@#$%^&_until_72_characters_because_bcrypt_______)"
			_, err := service.Hash(input)

			if err != nil {
				t.Fail()
			}
		})
	})

	t.Run("Verify method", func(t *testing.T) {
		input := "test"
		hashed, err := service.Hash(input)

		if err != nil {
			t.Fail()
		}

		t.Run("should return true when", func(t *testing.T) {
			t.Run("input is the same as hashed", func(t *testing.T) {
				if service.Verify(input, hashed) == false {
					t.Fail()
				}
			})
		})

		t.Run("should return false when", func(t *testing.T) {
			t.Run("input is not the same as hashed", func(t *testing.T) {
				if service.Verify("test2", hashed) == true {
					t.Fail()
				}
			})

			t.Run("input is empty", func(t *testing.T) {
				if service.Verify("", hashed) == true {
					t.Fail()
				}
			})
		})
	})
}
