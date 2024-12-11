part of "screen.dart";

final class _DateForm extends StatelessWidget {
  const _DateForm();

  @override
  Widget build(context) {
    return common_ui.DateField(
      controller: presenters.CurrencyControllers.date,
    );
  }
}

final class _QueryButton extends StatelessWidget {
  const _QueryButton();

  @override
  Widget build(context) {
    return common_ui.Button(
      size: const Size(
        double.infinity,
        48,
      ),
      isLoading: context.watch<usecases.GetExchangeRate>().state is usecases.GetExchangeRateStateLoading,
      onPressed: () {
        context.read<usecases.GetExchangeRate>().execute(
              currencyCode: (context.read<usecases.GetCurrencies>().state as usecases.GetCurrenciesStateSuccess).currentCurrency.code,
              date: presenters.CurrencyControllers.date.text,
            );
      },
      child: const Text("Query"),
    );
  }
}

final class _DropdownMenu extends StatelessWidget {
  const _DropdownMenu();

  @override
  Widget build(context) {
    final state = context.watch<usecases.GetCurrencies>().state as usecases.GetCurrenciesStateSuccess;

    return DropdownMenu(
      label: const Text("Currency Code"),
      initialSelection: state.currentCurrency,
      expandedInsets: EdgeInsets.zero,
      dropdownMenuEntries: [
        for (final currency in state.currencies)
          DropdownMenuEntry<entities.Currency>(
            value: currency,
            label: "${currency.code} (${currency.name})",
          ),
      ],
      onSelected: (currency) {
        if (currency != null) {
          context.read<usecases.GetCurrencies>().changeCurrency(currency);
        }
      },
    );
  }
}
