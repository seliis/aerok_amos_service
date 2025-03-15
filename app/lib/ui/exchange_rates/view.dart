import "package:flutter/material.dart";

final class ExchangeRatesView extends StatefulWidget {
  const ExchangeRatesView({super.key});

  @override
  State<ExchangeRatesView> createState() => _ExchangeRatesViewState();
}

final class _ExchangeRatesViewState extends State<ExchangeRatesView> {
  @override
  Widget build(BuildContext context) {
    return Text("Exchange Rates");
  }
}
