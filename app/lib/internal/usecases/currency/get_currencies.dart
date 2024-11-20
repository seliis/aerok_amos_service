import "package:flutter_bloc/flutter_bloc.dart";

import "package:app/internal/repositories/__index.dart" as repositories;
import "package:app/internal/entities/__index.dart" as entities;

final class GetCurrencies extends Cubit<GetCurrenciesState> {
  GetCurrencies({
    required this.currencyRepository,
  }) : super(GetCurrenciesStateInitial());

  final repositories.CurrencyRepository currencyRepository;

  void execute() async {
    emit(GetCurrenciesStateLoading());

    try {
      final currencies = await currencyRepository.getCurrencies();
      emit(
        GetCurrenciesStateSuccess(
          currencies: currencies,
          currentCurrency: currencies.first,
        ),
      );
    } catch (error, stackTrace) {
      emit(GetCurrenciesStateError(
        error: error,
        stackTrace: stackTrace,
      ));
    }
  }

  void changeCurrency(entities.Currency currency) {
    if (state is GetCurrenciesStateSuccess) {
      emit(
        (state as GetCurrenciesStateSuccess).copyWith(
          currentCurrency: currency,
        ),
      );
    }
  }
}

final class GetCurrenciesState {}

final class GetCurrenciesStateInitial extends GetCurrenciesState {}

final class GetCurrenciesStateLoading extends GetCurrenciesState {}

final class GetCurrenciesStateSuccess extends GetCurrenciesState {
  GetCurrenciesStateSuccess({
    required this.currencies,
    required this.currentCurrency,
  });

  final List<entities.Currency> currencies;
  final entities.Currency currentCurrency;

  GetCurrenciesStateSuccess copyWith({
    List<entities.Currency>? currencies,
    entities.Currency? currentCurrency,
  }) {
    return GetCurrenciesStateSuccess(
      currencies: currencies ?? this.currencies,
      currentCurrency: currentCurrency ?? this.currentCurrency,
    );
  }
}

final class GetCurrenciesStateError extends GetCurrenciesState {
  GetCurrenciesStateError({
    required this.error,
    required this.stackTrace,
  });

  final Object error;
  final StackTrace stackTrace;
}
