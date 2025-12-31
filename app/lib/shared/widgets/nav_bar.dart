import 'package:app/features/profile/widgets/profile_page.dart';
import 'package:app/main.dart';
import 'package:app/shared/models/nav_item_data.dart';
import 'package:app/shared/widgets/nav_item.dart';
import 'package:flutter/material.dart';

class NavBar extends StatelessWidget {
  final int currentIndex;
  final Function(int) onTap;

  const NavBar({super.key, required this.currentIndex, required this.onTap});

  static final List<NavItemData> _navItems = [
    NavItemData(
      icon: Icons.home_rounded,
      label: 'Home',
      page: const ProfilePage(),
    ),
    NavItemData(
      icon: Icons.location_on_rounded,
      label: 'Destinazioni',
      page: const ProfilePage(),
    ),
    NavItemData(
      icon: Icons.calendar_today_rounded,
      label: 'Calendario',
      page: const ProfilePage(),
      badgeCount: 3,
    ),
    NavItemData(
      icon: Icons.person_rounded,
      label: 'Profilo',
      page: const ProfilePage(),
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white,
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.05),
            blurRadius: 10,
            offset: const Offset(0, -2),
          ),
        ],
      ),
      child: SafeArea(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 20, vertical: 8),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: List.generate(
              _navItems.length,
              (index) => NavItem(
                icon: _navItems[index].icon,
                label: _navItems[index].label,
                isSelected: currentIndex == index,
                badgeCount: _navItems[index].badgeCount,
                onTap: () => onTap(index),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
