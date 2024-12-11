import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;

final class MasterAppBar extends StatelessWidget implements PreferredSizeWidget {
  const MasterAppBar({super.key});

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  Widget build(context) {
    return BlocBuilder(
      bloc: BlocProvider.of<usecases.RequestWebServiceKey>(context),
      builder: (context, state) {
        return AppBar(
          actions: const [
            _RequestWebServiceKeyButton(),
            SizedBox(
              width: 16,
            ),
          ],
        );
      },
    );
  }
}

final class _RequestWebServiceKeyButton extends StatelessWidget {
  const _RequestWebServiceKeyButton();

  @override
  Widget build(context) {
    final bool isSuccess = context.watch<usecases.RequestWebServiceKey>().state is usecases.RequestWebServiceKeyStateSuccess;

    return IconButton(
      icon: Icon(isSuccess ? Icons.vpn_key_outlined : Icons.vpn_key_off_outlined),
      onPressed: isSuccess
          ? null
          : () {
              showDialog<void>(
                context: context,
                builder: (context) {
                  return AlertDialog(
                    title: const Text("Request AMOS WebService Key"),
                    content: SizedBox(
                      width: 512,
                      child: common_ui.TextField(
                        enabled: false,
                        autofocus: true,
                        obscureText: true,
                        labelText: "Password",
                        controller: TextEditingController(),
                      ),
                    ),
                    actions: [
                      TextButton(
                        onPressed: () {
                          context.read<usecases.RequestWebServiceKey>().execute("8Z7Plc3p!!");
                          Navigator.of(context).pop();
                        },
                        child: const Text("Request"),
                      ),
                    ],
                  );
                },
              );
            },
    );
  }
}
