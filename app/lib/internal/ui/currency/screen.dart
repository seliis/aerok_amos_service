import "package:flutter/services.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/presenters/__index.dart" as presenters;
import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;
import "package:app/internal/entities/__index.dart" as entities;

part "exchange_rate_dialog.dart";
part "components.dart";
part "listeners.dart";

final class CurrencyScreen extends StatelessWidget {
  const CurrencyScreen({super.key});

  @override
  Widget build(context) {
    return MultiBlocListener(
      listeners: [
        _GetCurrenciesListener(),
        _GetExchangeRateListener(),
      ],
      child: BlocBuilder(
        bloc: BlocProvider.of<usecases.GetCurrencies>(context)..execute(),
        builder: (context, state) {
          if (state is usecases.GetCurrenciesStateLoading) {
            return const Scaffold(
              body: Center(
                child: common_ui.ProgressIndicator(),
              ),
            );
          }

          if (state is usecases.GetCurrenciesStateSuccess) {
            return const _View();
          }

          return const SizedBox.shrink();
        },
      ),
    );
  }
}

final class _View extends StatelessWidget {
  const _View();

  @override
  Widget build(context) {
    return const Column(
      children: [
        _DropdownMenu(),
        SizedBox(
          height: 16,
        ),
        _DateForm(),
        SizedBox(
          height: 16,
        ),
        _QueryButton(),
      ],
    );
  }
}
