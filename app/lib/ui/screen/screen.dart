import "package:aerok_amos_service/common_ui/index.dart" as common_ui;
import "package:aerok_amos_service/usecases/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:file_picker/file_picker.dart";
import "package:go_router/go_router.dart";
import "package:flutter/material.dart";
import "package:intl/intl.dart";

part "drawer.dart";
part "import_currency.dart";
part "transfer_flight_schedule.dart";

final class Screen extends StatelessWidget {
  const Screen({super.key, required this.child});

  final Widget child;

  @override
  Widget build(context) {
    return Scaffold(
      appBar: AppBar(
        notificationPredicate: (notification) {
          if (notification is OverscrollNotification) {
            return true;
          }

          return false;
        },
        actions: [_AmosServiceMenu()],
        actionsPadding: EdgeInsets.only(right: 16),
      ),
      drawer: _Drawer(),
      body: Column(
        children: [
          Expanded(child: child),
          Padding(padding: EdgeInsets.all(32), child: _Footer()),
        ],
      ),
    );
  }
}

final class _AmosServiceMenu extends StatelessWidget {
  const _AmosServiceMenu();

  @override
  Widget build(context) {
    return PopupMenuButton(
      tooltip: "AMOS Service",
      itemBuilder: (context) {
        return [
          PopupMenuItem<void>(
            child: Row(
              children: [
                Icon(Icons.monetization_on_outlined),
                SizedBox(width: 8),
                Text("Import Currency"),
              ],
            ),
            onTap: () {
              showDialog<void>(
                context: context,
                builder: (context) => const _ImportCurrencyDialog(),
              );
            },
          ),
          PopupMenuItem<void>(
            child: Row(
              children: [
                Icon(Icons.airplane_ticket_outlined),
                SizedBox(width: 8),
                Text("Transfer Future Flights"),
              ],
            ),
            onTap: () {
              showDialog<void>(
                context: context,
                builder: (context) => const _TransferFutureFlightsDialog(),
              );
            },
          ),
        ];
      },
    );
  }
}

final class _Footer extends StatelessWidget {
  const _Footer();

  @override
  Widget build(context) {
    final theme = Theme.of(context);

    return Column(
      children: [
        Text(
          "Aero_K Airlines AMOS Service",
          style: theme.textTheme.titleMedium?.copyWith(
            color: theme.colorScheme.secondary,
            fontWeight: FontWeight.w400,
          ),
        ),
        SizedBox(height: 8),
        Text(
          "© 2025 Developed by In Son",
          style: theme.textTheme.bodySmall?.copyWith(
            color: theme.colorScheme.secondary,
            fontWeight: FontWeight.w200,
          ),
        ),
      ],
    );
  }
}

final class _Dialog extends StatelessWidget {
  const _Dialog({
    required this.title,
    required this.actions,
    required this.child,
  });

  final String title;
  final List<Widget> actions;
  final Widget child;

  @override
  Widget build(context) {
    final themeData = Theme.of(context);

    return AlertDialog(
      title: Text(title, style: themeData.textTheme.titleLarge),
      contentPadding: EdgeInsets.symmetric(horizontal: 16, vertical: 32),
      content: SizedBox(width: 512, height: 256, child: child),
      actions: actions,
    );
  }
}
