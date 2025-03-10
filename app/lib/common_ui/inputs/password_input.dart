import "package:flutter/material.dart";

final class PasswordInput extends StatelessWidget {
  const PasswordInput({
    super.key,
    required this.controller,
    this.isLoading = false,
  });

  final TextEditingController controller;
  final bool isLoading;

  @override
  Widget build(context) {
    return TextFormField(
      controller: controller,
      decoration: InputDecoration(
        border: OutlineInputBorder(),
        labelText: "Password",
        floatingLabelBehavior: FloatingLabelBehavior.always,
        errorStyle: TextStyle(
          fontFamily: "CascadiaCode",
          fontWeight: FontWeight.w200,
        ),
      ),
      obscureText: true,
      enabled: !isLoading,
      validator: (value) {
        if (value == null || value.isEmpty) {
          return "Please Enter a Password";
        }

        return null;
      },
    );
  }
}
