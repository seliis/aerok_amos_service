import "package:flutter/services.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/presenters/__index.dart" as presenters;
import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;
import "package:app/internal/entities/__index.dart" as entities;

part "container.dart";
part "dialog.dart";

final class CurrencyScreen extends StatelessWidget {
  const CurrencyScreen({super.key});

  @override
  Widget build(context) {
    return BlocListener(
      bloc: BlocProvider.of<usecases.GetExchangeRate>(context),
      listener: (context, state) {
        if (state is usecases.GetExchangeRateStateError) {
          showDialog<void>(
            context: context,
            builder: (context) {
              return common_ui.ErrorDialog(
                message: state.error.toString(),
                stackTrace: state.stackTrace,
              );
            },
          );
        }

        if (state is usecases.GetExchangeRateStateSuccess) {
          showDialog<void>(
            context: context,
            builder: (context) {
              return _Dialog(
                exchangeRate: state.exchangeRate,
                width: MediaQuery.of(context).size.width * 0.75,
                height: MediaQuery.of(context).size.height * 0.25,
              );
            },
          );
        }
      },
      child: BlocBuilder<usecases.GetCurrencies, usecases.GetCurrenciesState>(
        bloc: BlocProvider.of<usecases.GetCurrencies>(context)..execute(),
        builder: (context, state) {
          if (state is usecases.GetCurrenciesStateLoading) {
            return const Scaffold(
              body: Center(
                child: CircularProgressIndicator(),
              ),
            );
          }

          if (state is usecases.GetCurrenciesStateSuccess) {
            return _View(state: state);
          }

          if (state is usecases.GetCurrenciesStateError) {
            return Column(
              children: [
                Text(state.error.toString()),
                Text(state.stackTrace.toString()),
              ],
            );
          }

          return const SizedBox();
        },
      ),
    );
  }
}

final class _View extends StatelessWidget {
  const _View({
    required this.state,
  });

  final usecases.GetCurrenciesStateSuccess state;

  @override
  Widget build(context) {
    final screenWidth = MediaQuery.of(context).size.width;

    return Scaffold(
      body: Padding(
        padding: EdgeInsets.symmetric(
          horizontal: screenWidth * 0.10,
          vertical: screenWidth * 0.025,
        ),
        child: const _Body(),
      ),
    );
  }
}

final class _Body extends StatelessWidget {
  const _Body();

  @override
  Widget build(context) {
    return const Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Spacer(),
        _Container(),
        Spacer(),
        _Footer(),
      ],
    );
  }
}

final class _Footer extends StatelessWidget {
  const _Footer();

  @override
  Widget build(context) {
    return const Text(
      "© 2024 Aero K Airlines, Developed by In Son",
      textAlign: TextAlign.center,
      style: TextStyle(
        fontWeight: FontWeight.w100,
        fontSize: 12,
      ),
    );
  }
}
