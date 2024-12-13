import "dart:typed_data";

import "package:file_picker/file_picker.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class OpenFutureFlights extends Cubit<OpenFutureFlightsState> {
  OpenFutureFlights() : super(OpenFutureFlightsStateInitial());

  Future<void> execute() async {
    emit(OpenFutureFlightsStateLoading());

    try {
      final result = await FilePicker.platform.pickFiles(
        initialDirectory: "%USERPROFILE%/Downloads",
        allowMultiple: false,
        allowedExtensions: [
          "xlsx"
        ],
        type: FileType.custom,
      );

      if (result == null) {
        emit(OpenFutureFlightsStateInitial());
        return;
      }

      final file = result.files.first;

      emit(OpenFutureFlightsStateSuccess(
        name: file.name,
        data: file.bytes,
      ));
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

final class OpenFutureFlightsStateSuccess extends OpenFutureFlightsState {
  OpenFutureFlightsStateSuccess({
    required this.name,
    required this.data,
  });

  final String name;
  final Uint8List? data;
}

final class OpenFutureFlightsStateFailure extends OpenFutureFlightsState {
  OpenFutureFlightsStateFailure({
    required this.message,
    required this.stackTrace,
  });

  final String message;
  final String stackTrace;
}
