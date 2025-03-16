import "package:flutter/material.dart";
import "package:flutter/services.dart";

final class DateInput extends StatelessWidget {
  const DateInput({
    super.key,
    required this.controller,
    this.width = 256,
    this.height = 48,
    this.enabled = true,
    this.isLimitedUpToNow = false,
  });

  final TextEditingController controller;
  final bool isLimitedUpToNow;
  final double width;
  final double height;
  final bool enabled;

  @override
  Widget build(context) {
    return SizedBox(
      width: width,
      child: TextFormField(
        enabled: enabled,
        controller: controller,
        decoration: InputDecoration(
          floatingLabelBehavior: FloatingLabelBehavior.always,
          border: OutlineInputBorder(),
          hintText: "YYYYMMDD",
          labelText: "Date",
          errorStyle: TextStyle(
            fontFamily: "CascadiaCode",
            fontWeight: FontWeight.w200,
          ),
        ),
        inputFormatters: [
          FilteringTextInputFormatter.allow(RegExp(r"[\d-]")),
          _DateInputFormatter(),
        ],
        validator: (value) {
          if (value == null ||
              !RegExp(r"^\d{4}-\d{2}-\d{2}$").hasMatch(value)) {
            return "Insert Date in Format YYYYMMDD";
          }
          try {
            final parts = value.split("-");
            final year = int.parse(parts[0]);
            final month = int.parse(parts[1]);
            final day = int.parse(parts[2]);
            final date = DateTime(year, month, day);

            if (date.year != year || date.month != month || date.day != day) {
              return "Invalid Year, Month, or Day";
            }

            if (isLimitedUpToNow && date.isAfter(DateTime.now())) {
              return "Date Cannot Be in the Future";
            }
          } catch (e) {
            return "Invalid Date Format";
          }
          return null;
        },
      ),
    );
  }
}

final class _DateInputFormatter extends TextInputFormatter {
  @override
  TextEditingValue formatEditUpdate(
    TextEditingValue oldValue,
    TextEditingValue newValue,
  ) {
    final digitsOnly = newValue.text.replaceAll(RegExp(r"\D"), "");

    final limitedDigits =
        digitsOnly.length > 8 ? digitsOnly.substring(0, 8) : digitsOnly;

    String newText = limitedDigits;
    if (limitedDigits.length == 8) {
      newText =
          "${limitedDigits.substring(0, 4)}-${limitedDigits.substring(4, 6)}-${limitedDigits.substring(6, 8)}";
    }

    return TextEditingValue(
      text: newText,
      selection: TextSelection.collapsed(offset: newText.length),
    );
  }
}
