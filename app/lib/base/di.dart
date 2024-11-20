import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/repositories/__index.dart" as repositories;
import "package:app/internal/usecases/__index.dart" as usecases;

final class DependencyInjector extends StatelessWidget {
  const DependencyInjector({
    super.key,
    required this.app,
  });

  final Widget app;

  @override
  Widget build(context) {
    return MultiRepositoryProvider(
      providers: [
        RepositoryProvider<repositories.CurrencyRepository>(
          create: (context) => const repositories.CurrencyRepository(),
        ),
      ],
      child: MultiBlocProvider(
        providers: [
          BlocProvider<usecases.GetCurrencies>(
            create: (context) => usecases.GetCurrencies(
              currencyRepository: RepositoryProvider.of<repositories.CurrencyRepository>(context),
            ),
          ),
          BlocProvider<usecases.GetExchangeRate>(
            create: (context) => usecases.GetExchangeRate(
              currencyRepository: RepositoryProvider.of<repositories.CurrencyRepository>(context),
            ),
          ),
        ],
        child: app,
      ),
    );
  }
}