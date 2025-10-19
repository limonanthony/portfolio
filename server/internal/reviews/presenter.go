package reviews

type ReviewPresenter struct {
	Name    string  `json:"name"`
	Email   *string `json:"email"`
	Rating  int     `json:"rating"`
	Message *string `json:"message"`
}
