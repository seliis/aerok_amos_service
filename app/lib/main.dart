import "dart:convert";

import "package:aerok_amos_service/main/router.dart";
import "package:flutter_dotenv/flutter_dotenv.dart";
import "package:aerok_amos_service/main/di.dart";
import "package:json_theme/json_theme.dart";
import "package:flutter/material.dart";
import "package:flutter/services.dart";

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await dotenv.load(fileName: "assets/.env");

  final themeData = jsonDecode(
    await rootBundle.loadString("assets/themes/aerok.json"),
  );

  runApp(
    DependencyInjector(App(theme: ThemeDecoder.decodeThemeData(themeData))),
  );
}

final class App extends StatelessWidget {
  const App({super.key, required this.theme});

  final ThemeData? theme;

  @override
  Widget build(context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      routerConfig: routerConfig,
      theme: theme,
    );
  }
}
