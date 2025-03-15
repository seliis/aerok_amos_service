import "package:aerok_amos_service/repositories/index.dart";
import "package:aerok_amos_service/entities/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class GetCurrencies extends Cubit<GetCurrenciesState> {
  GetCurrencies(this.exchangeRateRepository)
    : super(const GetCurrenciesStateInitial());

  final ExchangeRateRepository exchangeRateRepository;

  Future<void> execute() async {
    emit(const GetCurrenciesStateLoading());

    try {
      emit(
        GetCurrenciesStateSuccess(
          currencies: await exchangeRateRepository.getCurrencies(),
        ),
      );
    } catch (e) {
      emit(GetCurrenciesStateFailure(message: e.toString()));
    }
  }
}

final class GetCurrenciesState {
  const GetCurrenciesState();
}

final class GetCurrenciesStateInitial extends GetCurrenciesState {
  const GetCurrenciesStateInitial();
}

final class GetCurrenciesStateLoading extends GetCurrenciesState {
  const GetCurrenciesStateLoading();
}

final class GetCurrenciesStateSuccess extends GetCurrenciesState {
  const GetCurrenciesStateSuccess({required this.currencies});

  final List<Currency> currencies;
}

final class GetCurrenciesStateFailure extends GetCurrenciesState {
  const GetCurrenciesStateFailure({required this.message});

  final String message;
}
