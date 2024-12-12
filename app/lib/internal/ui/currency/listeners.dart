part of "screen.dart";

final class _GetCurrenciesListener extends BlocListener<usecases.GetCurrencies, usecases.GetCurrenciesState> {
  _GetCurrenciesListener()
      : super(
          listener: (context, state) {
            if (state is usecases.GetCurrenciesStateFailure) {
              showDialog<void>(
                context: context,
                builder: (context) {
                  return common_ui.ErrorDialog(
                    message: state.messsage,
                    stackTrace: state.stackTrace,
                  );
                },
              );
            }
          },
        );
}

final class _GetExchangeRateListener extends BlocListener<usecases.GetExchangeRate, usecases.GetExchangeRateState> {
  _GetExchangeRateListener()
      : super(
          listener: (context, state) {
            if (state is usecases.GetExchangeRateStateFailure) {
              showDialog<void>(
                context: context,
                builder: (context) {
                  return common_ui.ErrorDialog(
                    message: state.message,
                    stackTrace: state.stackTrace,
                  );
                },
              );
            }

            if (state is usecases.GetExchangeRateStateSuccess) {
              showDialog<void>(
                context: context,
                builder: (context) {
                  return _ExchangeRateDialog(
                    exchangeRate: state.exchangeRate,
                    width: 512,
                  );
                },
              );
            }
          },
        );
}
