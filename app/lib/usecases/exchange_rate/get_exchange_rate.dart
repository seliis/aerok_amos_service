import "package:aerok_amos_service/repositories/index.dart";
import "package:aerok_amos_service/entities/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class GetExchangeRate extends Cubit<GetExchangeRateState> {
  GetExchangeRate(this.exchangeRateRepository)
    : super(const GetExchangeRateStateInitial());

  final ExchangeRateRepository exchangeRateRepository;

  Future<void> execute({required String code, required String date}) async {
    emit(const GetExchangeRateStateLoading());

    try {
      emit(
        GetExchangeRateStateSuccess(
          await exchangeRateRepository.getExchangeRate(code, date),
        ),
      );
    } catch (e) {
      emit(GetExchangeRateStateFailure(e.toString()));
    }
  }
}

final class GetExchangeRateState {
  const GetExchangeRateState();
}

final class GetExchangeRateStateInitial extends GetExchangeRateState {
  const GetExchangeRateStateInitial();
}

final class GetExchangeRateStateLoading extends GetExchangeRateState {
  const GetExchangeRateStateLoading();
}

final class GetExchangeRateStateSuccess extends GetExchangeRateState {
  const GetExchangeRateStateSuccess(this.exchangeRate);

  final ExchangeRateWithCurrency exchangeRate;
}

final class GetExchangeRateStateFailure extends GetExchangeRateState {
  const GetExchangeRateStateFailure(this.message);

  final String message;
}
