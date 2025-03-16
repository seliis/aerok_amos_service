import "package:aerok_amos_service/repositories/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class GetAuth extends Cubit<GetAuthState> {
  GetAuth(this.amosAimWebServicesRepository) : super(GetAuthInitial());

  final AmosAimWebServicesRepository amosAimWebServicesRepository;

  void reset() {
    emit(GetAuthInitial());
  }

  void execute(String password) async {
    emit(GetAuthLoading());
    try {
      emit(
        GetAuthSuccess(
          token: await amosAimWebServicesRepository.getToken(password),
        ),
      );
    } catch (e) {
      emit(GetAuthFailure(message: e.toString()));
    }
  }
}

final class GetAuthState {}

final class GetAuthInitial extends GetAuthState {}

final class GetAuthLoading extends GetAuthState {}

final class GetAuthSuccess extends GetAuthState {
  GetAuthSuccess({required this.token});

  final String token;
}

final class GetAuthFailure extends GetAuthState {
  GetAuthFailure({required this.message});

  final String message;
}
