import "package:flutter/material.dart";

import "package:app/base/router.dart";

final class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      onGenerateRoute: onGenerateRoute,
      theme: ThemeData(
        colorSchemeSeed: Colors.teal,
        fontFamily: "NanumSquareNeo",
      ),
    );
  }
}
