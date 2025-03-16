import "package:aerok_amos_service/net/index.dart";
import "package:flutter/services.dart";

final class FlightScheduleRepository {
  Future<String?> updateFlightSchedule({
    required Uint8List bytes,
    required String fileName,
  }) async {
    final response = await Http.multipart(
      path: "/flight-schedule/update",
      bytes: bytes,
      fieldName: "flight_schedule",
      fileName: fileName,
    );

    return response.message;
  }
}
