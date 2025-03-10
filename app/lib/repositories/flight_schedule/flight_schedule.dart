import "package:aerok_amos_service/net/index.dart";
import "package:flutter/services.dart";
import "dart:convert";

final class FlightScheduleRepository {
  Future<String?> updateFlightSchedule({
    required Uint8List bytes,
    required String fileName,
    required String password,
  }) async {
    final response = await Http.multipart(
      path: "/flight-schedule/update",
      bytes: bytes,
      fieldName: "flight_schedule",
      fileName: fileName,
      headers: {
        "Authorization":
            "Basic ${base64.encode(utf8.encode("admin:$password"))}",
      },
    );

    return response.message;
  }
}
