import "package:flutter_dotenv/flutter_dotenv.dart";
import "package:flutter/material.dart";

import "package:app/base/__index.dart" as base;

Future<void> main() async {
  await dotenv.load(fileName: ".env");

  runApp(
    const base.DependencyInjector(
      app: base.App(),
    ),
  );
}
