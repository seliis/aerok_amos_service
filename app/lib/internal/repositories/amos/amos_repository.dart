import "dart:convert";
import "dart:typed_data";

import "package:app/infra/__index.dart" as infra;

final class AmosRepository {
  const AmosRepository();

  Future<String> requestWebServiceKey({
    required String password,
  }) async {
    return const infra.Client().get(
      "/amos/auth",
      headers: {
        "Authorization": "Basic ${base64.encode(utf8.encode("admin:$password"))}",
      },
    ).then(
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
    return const infra.Client().post(
      "/amos/updateCurrency/$date",
      headers: {
        "Authorization": "Basic $authorization",
      },
    ).then(
      (response) {
        if (!response.isSuccess) {
          throw Exception(response.message);
        }
      },
    );
  }

  Future<void> updateFutureFlights({
    required String authorization,
    required Uint8List? data,
  }) async {
    return const infra.Client().post(
      "/amos/updateFutureFlights",
      headers: {
        "Authorization": "Basic $authorization",
      },
      body: {
        "data": data,
      },
    ).then(
      (response) {
        if (!response.isSuccess) {
          throw Exception(response.message);
        }
      },
    );
  }
}
