package reviews

import (
	"context"
	"net/http"

	"github.com/limonanthony/portfolio/internal/errordefs"
	"github.com/limonanthony/portfolio/internal/router"
)

func RegisterRoutes(parent *router.Router) {
	parent.Group("/reviews", "Reviews", func(r *router.Router) {
		router.Register(r, router.Operation{
			Path:          "",
			Description:   "Add a review",
			Summary:       "Create review",
			Method:        http.MethodPost,
			DefaultStatus: http.StatusCreated,
		}, createReview(NewService(NewRepository())))
		router.Register(r, router.Operation{
			Path:        "",
			Description: "Fetch all visible reviews",
			Summary:     "Get all reviews",
			Method:      http.MethodGet,
		}, getAllReviews(NewService(NewRepository())))
	})
}

func createReview(service Service) func(ctx context.Context, i *creatioRequest) (*struct{ Body int }, error) {
	return func(ctx context.Context, i *creatioRequest) (*struct{ Body int }, error) {
		res, err := service.Create(ctx, i.Body)
		if err != nil {
			return nil, errordefs.ToHttpError(err)
		}
		return &struct {
			Body int
		}{
			Body: int(res),
		}, nil
	}
}

func getAllReviews(service Service) func(ctx context.Context, i *struct{}) (*getReviewResponses, error) {
	return func(ctx context.Context, i *struct{}) (*getReviewResponses, error) {
		values, err := service.GetAll(ctx)
		if err != nil {
			return nil, errordefs.ToHttpError(err)
		}

		presenters := make([]ReviewPresenter, len(values))
		for i, v := range values {
			presenters[i] = ReviewPresenter{
				Email:   v.Email,
				Name:    v.Name,
				Message: v.Message,
				Rating:  v.Rating,
			}
		}

		return &getReviewResponses{Body: presenters}, nil
	}
}
