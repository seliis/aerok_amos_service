import "package:flutter_bloc/flutter_bloc.dart";

import "package:app/internal/repositories/__index.dart" as repositories;

final class RequestWebServiceKey extends Cubit<RequestWebServiceKeyState> {
  RequestWebServiceKey({
    required this.amosRepository,
  }) : super(RequestWebServiceKeyStateInitial());

  final repositories.AmosRepository amosRepository;

  Future<void> execute(String password) async {
    emit(RequestWebServiceKeyStateLoading());

    try {
      emit(
        RequestWebServiceKeyStateSuccess(
          key: await amosRepository.requestWebServiceKey(password: password),
        ),
      );
    } catch (error, stackTrace) {
      emit(RequestWebServiceKeyStateFailure(
        messsage: error.toString(),
        stackTrace: stackTrace.toString(),
      ));
    }
  }
}

final class RequestWebServiceKeyState {}

final class RequestWebServiceKeyStateInitial extends RequestWebServiceKeyState {}

final class RequestWebServiceKeyStateLoading extends RequestWebServiceKeyState {}

final class RequestWebServiceKeyStateSuccess extends RequestWebServiceKeyState {
  RequestWebServiceKeyStateSuccess({
    required this.key,
  });

  final String key;
}

final class RequestWebServiceKeyStateFailure extends RequestWebServiceKeyState {
  RequestWebServiceKeyStateFailure({
    required this.messsage,
    required this.stackTrace,
  });

  final String messsage;
  final String stackTrace;
}
