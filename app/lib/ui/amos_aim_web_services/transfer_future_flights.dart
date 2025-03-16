part of "view.dart";

final class _TransferFutureFlights extends StatefulWidget {
  const _TransferFutureFlights({required this.token});

  final String token;

  @override
  State<_TransferFutureFlights> createState() => _TransferFutureFlightsState();
}

final class _TransferFutureFlightsState extends State<_TransferFutureFlights> {
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
                              allowedExtensions: ["xlsx"],
                              type: FileType.custom,
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
              common_ui.ActionButton(
                width: 512,
                enabled: pickedFile != null,
                onPressed: () {
                  if (pickedFile == null) {
                    common_ui.showError(context, "Please select a file");
                    return;
                  }

                  context.read<TransferFutureFlights>().execute(
                    token: widget.token,
                    file: pickedFile!,
                  );
                },
                title: "TRANSFER",
                isLoading: isLoading,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
