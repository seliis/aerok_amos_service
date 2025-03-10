part of "screen.dart";

final class _TransferFutureFlightsDialog extends StatefulWidget {
  const _TransferFutureFlightsDialog();

  @override
  State<_TransferFutureFlightsDialog> createState() =>
      _TransferFutureFlightsDialogState();
}

final class _TransferFutureFlightsDialogState
    extends State<_TransferFutureFlightsDialog> {
  final passwordController = TextEditingController();
  final dateController = TextEditingController(
    text: DateFormat("yyyy-MM-dd").format(DateTime.now()),
  );
  final formKey = GlobalKey<FormState>();
  PlatformFile? pickedFile;
  bool isLoading = false;

  @override
  Widget build(context) {
    return BlocListener<TransferFutureFlights, TransferFutureFlightsState>(
      listener: (context, state) {
        if (state is TransferFutureFlightsLoading) {
          setState(() {
            isLoading = true;
          });
        }

        if (state is TransferFutureFlightsFailure) {
          common_ui.showError(context, state.message);

          setState(() {
            isLoading = false;
          });
        }

        if (state is TransferFutureFlightsSuccess) {
          common_ui.showSuccess(context, "Transferred");

          context.pop();

          setState(() {
            isLoading = false;
          });
        }
      },
      child: _Dialog(
        title: "Transfer Future Flights",
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

                      await context.read<TransferFutureFlights>().execute(
                        file: pickedFile!,
                        password: passwordController.text,
                        date: dateController.text,
                      );
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
              SizedBox(
                width: 512,
                height: 48,
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(4),
                    ),
                  ),
                  onPressed:
                      isLoading
                          ? null
                          : () async {
                            final result = await FilePicker.platform.pickFiles(
                              type: FileType.custom,
                              allowedExtensions: ["xlsx"],
                            );

                            if (result == null) {
                              return;
                            }

                            setState(() {
                              pickedFile = result.files.single;
                            });
                          },
                  child: Text(
                    pickedFile == null ? "Select a File" : pickedFile!.name,
                  ),
                ),
              ),
              SizedBox(height: 16),
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
