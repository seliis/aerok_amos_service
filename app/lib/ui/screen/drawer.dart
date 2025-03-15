part of "screen.dart";

final class _Drawer extends StatelessWidget {
  const _Drawer();

  @override
  Widget build(context) {
    return Drawer(
      child: Padding(
        padding: EdgeInsets.all(16),
        child: Column(
          children: [
            _MenuGroup(
              groupName: "Public Services",
              menuItems: [
                _Menu(
                  title: "Exchange Rates",
                  icon: Icons.monetization_on_outlined,
                  onTap: () {},
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}

final class _MenuGroup extends StatelessWidget {
  const _MenuGroup({required this.groupName, required this.menuItems});

  final String groupName;
  final List<_Menu> menuItems;

  @override
  Widget build(context) {
    final theme = Theme.of(context);

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          groupName,
          style: theme.textTheme.bodySmall?.copyWith(
            fontWeight: FontWeight.w200,
          ),
        ),
        SizedBox(height: 4),
        Column(children: menuItems),
      ],
    );
  }
}

final class _Menu extends StatelessWidget {
  const _Menu({required this.title, required this.icon, required this.onTap});

  final String title;
  final IconData icon;
  final void Function() onTap;

  @override
  Widget build(context) {
    final theme = Theme.of(context);

    return ListTile(
      contentPadding: EdgeInsets.zero,
      title: Text(
        title,
        style: theme.textTheme.titleSmall?.copyWith(
          fontWeight: FontWeight.w200,
        ),
      ),
      leading: Icon(icon),
      onTap: onTap,
    );
  }
}
