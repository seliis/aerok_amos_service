import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;

part "update_amos_future_flights_dialog.dart";

final class FutureFlightScreen extends StatelessWidget {
  const FutureFlightScreen({super.key});

  @override
  Widget build(context) {
    return Scaffold(
      appBar: common_ui.MasterAppBar(
        title: "Future Flights",
        popupMenuEntries: [
          PopupMenuItem(
            onTap: () {
              showDialog<void>(
                context: context,
                builder: (context) {
                  return const _UpdateAmosFutureFlightsDialog();
                },
              );
            },
            child: const Text("Update AMOS"),
          ),
        ],
      ),
      drawer: const common_ui.MasterDrawer(),
      body: const _Body(),
    );
  }
}

final class _Body extends StatelessWidget {
  const _Body();

  @override
  Widget build(context) {
    return const Center(
      child: Text("Future Flights"),
    );
  }
}
