import "package:flutter/material.dart";

import "package:app/internal/ui/__index.dart" as ui;

final class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  List<NavigationRailDestination> get destinations {
    return [
      const NavigationRailDestination(
        icon: Icon(Icons.currency_exchange),
        label: Text("Currency Rates"),
      ),
      const NavigationRailDestination(
        icon: Icon(Icons.flight),
        label: Text("Future Flights"),
      ),
    ];
  }

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

final class _HomeScreenState extends State<HomeScreen> {
  int selectedIndex = 1;

  @override
  Widget build(context) {
    return Scaffold(
      body: Row(
        children: [
          NavigationRail(
            onDestinationSelected: (index) {
              setState(() {
                selectedIndex = index;
              });
            },
            selectedIndex: selectedIndex,
            destinations: widget.destinations,
          ),
          const VerticalDivider(
            thickness: 0.5,
            width: 1,
          ),
          Expanded(
            child: IndexedStack(
              index: selectedIndex,
              children: const [
                ui.CurrencyScreen(),
                ui.FutureFlightsScreen(),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
