import "package:flutter/material.dart";

import "package:app/base/__index.dart" as base;

final class MasterDrawer extends StatelessWidget {
  const MasterDrawer({super.key});

  @override
  Widget build(context) {
    return Drawer(
      child: Column(
        children: [
          const DrawerHeader(
            child: Text("Aero K Airlines"),
          ),
          Column(
            children: base.Route.values.map((route) {
              return ListTile(
                title: Text(route.title),
                leading: Icon(route.icon),
                onTap: () {
                  Navigator.of(context).pushReplacementNamed(route.path);
                },
              );
            }).toList(),
          ),
        ],
      ),
    );
  }
}
