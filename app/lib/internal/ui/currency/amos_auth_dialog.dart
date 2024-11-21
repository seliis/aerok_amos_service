part of "screen.dart";

final class _AmosAuthDialog extends StatelessWidget {
  const _AmosAuthDialog({
    required this.width,
    required this.height,
  });

  final double width;
  final double height;

  @override
  Widget build(context) {
    final isLoading = context.watch<usecases.UpdateAmosCurrency>().state is usecases.UpdateAmosCurrencyStateLoading;

    return AlertDialog(
      title: const Text("AMOS Web Service Sign-In"),
      content: SizedBox(
        width: width,
        height: height,
        child: BlocListener<usecases.UpdateAmosCurrency, usecases.UpdateAmosCurrencyState>(
          bloc: BlocProvider.of<usecases.UpdateAmosCurrency>(context),
          listener: (context, state) {
            if (state is usecases.UpdateAmosCurrencyStateSuccess) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text("Updated"),
                ),
              );
              Navigator.of(context).pop();
            }

            if (state is usecases.UpdateAmosCurrencyStateFailure) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(state.message),
                ),
              );
            }
          },
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              common_ui.TextField(
                controller: presenters.CurrencyControllers.amosId,
                labelText: "ID",
              ),
              const SizedBox(
                height: 16,
              ),
              common_ui.TextField(
                controller: presenters.CurrencyControllers.amosPassword,
                labelText: "Password",
                obscureText: true,
              ),
              const SizedBox(
                height: 16,
              ),
              SizedBox(
                width: double.infinity,
                height: 48,
                child: FilledButton(
                  onPressed: isLoading
                      ? null
                      : () {
                          context.read<usecases.UpdateAmosCurrency>().execute(
                                id: presenters.CurrencyControllers.amosId.text,
                                password: presenters.CurrencyControllers.amosPassword.text,
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
                      : const Text("Update"),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
