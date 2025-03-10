import "package:aerok_amos_service/net/index.dart";
import "dart:convert";

final class AmosRepository {
  static Map<String, String>? getHeader(String password) {
    return {
      "Authorization": "Basic ${base64.encode(utf8.encode("admin:$password"))}",
    };
  }

  Future<String?> importCurrency(String password, String date) async {
    final response = await Http.request(
      method: HttpMethod.post,
      path: "/amos-aim-webservice/import-currency?date=$date",
      headers: getHeader(password),
    );

    return response.message;
  }

  Future<String?> transferFutureFlights(String password, String date) async {
    final response = await Http.request(
      method: HttpMethod.post,
      path: "/amos-aim-webservice/transfer-future-flights?date=$date",
      headers: getHeader(password),
    );

    return response.message;
  }
}
