import "package:flutter/material.dart";

final class ProgressIndicator extends StatelessWidget {
  const ProgressIndicator({
    super.key,
    this.scale = 1.0,
  });

  final double scale;

  @override
  Widget build(context) {
    return Transform.scale(
      scale: scale,
      child: const CircularProgressIndicator(),
    );
  }
}
