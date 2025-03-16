part of "view.dart";

final class _ImportCurrency extends StatefulWidget {
  const _ImportCurrency({required this.token});

  final String token;

  @override
  State<_ImportCurrency> createState() => _ImportCurrencyState();
}

final class _ImportCurrencyState extends State<_ImportCurrency> {
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
          setState(() {
            isLoading = false;
          });
        }
      },
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Form(
          key: formKey,
          autovalidateMode: AutovalidateMode.always,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              common_ui.DateInput(
                width: 512,
                controller: dateController,
                enabled: !isLoading,
                isLimitedUpToNow: true,
              ),
              SizedBox(height: 16),
              common_ui.ActionButton(
                width: 512,
                onPressed: () {
                  if (!formKey.currentState!.validate()) {
                    common_ui.showError(context, "Invalid Date Format");
                    return;
                  }

                  context.read<ImportCurrency>().execute(
                    token: widget.token,
                    date: dateController.text,
                  );
                },
                isLoading: isLoading,
                title: "IMPORT",
              ),
            ],
          ),
        ),
      ),
    );
  }
}
