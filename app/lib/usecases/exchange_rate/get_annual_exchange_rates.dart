import "package:aerok_amos_service/repositories/index.dart";
import "package:aerok_amos_service/entities/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class GetAnnualExchangeRates extends Cubit<GetAnnualExchangeRatesState> {
  GetAnnualExchangeRates(this.exchangeRateRepository)
    : super(const GetAnnualExchangeRatesStateInitial());

  final ExchangeRateRepository exchangeRateRepository;

  Future<void> execute({required String code, required String year}) async {
    emit(const GetAnnualExchangeRatesStateLoading());

    try {
      emit(
        GetAnnualExchangeRatesStateSuccess(
          exchangeRates: await exchangeRateRepository.getAnnualExchangeRates(
            code: code,
            year: year,
          ),
        ),
      );
    } catch (e) {
      emit(GetAnnualExchangeRatesStateFailure(message: e.toString()));
    }
  }
}

final class GetAnnualExchangeRatesState {
  const GetAnnualExchangeRatesState();
}

final class GetAnnualExchangeRatesStateInitial
    extends GetAnnualExchangeRatesState {
  const GetAnnualExchangeRatesStateInitial();
}

final class GetAnnualExchangeRatesStateLoading
    extends GetAnnualExchangeRatesState {
  const GetAnnualExchangeRatesStateLoading();
}

final class GetAnnualExchangeRatesStateSuccess
    extends GetAnnualExchangeRatesState {
  const GetAnnualExchangeRatesStateSuccess({required this.exchangeRates});

  final List<ExchangeRateWithCurrency> exchangeRates;
}

final class GetAnnualExchangeRatesStateFailure
    extends GetAnnualExchangeRatesState {
  const GetAnnualExchangeRatesStateFailure({required this.message});

  final String message;
}
