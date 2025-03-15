import "package:aerok_amos_service/common_ui/index.dart" as common_ui;
import "package:aerok_amos_service/usecases/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

final class AmosAimWebServicesView extends StatefulWidget {
  const AmosAimWebServicesView({super.key});

  @override
  State<AmosAimWebServicesView> createState() => _AmosAimWebServicesViewState();
}

final class _AmosAimWebServicesViewState extends State<AmosAimWebServicesView> {
  @override
  void initState() {
    super.initState();
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
            return const Center(child: Text("Authorized"));
          }

          if (state is GetAuthFailure) {
            return const Center(child: Text("Unauthorized"));
          }

          return _Auth();
        },
      ),
    );
  }
}

final class _Auth extends StatefulWidget {
  const _Auth();

  @override
  State<_Auth> createState() => _AuthState();
}

final class _AuthState extends State<_Auth> {
  final TextEditingController controller = TextEditingController();
  final formKey = GlobalKey<FormState>();

  @override
  Widget build(context) {
    return Form(
      key: formKey,
      autovalidateMode: AutovalidateMode.onUserInteraction,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          SizedBox(
            width: 512,
            child: common_ui.PasswordInput(controller: controller),
          ),
          const SizedBox(height: 8),
          common_ui.ActionButton(
            width: 512,
            title: "Authorize",
            onPressed: () {
              if (!formKey.currentState!.validate()) {
                common_ui.showError(context, "Invalid Password");
                return;
              }

              context.read<GetAuth>().execute();
            },
          ),
        ],
      ),
    );
  }
}
