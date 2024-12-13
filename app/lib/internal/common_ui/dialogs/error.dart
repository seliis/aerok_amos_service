import "package:flutter/material.dart";

final class ErrorDialog extends StatelessWidget {
  const ErrorDialog({
    super.key,
    this.width = 512,
    this.height = 384,
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
        "Error",
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
            Text(
              message,
              textAlign: TextAlign.start,
            ),
            const SizedBox(
              height: 16,
            ),
            const Divider(),
            const Text(
              "Stack Trace",
              style: TextStyle(
                fontWeight: FontWeight.w700,
              ),
            ),
            const SizedBox(
              height: 8,
            ),
            Expanded(
              child: SingleChildScrollView(
                child: Text(
                  stackTrace,
                  style: const TextStyle(
                    fontFamily: "CascadiaCode",
                  ),
                ),
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
