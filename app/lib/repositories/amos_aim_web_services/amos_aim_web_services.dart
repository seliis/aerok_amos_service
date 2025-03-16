import "package:aerok_amos_service/net/index.dart";
import "dart:convert";

final class AmosAimWebServicesRepository {
  Future<String> getToken(String password) async {
    final response = await Http.request(
      method: HttpMethod.post,
      path: "/amos-aim-webservice/authorize",
      headers: {
        "Authorization":
            "Basic ${base64.encode(utf8.encode("admin:$password"))}",
      },
    );

    if (response.data == null) {
      throw Exception("No Basic Token Received From Server");
    }

    return response.data! as String;
  }

  Future<String?> importCurrency(String token, String date) async {
    final response = await Http.request(
      method: HttpMethod.post,
      path: "/amos-aim-webservice/import-currency?date=$date",
      headers: {"Authorization": "Basic $token"},
    );

    return response.message;
  }

  Future<String?> transferFutureFlights(String token) async {
    final response = await Http.request(
      method: HttpMethod.post,
      path: "/amos-aim-webservice/transfer-future-flights",
      headers: {"Authorization": "Basic $token"},
    );

    return response.message;
  }
}
