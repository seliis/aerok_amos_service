part of "screen.dart";

final class _UpdateAmosFutureFlightsDialog extends StatelessWidget {
  const _UpdateAmosFutureFlightsDialog();

  @override
  Widget build(context) {
    return BlocListener<usecases.UpdateAmosFutureFlights, usecases.UpdateAmosFutureFlightsState>(
      bloc: BlocProvider.of<usecases.UpdateAmosFutureFlights>(context),
      listener: (context, state) {
        if (state is usecases.UpdateAmosFutureFlightsStateFailure) {
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

        if (state is usecases.UpdateAmosFutureFlightsStateSuccess) {
          Navigator.pop(context);
          ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
            content: Text("Updated"),
            backgroundColor: Colors.teal,
          ));
        }
      },
      child: BlocBuilder<usecases.UpdateAmosFutureFlights, usecases.UpdateAmosFutureFlightsState>(
        bloc: BlocProvider.of<usecases.UpdateAmosFutureFlights>(context),
        builder: (context, state) {
          final bool isPrepared = context.watch<usecases.OpenFutureFlights>().state is usecases.OpenFutureFlightsStateSuccess;
          final bool isLoading = state is usecases.UpdateAmosFutureFlightsStateLoading;

          return AlertDialog(
            title: const Text("Update AMOS Future Flights"),
            content: SizedBox(
              width: 512,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  const _Picker(),
                  const SizedBox(height: 16),
                  common_ui.Button(
                    isFullWidth: true,
                    isLoading: isLoading,
                    onPressed: isPrepared
                        ? () {
                            final authorization = context.read<usecases.RequestWebServiceKey>().state as usecases.RequestWebServiceKeyStateSuccess;
                            final picked = context.read<usecases.OpenFutureFlights>().state as usecases.OpenFutureFlightsStateSuccess;

                            BlocProvider.of<usecases.UpdateAmosFutureFlights>(context).execute(
                              authorization: authorization.key,
                              data: picked.data,
                            );
                          }
                        : null,
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

final class _Picker extends StatelessWidget {
  const _Picker();

  @override
  Widget build(context) {
    return BlocListener<usecases.OpenFutureFlights, usecases.OpenFutureFlightsState>(
      listener: (context, state) {
        if (state is usecases.OpenFutureFlightsStateFailure) {
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
      },
      child: BlocBuilder<usecases.OpenFutureFlights, usecases.OpenFutureFlightsState>(
        bloc: BlocProvider.of<usecases.OpenFutureFlights>(context),
        builder: (context, state) {
          final bool isPrepared = state is usecases.OpenFutureFlightsStateSuccess;
          final bool isLoading = state is usecases.OpenFutureFlightsStateLoading;

          return ListTile(
            contentPadding: EdgeInsets.zero,
            title: Text(
              isPrepared ? state.name : "Please Pick a File",
              style: TextStyle(
                fontFamily: "CascadiaCode",
                color: isPrepared ? Colors.black : Colors.grey,
              ),
            ),
            trailing: isLoading
                ? const common_ui.ProgressIndicator(
                    scale: 0.50,
                  )
                : IconButton(
                    onPressed: () {
                      BlocProvider.of<usecases.OpenFutureFlights>(context).execute();
                    },
                    icon: const Icon(Icons.file_open_outlined),
                  ),
          );
        },
      ),
    );
  }
}
