import "package:flutter/material.dart";

void showSuccess(BuildContext context, String message) {
  final colorScheme = Theme.of(context).colorScheme;
  final messanger = ScaffoldMessenger.of(context);
  messanger.clearSnackBars();

  messanger.showSnackBar(
    SnackBar(
      dismissDirection: DismissDirection.horizontal,
      backgroundColor: colorScheme.secondary,
      content: Text(
        message,
        overflow: TextOverflow.ellipsis,
        style: TextStyle(
          color: colorScheme.onSecondary,
          fontWeight: FontWeight.w400,
          fontSize: 16,
        ),
      ),
    ),
  );
}

void showError(BuildContext context, String message) {
  final themeData = Theme.of(context);

  showDialog<void>(
    context: context,
    useRootNavigator: true,
    builder: (context) {
      return AlertDialog(
        title: Row(
          children: [
            Icon(
              Icons.warning_amber_outlined,
              color: themeData.colorScheme.error,
            ),
            SizedBox(width: 8),
            Text("ERROR", style: themeData.textTheme.titleLarge),
          ],
        ),
        contentPadding: EdgeInsets.all(16),
        content: SizedBox(
          width: 512,
          height: 128,
          child: SingleChildScrollView(
            child: Text(
              message,
              style: themeData.textTheme.bodySmall?.copyWith(
                fontFamily: "CascadiaCode",
              ),
            ),
          ),
        ),
        actions: [
          TextButton(
            onPressed: () {
              Navigator.of(context).pop();
            },
            child: Text("OK"),
          ),
        ],
      );
    },
  );
}
