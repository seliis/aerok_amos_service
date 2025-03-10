import "package:aerok_amos_service/repositories/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class ImportCurrency extends Cubit<ImportCurrencyState> {
  ImportCurrency(this.amosRepository)
    : super(const ImportCurrencyStateInitial());

  final AmosRepository amosRepository;

  Future<void> execute(String password, String date) async {
    emit(const ImportCurrencyStateLoading());

    try {
      emit(
        ImportCurrencyStateSuccess(
          await amosRepository.importCurrency(password, date),
        ),
      );
    } catch (e) {
      emit(ImportCurrencyStateFailure(e.toString()));
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
  const ImportCurrencyStateSuccess(this.result);

  final String? result;
}

final class ImportCurrencyStateFailure extends ImportCurrencyState {
  const ImportCurrencyStateFailure(this.message);

  final String message;
}
