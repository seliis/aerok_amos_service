import "package:flutter/material.dart";

final class Button extends StatelessWidget {
  const Button({
    super.key,
    this.size = const Size(128, 48),
    required this.onPressed,
    required this.child,
  });

  final void Function() onPressed;
  final Widget child;
  final Size size;

  @override
  Widget build(context) {
    return SizedBox(
      width: size.width,
      height: size.height,
      child: FilledButton(
        style: ButtonStyle(
          shape: WidgetStatePropertyAll(
            RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(8),
            ),
          ),
        ),
        onPressed: onPressed,
        child: child,
      ),
    );
  }
}
