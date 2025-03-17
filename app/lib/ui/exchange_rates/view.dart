import "package:aerok_amos_service/common_ui/index.dart" as common_ui;
import "package:aerok_amos_service/entities/index.dart";
import "package:aerok_amos_service/usecases/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:flutter/services.dart";
import "package:flutter/material.dart";

final class ExchangeRatesView extends StatefulWidget {
  const ExchangeRatesView({super.key});

  @override
  State<ExchangeRatesView> createState() => _ExchangeRatesViewState();
}

final class _ExchangeRatesViewState extends State<ExchangeRatesView> {
  @override
  void initState() {
    super.initState();

    context.read<GetCurrencies>().execute();
  }

  @override
  Widget build(context) {
    return BlocBuilder<GetCurrencies, GetCurrenciesState>(
      builder: (context, state) {
        if (state is GetCurrenciesStateLoading) {
          return Center(child: CircularProgressIndicator());
        }

        if (state is GetCurrenciesStateSuccess) {
          return Padding(
            padding: const EdgeInsets.all(16),
            child: _Body(currencies: state.currencies),
          );
        }

        if (state is GetCurrenciesStateFailure) {
          return Center(child: Text(state.message));
        }

        return const SizedBox.shrink();
      },
    );
  }
}

final class _Body extends StatefulWidget {
  const _Body({required this.currencies});

  final List<Currency> currencies;

  @override
  State<_Body> createState() => _BodyState();
}

final class _BodyState extends State<_Body> {
  final yearController = TextEditingController(
    text: DateTime.now().year.toString(),
  );
  final formKey = GlobalKey<FormState>();
  late Currency selectedCurrency;
  bool isLoading = false;

  @override
  void initState() {
    super.initState();
    selectedCurrency =
        widget.currencies.where((currency) => currency.code == "USD").first;
  }

  @override
  Widget build(context) {
    final theme = Theme.of(context);

    return BlocListener<GetAnnualExchangeRates, GetAnnualExchangeRatesState>(
      listener: (context, state) {
        if (state is GetAnnualExchangeRatesStateLoading) {
          setState(() {
            isLoading = true;
          });
        }

        setState(() {
          isLoading = false;
        });
      },
      child: Form(
        key: formKey,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              width: 384,
              child: DropdownMenu<String>(
                initialSelection: selectedCurrency.code,
                label: Text("Currency"),
                enabled: !isLoading,
                menuHeight: 256,
                dropdownMenuEntries:
                    widget.currencies.map((currency) {
                      return DropdownMenuEntry<String>(
                        value: currency.code,
                        label: "${currency.code} (${currency.name})",
                        labelWidget: Text(
                          "${currency.code} (${currency.name})",
                          style: theme.textTheme.bodyMedium?.copyWith(
                            fontFamily: "CascadiaCode",
                          ),
                        ),
                      );
                    }).toList(),
                expandedInsets: EdgeInsets.zero,
                onSelected: (code) {
                  setState(() {
                    selectedCurrency =
                        widget.currencies
                            .where((currency) => currency.code == code)
                            .first;
                  });
                },
              ),
            ),
            SizedBox(height: 16),
            SizedBox(
              width: 384,
              child: TextFormField(
                controller: yearController,
                keyboardType: TextInputType.number,
                autovalidateMode: AutovalidateMode.onUserInteraction,
                inputFormatters: [
                  FilteringTextInputFormatter.digitsOnly,
                  LengthLimitingTextInputFormatter(4),
                ],
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: "Year",
                ),
                validator: (value) {
                  final currentYear = DateTime.now().year;

                  if (value == null || value.isEmpty) {
                    return "1991 ~ $currentYear";
                  }

                  final year = int.tryParse(value);
                  if (year == null) {
                    return "1991 ~ $currentYear";
                  }

                  if (year < 1991 || year > currentYear) {
                    return "1991 ~ $currentYear";
                  }

                  return null;
                },
              ),
            ),
            SizedBox(height: 16),
            common_ui.ActionButton(
              title: "DOWNLOAD",
              width: 384,
              onPressed: () {
                if (formKey.currentState!.validate()) {
                  context.read<GetAnnualExchangeRates>().execute(
                    code: selectedCurrency.code,
                    year: yearController.text,
                  );
                }
              },
              isLoading: isLoading,
            ),
          ],
        ),
      ),
    );
  }
}
