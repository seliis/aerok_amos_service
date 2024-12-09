import "package:flutter/material.dart";

final class TextField extends StatelessWidget {
  const TextField({
    super.key,
    this.enabled = true,
    this.autofocus = false,
    this.obscureText = false,
    required this.labelText,
    required this.controller,
  });

  final bool enabled;
  final bool autofocus;
  final bool obscureText;
  final String labelText;
  final TextEditingController controller;

  @override
  Widget build(context) {
    return TextFormField(
      enabled: enabled,
      autofocus: autofocus,
      controller: controller,
      obscureText: obscureText,
      decoration: InputDecoration(
        border: const OutlineInputBorder(),
        labelText: labelText,
      ),
    );
  }
}
