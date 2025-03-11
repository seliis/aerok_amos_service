import "dart:convert";

import "package:flutter_dotenv/flutter_dotenv.dart";
import "package:http/http.dart" as http;
import "package:flutter/services.dart";
import "package:web/web.dart" as web;

part "response.dart";

enum HttpMethod { get, post }

final class ServerException implements Exception {
  const ServerException({
    required this.statusCode,
    this.reasonPhrase = "NO_REASON_PHRASE",
    this.message = "NO_MESSAGE",
  });

  final int statusCode;
  final String? reasonPhrase;
  final String? message;

  @override
  String toString() {
    return "[$statusCode $reasonPhrase]: $message)";
  }
}

final class Http {
  static Uri target = Uri.parse(
    "${web.window.location.protocol}//${web.window.location.hostname}:${dotenv.env["PORT"]}/api",
  );

  static Future<Response> request({
    required HttpMethod method,
    required String path,
    Map<String, String>? headers,
    Map<String, dynamic>? body,
  }) async {
    late http.Response primitive;

    switch (method) {
      case HttpMethod.get:
        primitive = await http.get(Uri.parse("$target$path"), headers: headers);
        break;
      case HttpMethod.post:
        primitive = await http.post(
          Uri.parse("$target$path"),
          headers: headers,
          body: jsonEncode(body),
        );
        break;
    }

    final response = Response.fromPrimitive(primitive);

    if (response.isOk) {
      return response;
    } else {
      throw ServerException(
        statusCode: response.statusCode,
        reasonPhrase: response.reasonPhrase,
        message: response.message,
      );
    }
  }

  static Future<Response> multipart({
    required String path,
    required Uint8List bytes,
    required String fieldName,
    required String fileName,
    Map<String, String>? headers,
  }) async {
    final request = http.MultipartRequest("POST", Uri.parse("$target$path"));

    request.files.add(
      http.MultipartFile.fromBytes(fieldName, bytes, filename: fileName),
    );

    request.headers.addAll(headers ?? {});

    final response = Response.fromPrimitive(
      await http.Response.fromStream(await request.send()),
    );

    if (response.isOk) {
      return response;
    } else {
      throw ServerException(
        statusCode: response.statusCode,
        reasonPhrase: response.reasonPhrase,
        message: response.message,
      );
    }
  }
}
