import 'package:flutter/widgets.dart';

class NavItemData {
  final IconData icon;
  final String label;
  final Widget page;
  final int? badgeCount;

  const NavItemData({
    required this.icon,
    required this.label,
    required this.page,
    this.badgeCount,
  })
}
