import "package:aerok_amos_service/repositories/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class ImportCurrency extends Cubit<ImportCurrencyState> {
  ImportCurrency(this.amosAimWebServicesRepository)
    : super(const ImportCurrencyStateInitial());

  final AmosAimWebServicesRepository amosAimWebServicesRepository;

  Future<void> execute({required String token, required String date}) async {
    emit(const ImportCurrencyStateLoading());

    try {
      emit(
        ImportCurrencyStateSuccess(
          result: await amosAimWebServicesRepository.importCurrency(
            token,
            date,
          ),
        ),
      );
    } catch (e) {
      emit(ImportCurrencyStateFailure(message: e.toString()));
    }
  }
}

final class ImportCurrencyState {
  const ImportCurrencyState();
}

final class ImportCurrencyStateInitial extends ImportCurrencyState {
  const ImportCurrencyStateInitial();
}

final class ImportCurrencyStateLoading extends ImportCurrencyState {
  const ImportCurrencyStateLoading();
}

final class ImportCurrencyStateSuccess extends ImportCurrencyState {
  const ImportCurrencyStateSuccess({required this.result});

  final String? result;
}

final class ImportCurrencyStateFailure extends ImportCurrencyState {
  const ImportCurrencyStateFailure({required this.message});

  final String message;
}
