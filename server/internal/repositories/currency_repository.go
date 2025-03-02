package repositories

import (
	"context"
	db "packages/prisma"
	"packages/server/database"
	"packages/server/internal/entities"
)

type CurrencyRepository struct{}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}

func (r *CurrencyRepository) Create(context context.Context, entity *entities.Currency) (*entities.Currency, error) {
	model, err := database.Client.Currency.CreateOne(
		db.Currency.Code.Set(entity.Code),
		db.Currency.Name.Set(entity.Name),
		db.Currency.Base.Set(entity.Base),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.Currency)(&model.InnerCurrency), nil
}

func (r *CurrencyRepository) Read(context context.Context, currencyCode string) (*entities.Currency, error) {
	model, err := database.Client.Currency.FindUnique(
		db.Currency.Code.Equals(currencyCode),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.Currency)(&model.InnerCurrency), nil
}

func (r *CurrencyRepository) Update(context context.Context, entity *entities.Currency) (*entities.Currency, error) {
	model, err := database.Client.Currency.FindUnique(
		db.Currency.Code.Equals(entity.Code),
	).Update(
		db.Currency.Code.Set(entity.Code),
		db.Currency.Name.Set(entity.Name),
		db.Currency.Base.Set(entity.Base),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.Currency)(&model.InnerCurrency), nil
}

func (r *CurrencyRepository) Delete(context context.Context, currencyCode string) (*entities.Currency, error) {
	model, err := database.Client.Currency.FindUnique(
		db.Currency.Code.Equals(currencyCode),
	).Delete().Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.Currency)(&model.InnerCurrency), nil
}

func (r *CurrencyRepository) All(context context.Context) ([]*entities.Currency, error) {
	models, err := database.Client.Currency.FindMany().Exec(context)

	if err != nil {
		return nil, err
	}

	var arr []*entities.Currency

	for _, model := range models {
		arr = append(arr, (*entities.Currency)(&model.InnerCurrency))
	}

	return arr, nil
}
