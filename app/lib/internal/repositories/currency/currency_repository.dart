import "package:app/internal/entities/__index.dart" as entities;
import "package:app/infra/__index.dart" as infra;

final class CurrencyRepository {
  const CurrencyRepository();

  Future<List<entities.Currency>> getCurrencies() async {
    return const infra.Client().get("/currency/currencies").then(
      (response) {
        final data = response.data as List<dynamic>;

        return data.map((currency) {
          return entities.Currency.fromMap(currency as Map<String, dynamic>);
        }).toList();
      },
    );
  }

  Future<entities.ExchangeRate> getExchangeRate({
    required String currencyCode,
    required String date,
  }) async {
    return const infra.Client().get(
      "/currency/exchangeRate",
      queryParameters: {
        "currencyCode": currencyCode,
        "date": date,
      },
    ).then(
      (response) {
        return entities.ExchangeRate.fromMap(response.data as Map<String, dynamic>);
      },
    );
  }
}
