import "package:aerok_amos_service/net/index.dart";
import "dart:convert";

final class AmosAimWebServicesRepository {
  Future<String?> getAuth(String id, String password) async {
    final response = await Http.request(
      method: HttpMethod.post,
      path: "amos-aim-webservice/auth",
      headers: {
        "Authorization": "Basic ${base64.encode(utf8.encode("$id:$password"))}",
      },
    );

    return response.message;
  }
}
