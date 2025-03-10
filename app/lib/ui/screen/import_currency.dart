part of "screen.dart";

final class _ImportCurrencyDialog extends StatefulWidget {
  const _ImportCurrencyDialog();

  @override
  State<_ImportCurrencyDialog> createState() => _ImportCurrencyDialogState();
}

final class _ImportCurrencyDialogState extends State<_ImportCurrencyDialog> {
  final passwordController = TextEditingController();
  final dateController = TextEditingController(
    text: DateFormat("yyyy-MM-dd").format(DateTime.now()),
  );
  final formKey = GlobalKey<FormState>();
  bool isLoading = false;

  @override
  Widget build(context) {
    return BlocListener<ImportCurrency, ImportCurrencyState>(
      listener: (context, state) {
        if (state is ImportCurrencyStateLoading) {
          setState(() {
            isLoading = true;
          });
        }

        if (state is ImportCurrencyStateFailure) {
          common_ui.showError(context, state.message);

          setState(() {
            isLoading = false;
          });
        }

        if (state is ImportCurrencyStateSuccess) {
          common_ui.showSuccess(context, "Imported");

          context.pop();

          setState(() {
            isLoading = false;
          });
        }
      },
      child: _Dialog(
        title: "Import Currency",
        actions: [
          TextButton(
            onPressed:
                isLoading
                    ? null
                    : () async {
                      if (!formKey.currentState!.validate()) {
                        common_ui.showError(context, "Invalid");

                        return;
                      }

                      setState(() {
                        isLoading = true;
                      });

                      await context.read<ImportCurrency>().execute(
                        passwordController.text,
                        dateController.text,
                      );

                      setState(() {
                        isLoading = false;
                      });
                    },
            child:
                isLoading
                    ? Transform.scale(
                      scale: 0.5,
                      child: CircularProgressIndicator(),
                    )
                    : Text("Execute"),
          ),
        ],
        child: Form(
          key: formKey,
          autovalidateMode: AutovalidateMode.always,
          child: Column(
            children: [
              common_ui.DateInput(
                width: 512,
                controller: dateController,
                enabled: !isLoading,
                isLimitedToNow: true,
              ),
              SizedBox(height: 16),
              common_ui.PasswordInput(
                controller: passwordController,
                isLoading: isLoading,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
