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
        RepositoryProvider<repositories.FlightRepository>(
          create: (context) => const repositories.FlightRepository(),
        ),
        RepositoryProvider<repositories.AmosRepository>(
          create: (context) => const repositories.AmosRepository(),
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
          BlocProvider<usecases.UpdateAmosCurrency>(
            create: (context) => usecases.UpdateAmosCurrency(
              currencyRepository: RepositoryProvider.of<repositories.CurrencyRepository>(context),
            ),
          ),
          BlocProvider<usecases.OpenFutureFlights>(
            create: (context) => usecases.OpenFutureFlights(
              flightRepository: RepositoryProvider.of<repositories.FlightRepository>(context),
            ),
          ),
          BlocProvider<usecases.RequestWebServiceKey>(
            create: (context) => usecases.RequestWebServiceKey(
              amosRepository: RepositoryProvider.of<repositories.AmosRepository>(context),
            ),
          ),
        ],
        child: app,
      ),
    );
  }
}
