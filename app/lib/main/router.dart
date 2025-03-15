import "package:aerok_amos_service/ui/index.dart" as ui;
import "package:go_router/go_router.dart";
import "package:flutter/material.dart";

final class _Route {
  const _Route({required this.path, required this.widget});

  final String path;
  final Widget widget;
}

final routes = <_Route>[
  _Route(path: "/", widget: ui.HomeView()),
  _Route(path: "/exchange-rates", widget: ui.ExchangeRatesView()),
  _Route(path: "/amos-aim-web-services", widget: ui.AmosAimWebServicesView()),
];

final routerConfig = GoRouter(
  initialLocation: "/",
  routes: [
    ShellRoute(
      builder: (context, state, child) {
        return ui.Screen(child: child);
      },
      routes:
          routes.map((route) {
            return GoRoute(
              path: route.path,
              builder: (context, state) {
                return route.widget;
              },
            );
          }).toList(),
    ),
  ],
);
