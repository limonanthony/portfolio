package reviews

import "github.com/limonanthony/portfolio/internal/common"

type CreationDto struct {
	Email   *string `json:"email,omitempty" doc:"Email of the person who created this review"`
	Rating  int     `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
	Name    *string `json:"name,omitempty" doc:"Name of the person who created this review"`
	Message *string `json:"message,omitempty" doc:"Message of the review"`
}

type CreateReviewResponse struct {
	ID      common.Id `json:"id" doc:"The ID of the created review"`
	Message string    `json:"message" doc:"Success message"`
}

type GetAllReviewsResponse struct {
	Reviews []ReviewPresenter `json:"reviews" doc:"List of all visible reviews"`
}

type DeleteReviewPath struct {
	ID common.Id `path:"id" doc:"The ID of the review to delete"`
}
