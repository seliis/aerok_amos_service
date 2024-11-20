import "package:flutter/material.dart";

import "package:app/internal/ui/__index.dart" as ui;

MaterialPageRoute<void> onGenerateRoute(RouteSettings settings) {
  switch (settings.name) {
    default:
      return _getRoute(const ui.CurrencyScreen());
  }
}

MaterialPageRoute<void> _getRoute(Widget screen) {
  return MaterialPageRoute(builder: (_) => screen);
}
