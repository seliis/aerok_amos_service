import "package:aerok_amos_service/ui/index.dart" as ui;
import "package:go_router/go_router.dart";
import "package:flutter/material.dart";

final class _Route {
  const _Route(this.path, this.name, this.widget, this.iconData);

  final String path;
  final String name;
  final Widget widget;
  final IconData iconData;
}

final routes = <_Route>[_Route("/", "Home", ui.Home(), Icons.home)];

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
