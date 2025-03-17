import "package:flutter_dotenv/flutter_dotenv.dart";
import "package:go_router/go_router.dart";
import "package:flutter/material.dart";

part "drawer.dart";

final class Screen extends StatelessWidget {
  const Screen({super.key, required this.child});

  final Widget child;

  @override
  Widget build(context) {
    final theme = Theme.of(context);

    return Scaffold(
      appBar: AppBar(
        title: InkWell(
          onTap: () {
            context.go("/");
          },
          child: Text(
            "Aero K Airlines AMOS Service",
            style: theme.textTheme.titleMedium?.copyWith(
              fontWeight: FontWeight.w400,
            ),
          ),
        ),
        notificationPredicate: (notification) {
          if (notification is OverscrollNotification) {
            return true;
          }

          return false;
        },
        actionsPadding: EdgeInsets.only(right: 16),
      ),
      drawer: _Drawer(),
      body: child,
    );
  }
}
