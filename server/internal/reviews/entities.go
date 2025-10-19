package reviews

import "github.com/limonanthony/portfolio/internal/common"

type Review struct {
	Id      common.Id
	Name    string
	Email   *string
	Rating  int
	Message *string
	Visible bool
}
