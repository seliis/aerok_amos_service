import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/presenters/__index.dart" as presenters;
import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;

final class MasterAppBar extends StatelessWidget implements PreferredSizeWidget {
  const MasterAppBar({
    super.key,
    this.title,
    required this.popupMenuEntries,
  });

  final String? title;
  final List<PopupMenuEntry<dynamic>> popupMenuEntries;

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  Widget build(context) {
    return BlocBuilder(
      bloc: BlocProvider.of<usecases.RequestWebServiceKey>(context),
      builder: (context, state) {
        return AppBar(
          title: Text(title ?? ""),
          actions: [
            const _RequestWebServiceKeyButton(),
            _PopUpMenuButton(
              popupMenuEntries: popupMenuEntries,
            ),
            const SizedBox(
              width: 16,
            ),
          ],
        );
      },
    );
  }
}

final class _PopUpMenuButton extends StatelessWidget {
  const _PopUpMenuButton({
    required this.popupMenuEntries,
  });

  final List<PopupMenuEntry<dynamic>> popupMenuEntries;

  @override
  Widget build(context) {
    return PopupMenuButton(
      enabled: context.watch<usecases.RequestWebServiceKey>().state is usecases.RequestWebServiceKeyStateSuccess,
      itemBuilder: (context) {
        return popupMenuEntries;
      },
    );
  }
}

final class _RequestWebServiceKeyButton extends StatelessWidget {
  const _RequestWebServiceKeyButton();

  @override
  Widget build(context) {
    final bool isLoading = context.watch<usecases.RequestWebServiceKey>().state is usecases.RequestWebServiceKeyStateLoading;
    final bool isSuccess = context.watch<usecases.RequestWebServiceKey>().state is usecases.RequestWebServiceKeyStateSuccess;

    return BlocListener<usecases.RequestWebServiceKey, usecases.RequestWebServiceKeyState>(
      bloc: BlocProvider.of<usecases.RequestWebServiceKey>(context),
      listener: (context, state) {
        if (state is usecases.RequestWebServiceKeyStateSuccess) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("Granted"),
              backgroundColor: Colors.teal,
            ),
          );
        }

        if (state is usecases.RequestWebServiceKeyStateFailure) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("Unauthorized"),
              backgroundColor: Colors.pink,
            ),
          );
        }
      },
      child: isLoading
          ? const common_ui.ProgressIndicator(
              scale: 0.5,
            )
          : IconButton(
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
                                enabled: true,
                                autofocus: true,
                                obscureText: true,
                                labelText: "Password",
                                hintText: "Web Service Password",
                                helperText: "AMOS Configure Server (APN: 10)",
                                controller: presenters.AppBarControllers.webServicePassword,
                              ),
                            ),
                            actions: [
                              TextButton(
                                onPressed: () {
                                  context.read<usecases.RequestWebServiceKey>().execute(presenters.AppBarControllers.webServicePassword.text);
                                  Navigator.of(context).pop();
                                },
                                child: const Text("Request"),
                              ),
                            ],
                          );
                        },
                      );
                    },
            ),
    );
  }
}
