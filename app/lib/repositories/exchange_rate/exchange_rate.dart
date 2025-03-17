import "package:aerok_amos_service/entities/index.dart";
import "package:aerok_amos_service/net/index.dart";

final class ExchangeRateRepository {
  Future<List<Currency>> getCurrencies() async {
    return await Http.request(method: HttpMethod.get, path: "/currency/").then((
      response,
    ) {
      if (response.data == null) {
        throw Exception("Currencies Not Found");
      }

      return (response.data as List<dynamic>).map((e) {
        return Currency.fromJson(e as Map<String, dynamic>);
      }).toList();
    });
  }

  Future<ExchangeRateWithCurrency> getExchangeRate({
    required String code,
    required String date,
  }) async {
    try {
      return await Http.request(
        method: HttpMethod.get,
        path: "/exchange-rate/currency?code=$code&date=$date",
      ).then((response) {
        return ExchangeRateWithCurrency.fromJson(
          response.data as Map<String, dynamic>,
        );
      });
    } catch (e) {
      rethrow;
    }
  }

  Future<List<ExchangeRateWithCurrency>> getAnnualExchangeRates({
    required String code,
    required String year,
  }) async {
    try {
      return await Http.request(
        method: HttpMethod.get,
        path: "/exchange-rate/annual?code=$code&year=$year",
      ).then((response) {
        return (response.data as List<dynamic>).map((e) {
          return ExchangeRateWithCurrency.fromJson(e as Map<String, dynamic>);
        }).toList();
      });
    } catch (e) {
      rethrow;
    }
  }
}
