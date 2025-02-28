package client

import (
	"log"
	"packages/server/config"
	"time"

	"resty.dev/v3"
)

type _KoreaEximCurrencyExchangeData struct {
	RequestResult           int    `json:"result"`
	CurrencyCode            string `json:"cur_unit"`
	CurrencyName            string `json:"cur_nm"`
	BuyingRate              string `json:"ttb"`
	SellingRate             string `json:"tts"`
	ExchangeRate            string `json:"deal_bas_r"`
	BookPrice               string `json:"bkpr"`
	YearHigh                string `json:"yy_efee_r"`
	TenDaysHigh             string `json:"ten_dd_efee_r"`
	MoneyBrokerExchangeRate string `json:"kftc_deal_bas_r"`
	MoneyBrokerBookPrice    string `json:"yy_kftc_bkpr"`
}

func RequestCurrencyExchangeDataFromKoreaExim(date string) ([]*_KoreaEximCurrencyExchangeData, error) {
	client := resty.New().
		SetBaseURL(config.KoreaExim.Host).
		SetRetryCount(30).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(60 * time.Second).
		AddRetryConditions(func(r *resty.Response, err error) bool {
			if err != nil {
				log.Printf("error: unknown error occured on fetching data from korea-exim, retrying...")
				return true
			}

			if r.StatusCode() != 200 {
				log.Printf("error: http status code %s received when fetching data from korea-exim, retrying....", r.Status())
				return true
			}

			return false
		})
	defer client.Close()

	queryParams := map[string]string{
		"authkey":    config.KoreaExim.Auth,
		"searchdate": date,
		"data":       "AP01",
	}

	var result []*_KoreaEximCurrencyExchangeData

	_, err := client.R().SetQueryParams(queryParams).SetResult(&result).Get(config.KoreaExim.Path)
	if err != nil {
		return nil, err
	}

	return result, nil
}
