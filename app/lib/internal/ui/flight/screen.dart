import "package:flutter/material.dart";
import "package:flutter_bloc/flutter_bloc.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;

final class FutureFlightsScreen extends StatelessWidget {
  const FutureFlightsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    bool isLoading = context.watch<usecases.OpenFutureFlights>().state is usecases.OpenFutureFlightsStateLoading;

    return Scaffold(
      appBar: AppBar(
        title: const Text("Future Flights"),
      ),
      body: Center(
        child: SizedBox(
          width: 256,
          height: 48,
          child: FilledButton(
            onPressed: isLoading
                ? null
                : () {
                    context.read<usecases.OpenFutureFlights>().execute();
                  },
            style: FilledButton.styleFrom(
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(8),
              ),
            ),
            child: isLoading
                ? const common_ui.ProgressIndicator(
                    scale: 0.5,
                  )
                : const Text("Open"),
          ),
        ),
      ),
    );
  }
}
