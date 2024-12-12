import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;

final class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(context) {
    return const Scaffold(
      appBar: common_ui.MasterAppBar(
        popupMenuEntries: [],
      ),
      drawer: common_ui.MasterDrawer(),
      body: Center(
        child: Text("Home"),
      ),
    );
  }
}
