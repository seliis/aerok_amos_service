import "package:aerok_amos_service/repositories/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:file_picker/file_picker.dart";

final class TransferFutureFlights extends Cubit<TransferFutureFlightsState> {
  TransferFutureFlights(this.flightScheduleRepository, this.amosRepository)
    : super(const TransferFutureFlightsInitial());

  final FlightScheduleRepository flightScheduleRepository;
  final AmosRepository amosRepository;

  Future<void> execute({
    required PlatformFile file,
    required String password,
    required String date,
  }) async {
    emit(const TransferFutureFlightsLoading());

    try {
      final result = await flightScheduleRepository.updateFlightSchedule(
        bytes: file.bytes!,
        fileName: file.name,
        password: password,
      );

      if (result != null) {
        throw Exception(result);
      }

      emit(
        TransferFutureFlightsSuccess(
          await amosRepository.transferFutureFlights(password, date),
        ),
      );
    } catch (e) {
      emit(TransferFutureFlightsFailure(e.toString()));
    }
  }
}

final class TransferFutureFlightsState {
  const TransferFutureFlightsState();
}

final class TransferFutureFlightsInitial extends TransferFutureFlightsState {
  const TransferFutureFlightsInitial();
}

final class TransferFutureFlightsLoading extends TransferFutureFlightsState {
  const TransferFutureFlightsLoading();
}

final class TransferFutureFlightsSuccess extends TransferFutureFlightsState {
  const TransferFutureFlightsSuccess(this.message);

  final String? message;
}

final class TransferFutureFlightsFailure extends TransferFutureFlightsState {
  const TransferFutureFlightsFailure(this.message);

  final String message;
}
