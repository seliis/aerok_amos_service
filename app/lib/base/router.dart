import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/ui/__index.dart" as ui;

enum Route {
  home(path: "/", name: "Home"),
  currency(path: "/currency", name: "Currency");

  const Route({
    required this.path,
    required this.name,
  });

  final String path;
  final String name;
}

PageRoute<void> onGenerateRoute(RouteSettings settings) {
  if (settings.name == Route.currency.path) {
    return _getRoute(const ui.CurrencyScreen());
  }

  return _getRoute(const ui.HomeScreen());
}

PageRoute<void> _getRoute(Widget body) {
  return NoTransition(
    widget: Scaffold(
      appBar: const common_ui.MasterAppBar(),
      drawer: const common_ui.MasterDrawer(),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: body,
      ),
    ),
  );
}

final class NoTransition extends PageRouteBuilder<void> {
  NoTransition({
    required this.widget,
  }) : super(
          pageBuilder: (context, animation, secondaryAnimation) => widget,
          transitionDuration: Duration.zero,
          reverseTransitionDuration: Duration.zero,
          maintainState: true,
        );

  final Widget widget;

  @override
  Widget buildTransitions(
    context,
    animation,
    secondaryAnimation,
    child,
  ) {
    return child;
  }
}
