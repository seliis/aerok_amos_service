import "dart:typed_data";

import "package:flutter_bloc/flutter_bloc.dart";

import "package:app/internal/repositories/__index.dart" as repositories;

final class UpdateAmosFutureFlights extends Cubit<UpdateAmosFutureFlightsState> {
  UpdateAmosFutureFlights({
    required this.amosRepository,
  }) : super(UpdateAmosFutureFlightsStateInitial());

  final repositories.AmosRepository amosRepository;

  Future<void> execute({
    required String authorization,
    required Uint8List? data,
  }) async {
    emit(UpdateAmosFutureFlightsStateLoading());

    try {
      await amosRepository.updateFutureFlights(
        authorization: authorization,
        data: data,
      );

      emit(UpdateAmosFutureFlightsStateSuccess());
    } on Exception catch (exception, stackTrace) {
      emit(UpdateAmosFutureFlightsStateFailure(
        message: exception.toString(),
        stackTrace: stackTrace.toString(),
      ));
    } catch (error, stackTrace) {
      emit(UpdateAmosFutureFlightsStateFailure(
        message: error.toString(),
        stackTrace: stackTrace.toString(),
      ));
    }
  }
}

final class UpdateAmosFutureFlightsState {}

final class UpdateAmosFutureFlightsStateInitial extends UpdateAmosFutureFlightsState {}

final class UpdateAmosFutureFlightsStateLoading extends UpdateAmosFutureFlightsState {}

final class UpdateAmosFutureFlightsStateSuccess extends UpdateAmosFutureFlightsState {}

final class UpdateAmosFutureFlightsStateFailure extends UpdateAmosFutureFlightsState {
  UpdateAmosFutureFlightsStateFailure({
    required this.message,
    required this.stackTrace,
  });

  final String message;
  final String stackTrace;
}
