package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/limonanthony/portfolio/internal/env"
	"github.com/limonanthony/portfolio/internal/tests"
)

func TestGet(t *testing.T) {
	t.Run("should return the correct value", func(t *testing.T) {
		const key = "TEST_KEY"
		const value = "randomValue"

		err := os.Setenv(key, value)
		if err != nil {
			t.Error(err)
		}

		tests.AssertNotPanic(t, func() {
			envValue := env.Get(key)
			if envValue != value {
				t.Errorf("Expected %s, got %s", value, envValue)
			}
		})
	})

	t.Run("should always return a string", func(t *testing.T) {
		const key = "TEST_KEY"
		const value = true

		err := os.Setenv(key, fmt.Sprintf("%t", value))
		if err != nil {
			t.Error(err)
		}

		tests.AssertNotPanic(t, func() {
			envValue := env.Get(key)
			if envValue != fmt.Sprintf("%t", value) {
				t.Errorf("Expected %s, got %s", fmt.Sprintf("%t", value), envValue)
			}
		})
	})

	t.Run("should panic when key not exist", func(t *testing.T) {
		tests.AssertPanic(t, func() {
			env.Get("RANDOM_KEY")
		})
	})
}

func TestGetBool(t *testing.T) {
	t.Run("should return the correct value as bool", func(t *testing.T) {
		const key = "TEST_KEY"
		const value = true

		err := os.Setenv(key, fmt.Sprintf("%t", value))
		if err != nil {
			t.Error(err)
		}

		tests.AssertNotPanic(t, func() {
			envValue := env.GetBool(key)
			if envValue != value {
				t.Errorf("Expected %t, got %t", value, envValue)
			}
		})
	})

	t.Run("should panic when key not exist", func(t *testing.T) {
		tests.AssertPanic(t, func() {
			env.GetBool("RANDOM_KEY")
		})
	})

	t.Run("should panic when value is not a bool", func(t *testing.T) {
		err := os.Setenv("RANDOM_KEY", "randomValue")
		if err != nil {
			t.Error(err)
		}

		tests.AssertPanic(t, func() {
			env.GetBool("RANDOM_KEY")
		})
	})
}

func TestGetInt(t *testing.T) {
	t.Run("should return the correct value as int", func(t *testing.T) {
		const key = "TEST_KEY"
		const value = 42

		err := os.Setenv(key, fmt.Sprintf("%d", value))
		if err != nil {
			t.Error(err)
		}

		tests.AssertNotPanic(t, func() {
			envValue := env.GetInt(key)
			if envValue != value {
				t.Errorf("Expected %d, got %d", value, envValue)
			}
		})
	})

	t.Run("should panic when key not exist", func(t *testing.T) {
		tests.AssertPanic(t, func() {
			env.GetInt("RANDOM_KEY")
		})
	})

	t.Run("should panic when value not int", func(t *testing.T) {
		err := os.Setenv("RANDOM_KEY", "randomValue")
		if err != nil {
			t.Error(err)
		}

		tests.AssertPanic(t, func() {
			env.GetInt("RANDOM_KEY")
		})
	})
}
