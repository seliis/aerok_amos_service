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
          fontWeight: FontWeight.w700,
          fontFamily: "CascadiaCode",
          fontSize: 16,
        ),
      ),
    ),
  );
}

void showError(BuildContext context, String message) {
  message = message.replaceAll("Exception:", "").toUpperCase();
  final colorScheme = Theme.of(context).colorScheme;
  final messanger = ScaffoldMessenger.of(context);
  messanger.clearSnackBars();

  messanger.showSnackBar(
    SnackBar(
      dismissDirection: DismissDirection.horizontal,
      backgroundColor: colorScheme.error,
      content: Text(
        message,
        overflow: TextOverflow.ellipsis,
        style: TextStyle(
          color: colorScheme.onError,
          fontWeight: FontWeight.w700,
          fontFamily: "CascadiaCode",
          fontSize: 16,
        ),
      ),
      duration: Duration(seconds: 60),
      action: SnackBarAction(
        label: "SHOW",
        textColor: colorScheme.onError,
        onPressed: () {
          showDialog<void>(
            context: context,
            builder: (context) {
              return AlertDialog(
                title: Text(
                  "ERROR",
                  style: TextStyle(fontWeight: FontWeight.w700),
                ),
                contentPadding: EdgeInsets.all(16),
                content: SizedBox(
                  width: 512,
                  height: 256,
                  child: Text(
                    message,
                    overflow: TextOverflow.ellipsis,
                    style: TextStyle(fontFamily: "CascadiaCode", fontSize: 16),
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
        },
      ),
    ),
  );
}
