part of "view.dart";

final class _SignIn extends StatefulWidget {
  const _SignIn();

  @override
  State<_SignIn> createState() => _SignInState();
}

final class _SignInState extends State<_SignIn> {
  final TextEditingController controller = TextEditingController();
  final formKey = GlobalKey<FormState>();

  @override
  Widget build(context) {
    return Form(
      key: formKey,
      autovalidateMode: AutovalidateMode.onUserInteraction,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          common_ui.PasswordInput(controller: controller, width: 512),
          const SizedBox(height: 16),
          common_ui.ActionButton(
            width: 512,
            title: "AUTHORIZE",
            onPressed: () {
              if (!formKey.currentState!.validate()) {
                common_ui.showError(context, "Invalid Password");
                return;
              }

              context.read<GetAuth>().execute(controller.text);
            },
          ),
        ],
      ),
    );
  }
}
