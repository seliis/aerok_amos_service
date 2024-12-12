import "package:flutter/material.dart";

import "package:app/internal/ui/__index.dart" as ui;

enum Route {
  home(path: "/", name: "Home", icon: Icons.home),
  currency(path: "/currency", name: "Currency", icon: Icons.attach_money);

  const Route({
    required this.path,
    required this.name,
    required this.icon,
  });

  final String path;
  final String name;
  final IconData icon;
}

PageRoute<void> onGenerateRoute(RouteSettings settings) {
  if (settings.name == Route.currency.path) {
    return _getRoute(const ui.CurrencyScreen());
  }

  return _getRoute(const ui.HomeScreen());
}

PageRoute<void> _getRoute(Widget widget) {
  return NoTransition(widget: widget);
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
