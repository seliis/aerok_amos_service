import "package:flutter_bloc/flutter_bloc.dart";

import "package:app/internal/repositories/__index.dart" as repositories;
import "package:app/internal/entities/__index.dart" as entities;

final class GetExchangeRate extends Cubit<GetExchangeRateState> {
  GetExchangeRate({
    required this.currencyRepository,
  }) : super(GetExchangeRateStateInitial());

  final repositories.CurrencyRepository currencyRepository;

  void execute({
    required String currencyCode,
    required String date,
  }) async {
    emit(GetExchangeRateStateLoading());

    try {
      final exchangeRate = await currencyRepository.getExchangeRate(
        currencyCode: currencyCode,
        date: date,
      );
      emit(
        GetExchangeRateStateSuccess(
          exchangeRate: exchangeRate,
        ),
      );
    } catch (error, stackTrace) {
      emit(GetExchangeRateStateFailure(
        message: error.toString(),
        stackTrace: stackTrace.toString(),
      ));
    }
  }
}

final class GetExchangeRateState {}

final class GetExchangeRateStateInitial extends GetExchangeRateState {}

final class GetExchangeRateStateLoading extends GetExchangeRateState {}

final class GetExchangeRateStateSuccess extends GetExchangeRateState {
  GetExchangeRateStateSuccess({
    required this.exchangeRate,
  });

  final entities.ExchangeRate exchangeRate;
}

final class GetExchangeRateStateFailure extends GetExchangeRateState {
  GetExchangeRateStateFailure({
    required this.message,
    required this.stackTrace,
  });

  final String message;
  final String stackTrace;
}
