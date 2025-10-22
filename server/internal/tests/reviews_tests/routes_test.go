package reviews_tests

import (
	"net/http"
	"strings"
	"testing"

	"github.com/limonanthony/portfolio/internal/reviews"
	"github.com/limonanthony/portfolio/internal/tests"
)

func TestCreateReviewRoute(t *testing.T) {
	t.Run("should return 201 when body is valid", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)
		resp := testApi.Post("/reviews", map[string]any{
			"email":   "joe@mail.com",
			"rating":  5,
			"name":    "Joe",
			"message": "Hello, world!",
		})

		if resp.Code != http.StatusCreated {
			t.Errorf("got %d, want %d", resp.Code, http.StatusCreated)
		}
	})

	t.Run("should return 422 when json is invalid", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)
		resp := testApi.Post("/reviews", strings.NewReader(`{"email":"joe@mail.com"`))
		if resp.Code != http.StatusUnprocessableEntity {
			t.Errorf("got %d, want %d", resp.Code, http.StatusUnprocessableEntity)
		}
	})

	t.Run("should return 422 when body is empty", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)
		resp := testApi.Post("/reviews", map[string]any{})

		if resp.Code != http.StatusUnprocessableEntity {
			t.Errorf("got %d, want %d", resp.Code, http.StatusUnprocessableEntity)
		}
	})

	t.Run("should return 422 when body is wrong type", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)
		resp := testApi.Post("/reviews", map[string]any{
			"email":  "joe@mail.com",
			"rating": "5",
		})

		if resp.Code != http.StatusUnprocessableEntity {
			t.Errorf("got %d, want %d", resp.Code, http.StatusUnprocessableEntity)
		}
	})

	t.Run("should return 422 when body is invalid", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)
		resp := testApi.Post("/reviews", map[string]any{
			"email":   "joe@mail.com",
			"rating":  6,
			"message": "Hello, world!",
			"name":    "Joe",
		})

		if resp.Code != http.StatusUnprocessableEntity {
			t.Errorf("got %d, want %d -> infos %s", resp.Code, http.StatusUnprocessableEntity, resp.Body)
			t.Errorf("got %d, want %d", resp.Code, http.StatusUnprocessableEntity)
		}
	})

	t.Run("should return 409 when email is already taken", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)

		// Create first review
		testApi.Post("/reviews", map[string]any{
			"email":   "joe@mail.com",
			"rating":  5,
			"message": "Hello",
			"name":    "Joe",
		})

		// Try to create another with the same email
		resp := testApi.Post("/reviews", map[string]any{
			"email":   "joe@mail.com",
			"rating":  4,
			"message": "Hello",
		})

		if resp.Code != http.StatusConflict {
			t.Errorf("got %d, want %d -> infos %s", resp.Code, http.StatusConflict, resp.Body)
		}
	})
}

func TestGetAllReviewsRoute(t *testing.T) {
	t.Run("should return 200 with empty array", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)
		resp := testApi.Get("/reviews")

		if resp.Code != http.StatusOK {
			t.Errorf("got %d, want %d", resp.Code, http.StatusOK)
		}

		if resp.Body == nil {
			t.Errorf("got nil body")
		}

		if !strings.Contains(resp.Body.String(), "[]") {
			t.Errorf("got %s, want %s", resp.Body.String(), "[]")
		}

	})

	t.Run("should return a review if one is there", func(t *testing.T) {
		testRouter, testApi := tests.NewRouterWithDb(t)
		reviews.RegisterRoutes(testRouter)
		testApi.Post("/reviews", map[string]any{
			"email":   "joe@mail.com",
			"name":    "Joe",
			"rating":  5,
			"message": "Hello, world!",
		})

		resp := testApi.Get("/reviews")

		if resp.Code != http.StatusOK {
			t.Errorf("got %d, want %d", resp.Code, http.StatusOK)
		}

		if resp.Body == nil {
			t.Errorf("got nil body")
		}

		if !strings.Contains(resp.Body.String(), "Joe") {
			t.Errorf("got %s, want %s", resp.Body.String(), "[\"Joe\"]")
		}

		if !strings.Contains(resp.Body.String(), "joe@mail.com") {
			t.Errorf("got %s, want %s", resp.Body.String(), "[\"Joe\"]")
		}
	})
}
