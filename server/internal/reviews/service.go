package reviews

import (
	"context"

	"github.com/limonanthony/portfolio/internal/common"
	"github.com/limonanthony/portfolio/internal/logger"
)

type Service interface {
	common.Creator[CreationDto]
	common.AllGetter[Review]
	common.GetterById[Review]
	common.DeleterById
	GetAllVisible(context.Context) ([]Review, error)
	GetByEmail(context.Context, string) (Review, error)
}

type service struct {
	repository Repository
}

func (s service) Create(ctx context.Context, dto CreationDto) (common.Id, error) {
	logger.SetLevel(logger.LevelDebug)
	logger.Debugf("dto.Email: %v", dto.Email)
	if dto.Email != nil {
		_, err := s.repository.GetByEmail(ctx, *dto.Email)
		if err == nil {
			logger.Debugf("duplicate email: %v", *dto.Email)
			return 0, EmailConflictErr(*dto.Email)
		}
	}

	var name string
	if dto.Name != nil {
		name = *dto.Name
	} else {
		name = "Anon"
	}

	review := Review{
		Id:      0,
		Email:   dto.Email,
		Name:    name,
		Message: dto.Message,
		Rating:  dto.Rating,
		Visible: true,
	}

	return s.repository.Create(ctx, review)
}

func (s service) GetById(ctx context.Context, id common.Id) (Review, error) {
	return s.repository.GetById(ctx, id)
}

func (s service) GetByEmail(ctx context.Context, email string) (Review, error) {
	return s.repository.GetByEmail(ctx, email)
}

func (s service) GetAll(ctx context.Context) ([]Review, error) {
	return s.repository.GetAll(ctx)
}

func (s service) GetAllVisible(ctx context.Context) ([]Review, error) {
	return s.repository.GetAllVisible(ctx)
}

func (s service) DeleteById(ctx context.Context, id common.Id) error {
	return s.repository.DeleteById(ctx, id)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
