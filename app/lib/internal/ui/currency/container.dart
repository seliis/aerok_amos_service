part of "screen.dart";

final class _Container extends StatelessWidget {
  const _Container();

  @override
  Widget build(context) {
    return Column(
      children: [
        const _DropdownMenu(),
        const SizedBox(
          height: 16,
        ),
        common_ui.DateField(
          controller: presenters.CurrencyControllers.date,
        ),
        const SizedBox(
          height: 16,
        ),
        const SizedBox(
          width: double.infinity,
          height: 48,
          child: _Button(),
        ),
      ],
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

final class _Button extends StatelessWidget {
  const _Button();

  @override
  Widget build(context) {
    final bool isLoading = context.watch<usecases.GetExchangeRate>().state is usecases.GetExchangeRateStateLoading;

    return FilledButton(
      onPressed: isLoading
          ? null
          : () {
              context.read<usecases.GetExchangeRate>().execute(
                    currencyCode: (context.read<usecases.GetCurrencies>().state as usecases.GetCurrenciesStateSuccess).currentCurrency.code,
                    date: presenters.CurrencyControllers.date.text,
                  );
            },
      style: FilledButton.styleFrom(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8),
        ),
      ),
      child: isLoading
          ? const common_ui.ProgressIndicator(
              scale: 0.5,
            )
          : const Text("Query"),
    );
  }
}
