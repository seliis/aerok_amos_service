part of "screen.dart";

final class _UpdateAmosCurrencyDialog extends StatelessWidget {
  const _UpdateAmosCurrencyDialog({
    required this.dateController,
  });

  final TextEditingController dateController;

  @override
  Widget build(context) {
    return BlocListener<usecases.UpdateAmosCurrency, usecases.UpdateAmosCurrencyState>(
      bloc: BlocProvider.of<usecases.UpdateAmosCurrency>(context),
      listener: (context, state) {
        if (state is usecases.UpdateAmosCurrencyStateFailure) {
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

        if (state is usecases.UpdateAmosCurrencyStateSuccess) {
          Navigator.pop(context);
          ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
            content: Text("Updated"),
            backgroundColor: Colors.teal,
          ));
        }
      },
      child: BlocBuilder<usecases.UpdateAmosCurrency, usecases.UpdateAmosCurrencyState>(
        bloc: BlocProvider.of<usecases.UpdateAmosCurrency>(context),
        builder: (context, state) {
          return AlertDialog(
            title: const Text("Update AMOS Currency"),
            content: SizedBox(
              width: 512,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  common_ui.DateField(
                    enabled: state is! usecases.UpdateAmosCurrencyStateLoading,
                    controller: dateController,
                  ),
                  const SizedBox(
                    height: 16,
                  ),
                  common_ui.Button(
                    onPressed: () {
                      final authorization = BlocProvider.of<usecases.RequestWebServiceKey>(context).state as usecases.RequestWebServiceKeyStateSuccess;

                      BlocProvider.of<usecases.UpdateAmosCurrency>(context).execute(
                        authorization: authorization.key,
                        date: dateController.text,
                      );
                    },
                    isFullWidth: true,
                    isLoading: state is usecases.UpdateAmosCurrencyStateLoading,
                    child: const Text("Update"),
                  ),
                ],
              ),
            ),
          );
        },
      ),
    );
  }
}
