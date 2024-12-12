import "package:flutter/material.dart";

final class CurrencyControllers {
  const CurrencyControllers();

  static final date = TextEditingController(
    text: DateTime.now().toString().substring(0, 10),
  );
}
