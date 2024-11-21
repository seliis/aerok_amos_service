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
    final size = MediaQuery.of(context).size;

    return BlocListener(
      bloc: BlocProvider.of<usecases.GetExchangeRate>(context),
      listener: (context, state) {
        final width = size.width * 0.75;
        final height = size.height * 0.25;

        if (state is usecases.GetExchangeRateStateFailure) {
          showDialog<void>(
            context: context,
            builder: (context) {
              return common_ui.ErrorDialog(
                message: state.message,
                stackTrace: state.stackTrace,
                width: width,
                height: height,
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
                width: width,
                height: height,
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
    final screenWidth = MediaQuery.of(context).size.width;

    return Scaffold(
      appBar: const _AppBar(),
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

final class _AppBar extends StatelessWidget implements PreferredSizeWidget {
  const _AppBar();

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  Widget build(context) {
    return AppBar(
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
                      final size = MediaQuery.of(context).size;

                      return _AmosAuthDialog(
                        width: size.width * 0.75,
                        height: size.height * 0.25,
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
