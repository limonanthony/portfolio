package reviews

import (
	"context"

	"github.com/limonanthony/portfolio/internal/common"
	"github.com/limonanthony/portfolio/internal/database"
	"gorm.io/gorm"
)

type Repository interface {
	common.Creator[Review]
	common.GetterById[Review]
	common.AllGetter[Review]
	common.DeleterById
	GetAllVisible(context.Context) ([]Review, error)
	GetByEmail(context.Context, string) (Review, error)
}

type repository struct {
}

func (r repository) Create(ctx context.Context, dto Review) (common.Id, error) {
	tx := ctx.Value(database.ContextKey).(*gorm.DB)

	if err := gorm.G[Review](tx).Create(ctx, &dto); err != nil {
		return 0, err
	}

	return dto.Id, nil
}

func (r repository) GetById(ctx context.Context, id common.Id) (Review, error) {
	tx := ctx.Value(database.ContextKey).(*gorm.DB)

	review, err := gorm.G[Review](tx).Where("id = ?", id).First(ctx)
	if err != nil {
		return Review{}, err
	}

	return review, nil
}

func (r repository) GetByEmail(ctx context.Context, email string) (Review, error) {
	tx := ctx.Value(database.ContextKey).(*gorm.DB)

	review, err := gorm.G[Review](tx).Where("email = ?", email).First(ctx)
	if err != nil {
		return Review{}, err
	}

	return review, nil
}

func (r repository) GetAll(ctx context.Context) ([]Review, error) {
	tx := ctx.Value(database.ContextKey).(*gorm.DB)

	reviews_, err := gorm.G[Review](tx).Find(ctx)
	if err != nil {
		return nil, err
	}

	return reviews_, nil
}

func (r repository) GetAllVisible(ctx context.Context) ([]Review, error) {
	tx := ctx.Value(database.ContextKey).(*gorm.DB)

	reviews_, err := gorm.G[Review](tx).Where("visible = ?", true).Find(ctx)
	if err != nil {
		return nil, err
	}

	return reviews_, nil
}

func (r repository) DeleteById(ctx context.Context, id common.Id) error {
	tx := ctx.Value(database.ContextKey).(*gorm.DB)

	amount, err := gorm.G[Review](tx).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if amount == 0 {
		return NotFoundErr(id)
	}

	return nil
}

func NewRepository() Repository {
	return &repository{}
}
