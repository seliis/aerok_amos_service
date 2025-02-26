package repositories

import (
	"context"
	db "packages/prisma"
	"packages/server/database"
	"packages/server/internal/entities"
)

type ExchangeRateRepository struct{}

func NewExchangeRateRepository() *ExchangeRateRepository {
	return &ExchangeRateRepository{}
}
func (r *ExchangeRateRepository) Create(context context.Context, entity *entities.ExchangeRate) (*entities.ExchangeRate, error) {
	model, err := database.Client.ExchangeRate.CreateOne(
		db.ExchangeRate.Date.Set(entity.Date),
		db.ExchangeRate.Rate.Set(entity.Rate),
		db.ExchangeRate.Currency.Link(
			db.Currency.ID.Equals(entity.CurrencyID),
		),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.ExchangeRate)(&model.InnerExchangeRate), nil
}

func (r *ExchangeRateRepository) Read(context context.Context, id string) (*entities.ExchangeRate, error) {
	model, err := database.Client.ExchangeRate.FindUnique(
		db.ExchangeRate.ID.Equals(id),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.ExchangeRate)(&model.InnerExchangeRate), nil
}

func (r *ExchangeRateRepository) Update(context context.Context, entity *entities.ExchangeRate) (*entities.ExchangeRate, error) {
	model, err := database.Client.ExchangeRate.FindUnique(
		db.ExchangeRate.ID.Equals(entity.ID),
	).Update(
		db.ExchangeRate.Date.Set(entity.Date),
		db.ExchangeRate.Rate.Set(entity.Rate),
		db.ExchangeRate.Currency.Link(
			db.Currency.ID.Equals(entity.CurrencyID),
		),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.ExchangeRate)(&model.InnerExchangeRate), nil
}

func (r *ExchangeRateRepository) Delete(context context.Context, id string) (*entities.ExchangeRate, error) {
	model, err := database.Client.ExchangeRate.FindUnique(
		db.ExchangeRate.ID.Equals(id),
	).Delete().Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.ExchangeRate)(&model.InnerExchangeRate), nil
}

func (r *ExchangeRateRepository) All(context context.Context) ([]*entities.ExchangeRate, error) {
	models, err := database.Client.ExchangeRate.FindMany().Exec(context)

	if err != nil {
		return nil, err
	}

	var arr []*entities.ExchangeRate

	for _, model := range models {
		arr = append(arr, (*entities.ExchangeRate)(&model.InnerExchangeRate))
	}

	return arr, nil
}
