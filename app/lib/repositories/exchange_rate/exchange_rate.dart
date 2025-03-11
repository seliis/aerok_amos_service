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

  Future<ExchangeRateWithCurrency> getExchangeRate(
    String code,
    String date,
  ) async {
    return await Http.request(
      method: HttpMethod.get,
      path: "/exchange-rate/currency?code=$code&date=$date",
    ).then((response) {
      if (response.data == null) {
        throw Exception("Data Not Found");
      }

      return ExchangeRateWithCurrency.fromJson(
        response.data as Map<String, dynamic>,
      );
    });
  }
}
