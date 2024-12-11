import "package:flutter/material.dart";

final class ErrorDialog extends StatelessWidget {
  const ErrorDialog({
    super.key,
    this.width = 512,
    this.height = 256,
    required this.message,
    required this.stackTrace,
  });

  final String message;
  final String stackTrace;
  final double width;
  final double height;

  @override
  Widget build(context) {
    return AlertDialog(
      title: const Text(
        "ERROR",
        style: TextStyle(
          fontWeight: FontWeight.w700,
        ),
      ),
      content: SizedBox(
        width: width,
        height: height,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(message),
            const SizedBox(
              height: 16,
            ),
            Expanded(
              child: SingleChildScrollView(
                child: Text(stackTrace),
              ),
            ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: const Text("Close"),
        ),
      ],
    );
  }
}
