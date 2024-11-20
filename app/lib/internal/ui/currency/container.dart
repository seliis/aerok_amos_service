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
          controller: presenters.Currency.controller,
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

    void Function()? onPressed() {
      if (isLoading) {
        return null;
      }

      return () {
        context.read<usecases.GetExchangeRate>().execute(
              currencyCode: (context.read<usecases.GetCurrencies>().state as usecases.GetCurrenciesStateSuccess).currentCurrency.code,
              date: presenters.Currency.controller.text,
            );
      };
    }

    return FilledButton(
      onPressed: onPressed(),
      style: FilledButton.styleFrom(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8),
        ),
      ),
      child: isLoading
          ? Transform.scale(
              scale: 0.50,
              child: const CircularProgressIndicator(),
            )
          : const Text("Query"),
    );
  }
}
