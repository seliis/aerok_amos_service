import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;
import "package:app/internal/usecases/__index.dart" as usecases;

final class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(context) {
    return BlocBuilder(
      bloc: context.read<usecases.RequestWebServiceKey>(),
      builder: (context, state) {
        return _Screen();
      },
    );
  }
}

final class _RequestWebServiceKeyButton extends StatelessWidget {
  const _RequestWebServiceKeyButton();

  @override
  Widget build(context) {
    return IconButton(
      icon: const Icon(Icons.key),
      onPressed: () {
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

final class _Screen extends StatelessWidget {
  const _Screen();

  @override
  Widget build(context) {
    return Scaffold(
      appBar: AppBar(
        actions: [
          if (context.read<usecases.RequestWebServiceKey>().state is! usecases.RequestWebServiceKeyStateSuccess) ...[
            const _RequestWebServiceKeyButton(),
          ],
          const SizedBox(
            width: 16,
          ),
        ],
      ),
      drawer: _Drawer(),
      body: SizedBox.shrink(),
    );
  }
}

final class _Drawer extends StatelessWidget {
  const _Drawer();

  @override
  Widget build(context) {
    return Drawer();
  }
}
