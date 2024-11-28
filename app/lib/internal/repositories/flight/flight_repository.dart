import "dart:typed_data";

import "package:app/infra/__index.dart" as infra;

final class FlightRepository {
  const FlightRepository();

  Future<void> updateAmosFutureFlights({
    required String id,
    required String password,
    required Uint8List? data,
  }) async {
    return const infra.Client().post("/flight/updateAmos", body: {
      "web_service_auth": {
        "id": id,
        "password": password,
      },
      "data": data,
    }).then(
      (response) {
        if (!response.isSuccess) {
          throw Exception(response.message);
        }
      },
    );
  }
}
