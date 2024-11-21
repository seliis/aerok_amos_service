import "package:flutter/material.dart";

final class CurrencyControllers {
  const CurrencyControllers();

  static final date = TextEditingController(
    text: DateTime.now().toString().substring(0, 10),
  );

  static final amosId = TextEditingController(
    text: "admin",
  );

  static final amosPassword = TextEditingController();
}
