import "package:flutter/material.dart";

import "package:app/internal/common_ui/__index.dart" as common_ui;

final class Button extends StatelessWidget {
  const Button({
    super.key,
    this.size = const Size(128, 48),
    this.isFullWidth = false,
    this.isLoading = false,
    required this.onPressed,
    required this.child,
  });

  final void Function()? onPressed;
  final bool isFullWidth;
  final bool isLoading;
  final Widget child;
  final Size size;

  @override
  Widget build(context) {
    return SizedBox(
      width: isFullWidth ? double.infinity : size.width,
      height: size.height,
      child: FilledButton(
        style: ButtonStyle(
          shape: WidgetStatePropertyAll(
            RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(8),
            ),
          ),
        ),
        onPressed: isLoading ? null : onPressed,
        child: isLoading
            ? const common_ui.ProgressIndicator(
                scale: 0.5,
              )
            : child,
      ),
    );
  }
}
