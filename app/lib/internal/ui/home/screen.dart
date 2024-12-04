import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;

final class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(context) {
    return Scaffold(
      appBar: AppBar(
        actions: [
          IconButton(
            icon: const Icon(Icons.key),
            onPressed: () {
              showDialog<void>(
                context: context,
                builder: (context) {
                  return AlertDialog(
                    title: Text("AMOS Web Service Authentication"),
                    content: SizedBox(
                      width: 512,
                      child: common_ui.TextField(
                        autofocus: true,
                        labelText: "Password",
                        controller: TextEditingController(),
                      ),
                    ),
                    actions: [
                      TextButton(
                        onPressed: () {},
                        child: Text("Submit"),
                      ),
                    ],
                  );
                },
              );
            },
          ),
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
