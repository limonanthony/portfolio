package logger

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/limonanthony/portfolio/internal/tests"
)

func TestLoggerFunctions(t *testing.T) {
	t.Run("Debug method", func(t *testing.T) {
		t.Run("should have the DEBUG key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelDebug)
			defer SetLevel(LevelInfo)

			Debug("test debug message")

			output := buf.String()
			if !strings.Contains(output, "DEBUG") {
				t.Errorf("Expected DEBUG key, got: %s", output)
			}
		})

		t.Run("should contain the message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelDebug)
			defer SetLevel(LevelInfo)

			Debug("test debug message")

			output := buf.String()
			if !strings.Contains(output, "test debug message") {
				t.Errorf("Expected message content, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelDebug)
			defer SetLevel(LevelInfo)

			EnableColor(true)
			defer EnableColor(true)

			Debug("test debug message")

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in debug output")
			}
		})
	})

	t.Run("Info method", func(t *testing.T) {
		t.Run("should have the INFO key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Info("test info message")

			output := buf.String()
			if !strings.Contains(output, "INFO") {
				t.Errorf("Expected INFO key, got: %s", output)
			}
		})

		t.Run("should contain the message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Info("test info message")

			output := buf.String()
			if !strings.Contains(output, "test info message") {
				t.Errorf("Expected message content, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Info("test info message")

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in info output")
			}
		})
	})

	t.Run("Success method", func(t *testing.T) {
		t.Run("should have the SUCCESS key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Success("test success message")

			output := buf.String()
			if !strings.Contains(output, "SUCCESS") {
				t.Errorf("Expected SUCCESS key, got: %s", output)
			}
		})

		t.Run("should contain the message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Success("test success message")

			output := buf.String()
			if !strings.Contains(output, "test success message") {
				t.Errorf("Expected message content, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Success("test success message")

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in success output")
			}
		})
	})

	t.Run("Warn method", func(t *testing.T) {
		t.Run("should have the WARN key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Warn("test warning message")

			output := buf.String()
			if !strings.Contains(output, "WARN") {
				t.Errorf("Expected WARN key, got: %s", output)
			}
		})

		t.Run("should contain the message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Warn("test warning message")

			output := buf.String()
			if !strings.Contains(output, "test warning message") {
				t.Errorf("Expected message content, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Warn("test warning message")

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in warn output")
			}
		})
	})

	t.Run("Error method", func(t *testing.T) {
		t.Run("should have the ERROR key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Error("test error message")

			output := buf.String()
			if !strings.Contains(output, "ERROR") {
				t.Errorf("Expected ERROR key, got: %s", output)
			}
		})

		t.Run("should contain the message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Error("test error message")

			output := buf.String()
			if !strings.Contains(output, "test error message") {
				t.Errorf("Expected message content, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Error("test error message")

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in error output")
			}
		})
	})

	t.Run("Panic method", func(t *testing.T) {
		t.Run("should have the PANIC key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			tests.AssertPanic(t, func() {
				Panic("test panic message")
			})

			output := buf.String()
			if !strings.Contains(output, "PANIC") {
				t.Errorf("Expected PANIC key, got: %s", output)
			}
		})

		t.Run("should contain the message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			tests.AssertPanic(t, func() {
				Panic("test panic message")
			})

			output := buf.String()
			if !strings.Contains(output, "test panic message") {
				t.Errorf("Expected message content, got: %s", output)
			}
		})

		t.Run("should panic", func(t *testing.T) {
			tests.AssertPanic(t, func() {
				Panic("test panic message")
			})
		})
	})
}

func TestLoggerFormattedFunctions(t *testing.T) {
	t.Run("Debugf method", func(t *testing.T) {
		t.Run("should have the DEBUG key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelDebug)
			defer SetLevel(LevelInfo)

			Debugf("test %s message with %d", "debug", 42)

			output := buf.String()
			if !strings.Contains(output, "DEBUG") {
				t.Errorf("Expected DEBUG key, got: %s", output)
			}
		})

		t.Run("should contain the formatted message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelDebug)
			defer SetLevel(LevelInfo)

			Debugf("test %s message with %d", "debug", 42)

			output := buf.String()
			expected := "test debug message with 42"
			if !strings.Contains(output, expected) {
				t.Errorf("Expected formatted message, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelDebug)
			defer SetLevel(LevelInfo)

			EnableColor(true)
			defer EnableColor(true)

			Debugf("test %s message with %d", "debug", 42)

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in debugf output")
			}
		})
	})

	t.Run("Infof method", func(t *testing.T) {
		t.Run("should have the INFO key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Infof("test %s message with %d", "info", 42)

			output := buf.String()
			if !strings.Contains(output, "INFO") {
				t.Errorf("Expected INFO key, got: %s", output)
			}
		})

		t.Run("should contain the formatted message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Infof("test %s message with %d", "info", 42)

			output := buf.String()
			expected := "test info message with 42"
			if !strings.Contains(output, expected) {
				t.Errorf("Expected formatted message, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Infof("test %s message with %d", "info", 42)

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in infof output")
			}
		})
	})

	t.Run("Successf method", func(t *testing.T) {
		t.Run("should have the SUCCESS key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Successf("test %s message with %d", "success", 42)

			output := buf.String()
			if !strings.Contains(output, "SUCCESS") {
				t.Errorf("Expected SUCCESS key, got: %s", output)
			}
		})

		t.Run("should contain the formatted message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Successf("test %s message with %d", "success", 42)

			output := buf.String()
			expected := "test success message with 42"
			if !strings.Contains(output, expected) {
				t.Errorf("Expected formatted message, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Successf("test %s message with %d", "success", 42)

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in successf output")
			}
		})
	})

	t.Run("Warnf method", func(t *testing.T) {
		t.Run("should have the WARN key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Warnf("test %s message with %d", "warning", 42)

			output := buf.String()
			if !strings.Contains(output, "WARN") {
				t.Errorf("Expected WARN key, got: %s", output)
			}
		})

		t.Run("should contain the formatted message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Warnf("test %s message with %d", "warning", 42)

			output := buf.String()
			expected := "test warning message with 42"
			if !strings.Contains(output, expected) {
				t.Errorf("Expected formatted message, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Warnf("test %s message with %d", "warning", 42)

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in warnf output")
			}
		})
	})

	t.Run("Errorf method", func(t *testing.T) {
		t.Run("should have the ERROR key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Errorf("test %s message with %d", "error", 42)

			output := buf.String()
			if !strings.Contains(output, "ERROR") {
				t.Errorf("Expected ERROR key, got: %s", output)
			}
		})

		t.Run("should contain the formatted message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			Errorf("test %s message with %d", "error", 42)

			output := buf.String()
			expected := "test error message with 42"
			if !strings.Contains(output, expected) {
				t.Errorf("Expected formatted message, got: %s", output)
			}
		})

		t.Run("should have the correct color", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Errorf("test %s message with %d", "error", 42)

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in errorf output")
			}
		})
	})

	t.Run("Panicf method", func(t *testing.T) {
		t.Run("should have the PANIC key", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			tests.AssertPanic(t, func() {
				Panicf("test %s message with %d", "panic", 42)
			})

			output := buf.String()
			if !strings.Contains(output, "PANIC") {
				t.Errorf("Expected PANIC key, got: %s", output)
			}
		})

		t.Run("should contain the formatted message", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			tests.AssertPanic(t, func() {
				Panicf("test %s message with %d", "panic", 42)
			})

			output := buf.String()
			expected := "test panic message with 42"
			if !strings.Contains(output, expected) {
				t.Errorf("Expected formatted message, got: %s", output)
			}
		})

		t.Run("should panic", func(t *testing.T) {
			tests.AssertPanic(t, func() {
				Panicf("test %s message with %d", "panic", 42)
			})
		})
	})
}

func TestLevelManagement(t *testing.T) {
	t.Run("SetLevel method", func(t *testing.T) {
		t.Run("should filter messages below current level", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelInfo)
			defer SetLevel(LevelInfo)

			Debug("this should not appear")
			Info("this should appear")

			output := buf.String()
			if strings.Contains(output, "this should not appear") {
				t.Error("Debug message should be filtered out")
			}
			if !strings.Contains(output, "this should appear") {
				t.Error("Info message should not be filtered")
			}
		})

		t.Run("should allow all messages when level is DEBUG", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelDebug)
			defer SetLevel(LevelInfo)

			Debug("debug message")
			Info("info message")
			Success("success message")
			Warn("warning message")
			Error("error message")

			output := buf.String()
			expectedMessages := []string{"debug message", "info message", "success message", "warning message", "error message"}
			for _, msg := range expectedMessages {
				if !strings.Contains(output, msg) {
					t.Errorf("Expected message '%s' to be logged", msg)
				}
			}
		})

		t.Run("should filter messages when level is ERROR", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			SetLevel(LevelError)
			defer SetLevel(LevelInfo)

			Debug("debug message")
			Info("info message")
			Success("success message")
			Warn("warning message")
			Error("error message")

			output := buf.String()
			filteredMessages := []string{"debug message", "info message", "success message", "warning message"}
			for _, msg := range filteredMessages {
				if strings.Contains(output, msg) {
					t.Errorf("Message '%s' should be filtered out", msg)
				}
			}
			if !strings.Contains(output, "error message") {
				t.Error("Error message should not be filtered")
			}
		})
	})
}

func TestColorSettings(t *testing.T) {
	t.Run("EnableColor method", func(t *testing.T) {
		t.Run("should include color codes when color is enabled", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(true)
			defer EnableColor(true)

			Info("test message")

			output := buf.String()
			if !strings.Contains(output, "\033[") {
				t.Error("Expected color codes in output when color is enabled")
			}
		})

		t.Run("should not include color codes when color is disabled", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			EnableColor(false)
			defer EnableColor(true)

			Info("test message")

			output := buf.String()
			if strings.Contains(output, "\033[") {
				t.Error("Expected no color codes in output when color is disabled")
			}
		})
	})
}

func TestLoggingMiddleware(t *testing.T) {
	t.Run("LoggingMiddleware function", func(t *testing.T) {
		t.Run("should log request start and end", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("test response"))
			})

			loggedHandler := LoggingMiddleware(handler)

			req := httptest.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()

			loggedHandler.ServeHTTP(w, req)

			output := buf.String()

			if !strings.Contains(output, "request_start") {
				t.Error("Expected request_start log")
			}

			if !strings.Contains(output, "request_end") {
				t.Error("Expected request_end log")
			}

			if !strings.Contains(output, "method=GET") {
				t.Error("Expected method in log")
			}
			if !strings.Contains(output, "path=/test") {
				t.Error("Expected path in log")
			}
			if !strings.Contains(output, "status=200") {
				t.Error("Expected status in log")
			}
		})

		t.Run("should generate request ID when not provided", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			loggedHandler := LoggingMiddleware(handler)
			req := httptest.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()

			loggedHandler.ServeHTTP(w, req)

			requestID := w.Header().Get("X-Request-ID")
			if requestID == "" {
				t.Error("Expected request ID to be set in response header")
			}

			output := buf.String()
			if !strings.Contains(output, fmt.Sprintf("request_id=%s", requestID)) {
				t.Error("Expected request ID in log output")
			}
		})

		t.Run("should use provided request ID", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			defer SetOutput(nil)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			loggedHandler := LoggingMiddleware(handler)
			req := httptest.NewRequest("GET", "/test", nil)
			req.Header.Set("X-Request-ID", "custom-request-id")
			w := httptest.NewRecorder()

			loggedHandler.ServeHTTP(w, req)

			output := buf.String()
			if !strings.Contains(output, "request_id=custom-request-id") {
				t.Error("Expected custom request ID in log output")
			}
		})

		t.Run("should log different levels based on status code", func(t *testing.T) {
			testCases := []struct {
				statusCode    int
				expectedLevel string
			}{
				{200, "SUCCESS"},
				{400, "WARN"},
				{500, "ERROR"},
			}

			for _, tc := range testCases {
				t.Run(fmt.Sprintf("status %d should log as %s", tc.statusCode, tc.expectedLevel), func(t *testing.T) {
					var buf bytes.Buffer
					SetOutput(&buf)
					defer SetOutput(nil)

					handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(tc.statusCode)
					})

					loggedHandler := LoggingMiddleware(handler)
					req := httptest.NewRequest("GET", "/test", nil)
					w := httptest.NewRecorder()

					loggedHandler.ServeHTTP(w, req)

					output := buf.String()
					if !strings.Contains(output, tc.expectedLevel) {
						t.Errorf("Expected %s level for status %d, got: %s", tc.expectedLevel, tc.statusCode, output)
					}
				})
			}
		})
	})
}
