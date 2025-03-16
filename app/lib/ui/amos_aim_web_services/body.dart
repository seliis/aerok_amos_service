part of "view.dart";

final class _Body extends StatefulWidget {
  const _Body({required this.token});

  final String token;

  @override
  State<_Body> createState() => _BodyState();
}

final class _BodyState extends State<_Body> with TickerProviderStateMixin {
  late final TabController controller;

  @override
  void initState() {
    super.initState();
    controller = TabController(length: 2, vsync: this);
  }

  @override
  void dispose() {
    controller.dispose();
    super.dispose();
  }

  @override
  Widget build(context) {
    final theme = Theme.of(context);

    return Column(
      children: [
        TabBar(
          controller: controller,
          tabs: [
            Tab(
              child: Text(
                "Import Currency",
                style: theme.textTheme.titleMedium?.copyWith(
                  fontWeight: FontWeight.w200,
                ),
              ),
            ),
            Tab(
              child: Text(
                "Transfer Future Flights",
                style: theme.textTheme.titleMedium?.copyWith(
                  fontWeight: FontWeight.w200,
                ),
              ),
            ),
          ],
        ),
        Expanded(
          child: TabBarView(
            controller: controller,
            children: [
              _ImportCurrency(token: widget.token),
              _TransferFutureFlights(token: widget.token),
            ],
          ),
        ),
      ],
    );
  }
}
