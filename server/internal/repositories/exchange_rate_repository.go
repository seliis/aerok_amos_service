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

func (r *ExchangeRateRepository) ListByDate(context context.Context, date string) ([]*entities.ExchangeRate, error) {
	models, err := database.Client.ExchangeRate.FindMany(
		db.ExchangeRate.Date.Equals(date),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	var arr []*entities.ExchangeRate

	for _, model := range models {
		arr = append(arr, (*entities.ExchangeRate)(&model.InnerExchangeRate))
	}

	return arr, nil
}

func (r *ExchangeRateRepository) UpsertExchangeRate(context context.Context, entity *entities.ExchangeRate) error {
	_, err := database.Client.Prisma.ExecuteRaw(`
			insert or replace into exchange_rates (
				id,
				date,
				rate,
				currency_id
			) values (
			 	?,
			 	?,
				?,
				?
			) on conflict(
				currency_id,
				date
			) do update set
				rate = excluded.rate
		`,
		entity.ID,
		entity.Date,
		entity.Rate,
		entity.CurrencyID,
	).Exec(context)

	if err != nil {
		return err
	}

	return nil
}

func (r *ExchangeRateRepository) GetExchangeRates(context context.Context, date string) ([]*entities.ExchangeRateWithCurrency, error) {
	models, err := database.Client.ExchangeRate.FindMany(
		db.ExchangeRate.Date.Equals(date),
	).With(
		db.ExchangeRate.Currency.Fetch(),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	var arr []*entities.ExchangeRateWithCurrency

	for _, model := range models {
		arr = append(arr, &entities.ExchangeRateWithCurrency{
			ExchangeRateID: model.ID,
			CurrencyID:     model.CurrencyID,
			Code:           model.RelationsExchangeRate.Currency.Code,
			Name:           model.RelationsExchangeRate.Currency.Name,
			Base:           model.RelationsExchangeRate.Currency.Base,
			Date:           model.Date,
			Rate:           model.Rate,
		})
	}

	return arr, nil
}
