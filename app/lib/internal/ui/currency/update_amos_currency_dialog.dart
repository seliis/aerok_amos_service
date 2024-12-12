part of "screen.dart";

final class _UpdateAmosCurrencyDialog extends StatelessWidget {
  const _UpdateAmosCurrencyDialog({
    required this.dateController,
  });

  final TextEditingController dateController;

  @override
  Widget build(context) {
    return BlocBuilder<usecases.UpdateAmosCurrency, usecases.UpdateAmosCurrencyState>(
      bloc: BlocProvider.of<usecases.UpdateAmosCurrency>(context),
      builder: (context, state) {
        final keyState = context.read<usecases.RequestWebServiceKey>().state;
        final isSuccessKeyState = keyState is usecases.RequestWebServiceKeyStateSuccess;

        return AlertDialog(
          title: const Text("Update AMOS Currency"),
          content: SizedBox(
            width: 512,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                common_ui.DateField(
                  controller: dateController,
                ),
                const SizedBox(
                  height: 16,
                ),
                common_ui.Button(
                  onPressed: isSuccessKeyState
                      ? () {
                          BlocProvider.of<usecases.UpdateAmosCurrency>(context).execute(
                            authKey: keyState.key,
                            date: dateController.text,
                          );
                        }
                      : null,
                  isFullWidth: true,
                  isLoading: state is usecases.UpdateAmosCurrencyStateLoading,
                  child: Text(isSuccessKeyState ? "Update" : "Not Granted"),
                ),
              ],
            ),
          ),
        );
      },
    );
  }
}
