part of "http.dart";

final class Response {
  const Response({
    required this.isOk,
    required this.statusCode,
    this.reasonPhrase,
    this.message,
    this.data,
  });

  final bool isOk;
  final int statusCode;
  final String? reasonPhrase;
  final String? message;
  final dynamic data;

  factory Response.fromPrimitive(http.Response response) {
    final body = jsonDecode(response.body);

    return Response(
      isOk: body["is_ok"] as bool,
      statusCode: response.statusCode,
      reasonPhrase: response.reasonPhrase,
      message: body["message"] as String?,
      data: body["data"] as dynamic,
    );
  }
}
