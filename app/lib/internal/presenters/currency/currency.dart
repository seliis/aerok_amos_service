import "package:flutter/material.dart";

final class Currency {
  const Currency();

  static final controller = TextEditingController(
    text: DateTime.now().toString().substring(0, 10),
  );
}
