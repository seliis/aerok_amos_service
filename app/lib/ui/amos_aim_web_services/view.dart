import "package:aerok_amos_service/common_ui/index.dart" as common_ui;
import "package:aerok_amos_service/usecases/index.dart";
import "package:file_picker/file_picker.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";
import "package:intl/intl.dart";

part "body.dart";
part "sign_in.dart";
part "import_currency.dart";
part "transfer_future_flights.dart";

final class AmosAimWebServicesView extends StatefulWidget {
  const AmosAimWebServicesView({super.key});

  @override
  State<AmosAimWebServicesView> createState() => _AmosAimWebServicesViewState();
}

final class _AmosAimWebServicesViewState extends State<AmosAimWebServicesView> {
  @override
  void initState() {
    super.initState();
    context.read<GetAuth>().reset();
  }

  @override
  Widget build(context) {
    return Padding(
      padding: const EdgeInsets.all(16),
      child: BlocBuilder<GetAuth, GetAuthState>(
        builder: (context, state) {
          if (state is GetAuthLoading) {
            return const Center(child: CircularProgressIndicator());
          }

          if (state is GetAuthSuccess) {
            return _Body(token: state.token);
          }

          if (state is GetAuthFailure) {
            return Center(child: Text(state.message));
          }

          return _SignIn();
        },
      ),
    );
  }
}
