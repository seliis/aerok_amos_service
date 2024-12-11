import "package:flutter_dotenv/flutter_dotenv.dart";
import "dart:convert";

import "package:http/http.dart" as http;
import "package:web/web.dart" as web;

final class Response {
  const Response({
    required this.isSuccess,
    required this.message,
    required this.data,
  });

  final bool isSuccess;
  final String? message;
  final dynamic data;

  factory Response.fromBody(String body) {
    final json = jsonDecode(body);

    return Response(
      isSuccess: json["is_success"] as bool,
      message: json["message"] as String?,
      data: json["data"] as dynamic,
    );
  }
}

final class Client {
  const Client();

  Uri get baseUrl {
    final int port = int.parse(dotenv.env["PORT"] ?? "8080");

    return Uri.parse(web.window.location.href).replace(port: port);
  }

  Uri _getUri(String path, {Map<String, String>? queryParameters}) {
    return baseUrl.replace(
      path: "/api$path",
      queryParameters: queryParameters,
    );
  }

  Future<Response> get(String path, {Map<String, String>? queryParameters, Map<String, String>? headers}) async {
    await Future<void>.delayed(const Duration(seconds: 1)); // TODO: Remove Before Production

    final response = await http.get(_getUri(path, queryParameters: queryParameters), headers: headers);

    return Response.fromBody(response.body);
  }

  Future<Response> post(String path, {Map<String, String>? queryParameters, Map<String, dynamic>? body}) async {
    await Future<void>.delayed(const Duration(seconds: 1)); // TODO: Remove Before Production

    final response = await http.post(
      _getUri(path, queryParameters: queryParameters),
      body: jsonEncode(body),
      headers: {
        "Content-Type": "application/json",
      },
    );

    return Response.fromBody(response.body);
  }
}
