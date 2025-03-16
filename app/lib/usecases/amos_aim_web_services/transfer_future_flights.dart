import "package:aerok_amos_service/repositories/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:file_picker/file_picker.dart";

final class TransferFutureFlights extends Cubit<TransferFutureFlightsState> {
  TransferFutureFlights(
    this.flightScheduleRepository,
    this.amosAimWebServiceRepository,
  ) : super(const TransferFutureFlightsInitial());

  final FlightScheduleRepository flightScheduleRepository;
  final AmosAimWebServicesRepository amosAimWebServiceRepository;

  Future<void> execute({
    required PlatformFile file,
    required String token,
  }) async {
    emit(const TransferFutureFlightsLoading());

    try {
      final result = await flightScheduleRepository.updateFlightSchedule(
        bytes: file.bytes!,
        fileName: file.name,
      );

      if (result != null) {
        throw Exception(result);
      }

      emit(
        TransferFutureFlightsSuccess(
          message: await amosAimWebServiceRepository.transferFutureFlights(
            token,
          ),
        ),
      );
    } catch (e) {
      emit(TransferFutureFlightsFailure(message: e.toString()));
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
  const TransferFutureFlightsSuccess({required this.message});

  final String? message;
}

final class TransferFutureFlightsFailure extends TransferFutureFlightsState {
  const TransferFutureFlightsFailure({required this.message});

  final String message;
}
