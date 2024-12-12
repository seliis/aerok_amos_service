import "package:flutter_bloc/flutter_bloc.dart";

import "package:app/internal/repositories/__index.dart" as repositories;

final class UpdateAmosCurrency extends Cubit<UpdateAmosCurrencyState> {
  UpdateAmosCurrency({
    required this.amosRepository,
  }) : super(UpdateAmosCurrencyStateInitial());

  final repositories.AmosRepository amosRepository;

  Future<void> execute({
    required String authorization,
    required String date,
  }) async {
    emit(UpdateAmosCurrencyStateLoading());

    try {
      await amosRepository.updateCurrency(authorization: authorization, date: date);
      emit(UpdateAmosCurrencyStateSuccess());
    } on Exception catch (exception, stackTrace) {
      emit(UpdateAmosCurrencyStateFailure(
        message: exception.toString(),
        stackTrace: stackTrace.toString(),
      ));
    } catch (error, stackTrace) {
      emit(UpdateAmosCurrencyStateFailure(
        message: error.toString(),
        stackTrace: stackTrace.toString(),
      ));
    }
  }
}

final class UpdateAmosCurrencyState {}

final class UpdateAmosCurrencyStateInitial extends UpdateAmosCurrencyState {}

final class UpdateAmosCurrencyStateLoading extends UpdateAmosCurrencyState {}

final class UpdateAmosCurrencyStateSuccess extends UpdateAmosCurrencyState {}

final class UpdateAmosCurrencyStateFailure extends UpdateAmosCurrencyState {
  UpdateAmosCurrencyStateFailure({
    required this.message,
    required this.stackTrace,
  });

  final String message;
  final String stackTrace;
}
