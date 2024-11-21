import "package:flutter/material.dart";

final class TextField extends StatelessWidget {
  const TextField({
    super.key,
    this.obscureText = false,
    required this.labelText,
    required this.controller,
  });

  final bool obscureText;
  final String labelText;
  final TextEditingController controller;

  @override
  Widget build(context) {
    return TextFormField(
      controller: controller,
      obscureText: obscureText,
      decoration: InputDecoration(
        border: const OutlineInputBorder(),
        labelText: labelText,
      ),
    );
  }
}
