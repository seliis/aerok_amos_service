import "package:flutter/material.dart";

final class ActionButton extends StatelessWidget {
  const ActionButton({
    super.key,
    required this.onPressed,
    this.width = 128,
    this.height = 48,
    this.title = "Execute",
    this.isLoading = false,
  });

  final void Function()? onPressed;
  final double width;
  final double height;
  final String title;
  final bool isLoading;

  @override
  Widget build(context) {
    return SizedBox(
      width: width,
      height: height,
      child: FilledButton(
        style: FilledButton.styleFrom(
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(4)),
        ),
        onPressed: isLoading ? null : onPressed,
        child:
            isLoading
                ? Transform.scale(
                  scale: 0.50,
                  child: CircularProgressIndicator(),
                )
                : Text(
                  title,
                  style: TextStyle(
                    fontFamily: "CascadiaCode",
                    fontWeight: FontWeight.w700,
                  ),
                ),
      ),
    );
  }
}
