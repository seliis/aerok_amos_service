import "package:aerok_amos_service/repositories/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";

final class GetAuth extends Cubit<GetAuthState> {
  GetAuth() : super(GetAuthInitial());

  final 

  void execute() async {
    emit(GetAuthLoading());
    try {
      emit(GetAuthSuccess(auth: ));
    } catch (e) {
      emit(GetAuthFailure(message: e.toString()));
    }
  }
}

final class GetAuthState {}

final class GetAuthInitial extends GetAuthState {}

final class GetAuthLoading extends GetAuthState {}

final class GetAuthSuccess extends GetAuthState {
  GetAuthSuccess({required this.auth});

  final String auth;
}

final class GetAuthFailure extends GetAuthState {
  GetAuthFailure({required this.message});

  final String message;
}
