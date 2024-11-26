import "package:flutter_bloc/flutter_bloc.dart";
import "package:file_picker/file_picker.dart";

import "package:app/internal/repositories/__index.dart" as repositories;

final class OpenFutureFlights extends Cubit<OpenFutureFlightsState> {
  OpenFutureFlights({
    required this.flightRepository,
  }) : super(OpenFutureFlightsStateInitial());

  final repositories.FlightRepository flightRepository;

  void execute() async {
    emit(OpenFutureFlightsStateLoading());

    try {
      final result = await FilePicker.platform.pickFiles(
        allowedExtensions: [
          "csv"
        ],
        type: FileType.custom,
      );
      print(result?.files.first.bytes);
      emit(OpenFutureFlightsStateSuccess());
    } on Exception catch (exception, stackTrace) {
      emit(OpenFutureFlightsStateFailure(
        message: exception.toString(),
        stackTrace: stackTrace.toString(),
      ));
    } catch (error, stackTrace) {
      emit(OpenFutureFlightsStateFailure(
        message: error.toString(),
        stackTrace: stackTrace.toString(),
      ));
    }
  }
}

final class OpenFutureFlightsState {}

final class OpenFutureFlightsStateInitial extends OpenFutureFlightsState {}

final class OpenFutureFlightsStateLoading extends OpenFutureFlightsState {}

final class OpenFutureFlightsStateSuccess extends OpenFutureFlightsState {}

final class OpenFutureFlightsStateFailure extends OpenFutureFlightsState {
  OpenFutureFlightsStateFailure({
    required this.message,
    required this.stackTrace,
  });

  final String message;
  final String stackTrace;
}
