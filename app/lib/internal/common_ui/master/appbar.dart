import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;

final class MasterAppBar extends StatelessWidget implements PreferredSizeWidget {
  const MasterAppBar({
    super.key,
    required this.popupMenuEntries,
  });

  final List<PopupMenuEntry<dynamic>> popupMenuEntries;

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  Widget build(context) {
    return BlocBuilder(
      bloc: BlocProvider.of<usecases.RequestWebServiceKey>(context),
      builder: (context, state) {
        return AppBar(
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

    return isLoading
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
