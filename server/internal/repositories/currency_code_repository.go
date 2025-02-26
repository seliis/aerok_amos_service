package repositories

import (
	"context"
	db "packages/prisma"
	"packages/server/database"
	"packages/server/internal/entities"
)

type CurrencyCodeRepository struct{}

func NewCurrencyCodeRepository() *CurrencyCodeRepository {
	return &CurrencyCodeRepository{}
}

func (r *CurrencyCodeRepository) Create(context context.Context, entity *entities.CurrencyCode) (*entities.CurrencyCode, error) {
	model, err := database.Client.CurrencyCode.CreateOne(
		db.CurrencyCode.Code.Set(entity.Code),
		db.CurrencyCode.Name.Set(entity.Name),
		db.CurrencyCode.Base.Set(entity.Base),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.CurrencyCode)(&model.InnerCurrencyCode), nil
}

func (r *CurrencyCodeRepository) Read(context context.Context, id string) (*entities.CurrencyCode, error) {
	model, err := database.Client.CurrencyCode.FindUnique(
		db.CurrencyCode.ID.Equals(id),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.CurrencyCode)(&model.InnerCurrencyCode), nil
}

func (r *CurrencyCodeRepository) Update(context context.Context, entity *entities.CurrencyCode) (*entities.CurrencyCode, error) {
	model, err := database.Client.CurrencyCode.FindUnique(
		db.CurrencyCode.ID.Equals(entity.ID),
	).Update(
		db.CurrencyCode.Code.Set(entity.Code),
		db.CurrencyCode.Name.Set(entity.Name),
		db.CurrencyCode.Base.Set(entity.Base),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.CurrencyCode)(&model.InnerCurrencyCode), nil
}

func (r *CurrencyCodeRepository) Delete(context context.Context, id string) (*entities.CurrencyCode, error) {
	model, err := database.Client.CurrencyCode.FindUnique(
		db.CurrencyCode.ID.Equals(id),
	).Delete().Exec(context)

	if err != nil {
		return nil, err
	}

	return (*entities.CurrencyCode)(&model.InnerCurrencyCode), nil
}

func (r *CurrencyCodeRepository) All(context context.Context) ([]*entities.CurrencyCode, error) {
	models, err := database.Client.CurrencyCode.FindMany().Exec(context)

	if err != nil {
		return nil, err
	}

	var arr []*entities.CurrencyCode

	for _, model := range models {
		arr = append(arr, (*entities.CurrencyCode)(&model.InnerCurrencyCode))
	}

	return arr, nil
}
