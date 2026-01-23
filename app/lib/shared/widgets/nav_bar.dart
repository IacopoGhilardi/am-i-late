import 'package:app/features/calendar/widgets/calendar_page.dart';
import 'package:app/features/destination/widgets/destination_page.dart';
import 'package:app/features/home/widgets/home_page.dart';
import 'package:app/features/profile/widgets/profile_page.dart';
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
      page: const HomePage(),
    ),
    NavItemData(
      icon: Icons.calendar_today_rounded,
      label: 'Calendario',
      page: const CalendarPage(),
      badgeCount: 3,
    ),
    // NavItemData(
    //   icon: Icons.location_on_rounded,
    //   label: 'Destinazioni',
    //   page: const DestinationPage(),
    // ),
    NavItemData(
      icon: Icons.person_rounded,
      label: 'Profilo',
      page: const ProfilePage(),
    ),
  ];

  static List<Widget> get pages => _navItems.map((item) => item.page).toList();

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
          padding: EdgeInsets.symmetric(horizontal: 0, vertical: 8),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: List.generate(
              _navItems.length,
              (index) => NavItem(
                icon: _navItems[index].icon,
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
