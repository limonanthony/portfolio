package reviews

import "github.com/limonanthony/portfolio/internal/common"

type createReviewResponse struct {
	Body common.Id
}

type getReviewResponses struct {
	Body []ReviewPresenter
}
