import "package:aerok_amos_service/repositories/index.dart";
import "package:aerok_amos_service/usecases/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

final class DependencyInjector extends StatelessWidget {
  const DependencyInjector(this.child, {super.key});

  final Widget child;

  @override
  Widget build(context) {
    return _Repositories(_UseCases(child));
  }
}

final class _Repositories extends StatelessWidget {
  const _Repositories(this.child);

  final Widget child;

  @override
  Widget build(context) {
    return MultiRepositoryProvider(
      providers: [
        RepositoryProvider<ExchangeRateRepository>(
          create: (context) {
            return ExchangeRateRepository();
          },
        ),
        RepositoryProvider<FlightScheduleRepository>(
          create: (context) {
            return FlightScheduleRepository();
          },
        ),
        RepositoryProvider<AmosAimWebServicesRepository>(
          create: (context) {
            return AmosAimWebServicesRepository();
          },
        ),
      ],
      child: child,
    );
  }
}

final class _UseCases extends StatelessWidget {
  const _UseCases(this.child);

  final Widget child;

  @override
  Widget build(context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider<GetCurrencies>(
          create: (context) {
            return GetCurrencies(context.read<ExchangeRateRepository>());
          },
        ),
        BlocProvider<GetExchangeRate>(
          create: (context) {
            return GetExchangeRate(context.read<ExchangeRateRepository>());
          },
        ),
        BlocProvider<GetAnnualExchangeRates>(
          create: (context) {
            return GetAnnualExchangeRates(
              context.read<ExchangeRateRepository>(),
            );
          },
        ),
        BlocProvider<GetAuth>(
          create: (context) {
            return GetAuth(context.read<AmosAimWebServicesRepository>());
          },
        ),
        BlocProvider<ImportCurrency>(
          create: (context) {
            return ImportCurrency(context.read<AmosAimWebServicesRepository>());
          },
        ),
        BlocProvider<TransferFutureFlights>(
          create: (context) {
            return TransferFutureFlights(
              context.read<FlightScheduleRepository>(),
              context.read<AmosAimWebServicesRepository>(),
            );
          },
        ),
      ],
      child: child,
    );
  }
}
