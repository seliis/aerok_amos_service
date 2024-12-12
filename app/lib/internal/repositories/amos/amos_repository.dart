import "dart:convert";

import "package:app/infra/__index.dart" as infra;

final class AmosRepository {
  const AmosRepository();

  Future<String> requestWebServiceKey({
    required String password,
  }) async {
    return const infra.Client().get("/amos/auth", headers: {
      "Authorization": "Basic ${base64.encode(utf8.encode("admin:$password"))}",
    }).then(
      (response) {
        if (!response.isSuccess) {
          throw Exception(response.message);
        }

        return response.data as String;
      },
    );
  }

  Future<void> updateCurrency({
    required String authorization,
    required String date,
  }) async {
    return const infra.Client().post("/amos/updateCurrency/$date", headers: {
      "Authorization": "Basic $authorization",
    }).then(
      (response) {
        if (!response.isSuccess) {
          throw Exception(response.message);
        }
      },
    );
  }
}
