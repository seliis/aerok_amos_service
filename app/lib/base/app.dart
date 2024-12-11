import "package:flutter/material.dart";

import "package:app/base/router.dart" as router;

final class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      onGenerateRoute: router.onGenerateRoute,
      initialRoute: router.Route.currency.path,
      theme: ThemeData(
        colorSchemeSeed: Colors.teal,
        fontFamily: "NanumSquareNeo",
      ),
    );
  }
}
