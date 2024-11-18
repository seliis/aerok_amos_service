package repositories

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"net/http"
	"packages/src/amos"
	"packages/src/config"
	"packages/src/infra/database"
	"packages/src/infra/rest"
	"packages/src/internal/entities"
	"time"
)

type CurrencyRepository struct{}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}

func (r *CurrencyRepository) GetPrimitiveExchangeRates(date string) ([]entities.ExchangeRatePrimitive, error) {
	var primitives []entities.ExchangeRatePrimitive

	client, err := rest.NewClient()
	if err != nil {
		return nil, err
	}

	for i := 0; i < 3; i++ {
		response, err := client.Get(fmt.Sprintf(config.Server.Service.DataSource, date))
		if err != nil {
			if i == 2 {
				return nil, err
			}

			// Exponential Back-Off
			time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
			continue
		}
		defer response.Body.Close()

		if err := json.NewDecoder(response.Body).Decode(&primitives); err != nil {
			return nil, err
		}

		return primitives, nil
	}

	return nil, fmt.Errorf("failed to get exchange rates")
}

func (r *CurrencyRepository) SaveExchangeRate(data []entities.ExchangeRate) error {
	statement := `
			insert or replace
			into exchange_rates (
				currency_code,
				currency_name,
				date,
				direct_rate
			) values (
			 	?, ?, ?, ?
			) on conflict(
				currency_code,
				date
			) do update set
				direct_rate = excluded.direct_rate
		`

	for _, e := range data {
		if _, err := database.Client.Exec(statement, e.CurrencyCode, e.CurrencyName, e.Date, e.DirectRate); err != nil {
			return err
		}
	}

	return nil
}

func (r *CurrencyRepository) UpdateAmos(exchangeRates []entities.ExchangeRate) error {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Amos.WebserviceId, config.Amos.WebservicePassword)))
	url := fmt.Sprintf("%s%s", config.Amos.BaseUrl, config.Amos.ImportCurrency.Endpoint)

	data, err := xml.Marshal(amos.NewImportCurrency(exchangeRates))
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/xml")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))

	client, err := rest.NewClient()
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("%s", bytes)
	}

	return nil
}
