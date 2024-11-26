import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";
import "package:flutter/services.dart";

import "package:app/internal/presenters/__index.dart" as presenters;
import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;
import "package:app/internal/entities/__index.dart" as entities;

part "exchange_rate_dialog.dart";
part "amos_auth_dialog.dart";
part "container.dart";

final class CurrencyScreen extends StatelessWidget {
  const CurrencyScreen({super.key});

  @override
  Widget build(context) {
    return BlocListener(
      bloc: BlocProvider.of<usecases.GetExchangeRate>(context),
      listener: (context, state) {
        if (state is usecases.GetExchangeRateStateFailure) {
          showDialog<void>(
            context: context,
            builder: (context) {
              return common_ui.ErrorDialog(
                message: state.message,
                stackTrace: state.stackTrace,
                width: 640,
                height: 480,
              );
            },
          );
        }

        if (state is usecases.GetExchangeRateStateSuccess) {
          showDialog<void>(
            context: context,
            builder: (context) {
              return _ExchangeRateDialog(
                exchangeRate: state.exchangeRate,
                width: 640,
                height: 480,
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
                child: common_ui.ProgressIndicator(),
              ),
            );
          }

          if (state is usecases.GetCurrenciesStateSuccess) {
            return _View(state: state);
          }

          if (state is usecases.GetCurrenciesStateFailure) {
            return Column(
              children: [
                Text(state.messsage),
                Text(state.stackTrace),
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
    return const Scaffold(
      appBar: _AppBar(),
      body: Padding(
        padding: EdgeInsets.all(16),
        child: Center(
          child: _Body(),
        ),
      ),
    );
  }
}

final class _AppBar extends StatelessWidget implements PreferredSizeWidget {
  const _AppBar();

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  Widget build(context) {
    return AppBar(
      title: const Text("Currency Rates"),
      actions: [
        PopupMenuButton(
          itemBuilder: (context) {
            return [
              PopupMenuItem<void>(
                child: const Text("Update AMOS Currency"),
                onTap: () {
                  showDialog<void>(
                    context: context,
                    builder: (context) {
                      return const _AmosAuthDialog(
                        width: 640,
                        height: 480,
                      );
                    },
                  );
                },
              ),
            ];
          },
        ),
        const SizedBox(
          width: 8,
        ),
      ],
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
        SizedBox(
          width: 512,
          child: _Container(),
        ),
      ],
    );
  }
}
