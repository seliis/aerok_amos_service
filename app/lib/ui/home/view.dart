import "package:aerok_amos_service/common_ui/index.dart" as common_ui;
import "package:aerok_amos_service/usecases/index.dart";
import "package:aerok_amos_service/entities/index.dart";
import "package:flutter_bloc/flutter_bloc.dart";
import "package:go_router/go_router.dart";
import "package:flutter/material.dart";
import "package:flutter/services.dart";
import "package:intl/intl.dart";

final class HomeView extends StatefulWidget {
  const HomeView({super.key});

  @override
  State<HomeView> createState() => _HomeViewState();
}

final class _HomeViewState extends State<HomeView> {
  @override
  void initState() {
    super.initState();

    context.read<GetCurrencies>().execute();
  }

  @override
  Widget build(context) {
    return Center(
      child: BlocBuilder<GetCurrencies, GetCurrenciesState>(
        builder: (context, state) {
          if (state is GetCurrenciesStateLoading) {
            return CircularProgressIndicator();
          }

          if (state is GetCurrenciesStateSuccess) {
            return _Body(currencies: state.currencies);
          }

          if (state is GetCurrenciesStateFailure) {
            return Text(
              state.message,
              style: Theme.of(context).textTheme.titleLarge,
            );
          }

          return const SizedBox.shrink();
        },
      ),
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
  final dateInputController = TextEditingController(
    text: DateFormat("yyyy-MM-dd").format(DateTime.now()),
  );

  final formKey = GlobalKey<FormState>();
  late Currency selectedCurrency;
  bool isLoading = false;
  bool isCopied = false;

  @override
  void initState() {
    super.initState();
    selectedCurrency =
        widget.currencies.where((currency) => currency.code == "USD").first;
  }

  @override
  Widget build(context) {
    final theme = Theme.of(context);

    return BlocListener<GetExchangeRate, GetExchangeRateState>(
      listener: (context, state) {
        if (state is GetExchangeRateStateLoading) {
          setState(() {
            isLoading = true;
          });
        }

        if (state is GetExchangeRateStateSuccess) {
          showDialog<void>(
            context: context,
            builder: (context) {
              return _Dialog(exchangeRate: state.exchangeRate);
            },
          );

          setState(() {
            isLoading = false;
          });
        }

        if (state is GetExchangeRateStateFailure) {
          common_ui.showError(context, state.message);

          setState(() {
            isLoading = false;
          });
        }
      },
      child: Container(
        width: 384,
        padding: const EdgeInsets.all(16),
        child: Form(
          key: formKey,
          autovalidateMode: AutovalidateMode.always,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              DropdownMenu<String>(
                enabled: !isLoading,
                label: Text("Currency"),
                menuHeight: 256,
                initialSelection: selectedCurrency.code,
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
              const SizedBox(height: 16),
              common_ui.DateInput(
                controller: dateInputController,
                width: double.infinity,
                isLimitedUpToNow: true,
                enabled: !isLoading,
              ),
              const SizedBox(height: 16),
              common_ui.ActionButton(
                title: "FETCH",
                width: double.infinity,
                onPressed: () {
                  if (!formKey.currentState!.validate()) {
                    common_ui.showError(context, "Invalid Date Format");
                    return;
                  }

                  context.read<GetExchangeRate>().execute(
                    code: selectedCurrency.code,
                    date: dateInputController.text,
                  );
                },
                isLoading: isLoading,
              ),
            ],
          ),
        ),
      ),
    );
  }
}

final class _Dialog extends StatefulWidget {
  const _Dialog({required this.exchangeRate});

  final ExchangeRateWithCurrency exchangeRate;

  @override
  State<_Dialog> createState() => _DialogState();
}

final class _DialogState extends State<_Dialog> {
  bool isCopied = false;

  @override
  Widget build(context) {
    final rate = widget.exchangeRate.rate.toStringAsFixed(2);
    final themeData = Theme.of(context);

    return AlertDialog(
      title: Text(
        "${widget.exchangeRate.code} (${widget.exchangeRate.name})",
        style: themeData.textTheme.titleLarge,
      ),
      content: SizedBox(
        width: 512,
        height: 128,
        child: ListTile(
          contentPadding: EdgeInsets.zero,
          title: Text(rate, style: themeData.textTheme.titleMedium),
          subtitle: Padding(
            padding: EdgeInsets.only(top: 8),
            child: Text(
              widget.exchangeRate.date,
              style: themeData.textTheme.bodySmall?.copyWith(
                fontWeight: FontWeight.w200,
              ),
            ),
          ),
          trailing: TextButton(
            onPressed: () {
              Clipboard.setData(ClipboardData(text: rate));
              setState(() {
                isCopied = true;
              });
            },
            child: isCopied ? Text("Copied") : Icon(Icons.copy),
          ),
        ),
      ),
      actions: [
        TextButton(
          onPressed: () {
            context.pop();
          },
          child: Text("OK"),
        ),
      ],
    );
  }
}
