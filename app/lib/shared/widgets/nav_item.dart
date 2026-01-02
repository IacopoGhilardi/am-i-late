import 'package:app/shared/theme/app_theme.dart';
import 'package:flutter/material.dart';

class NavItem extends StatelessWidget {
  final IconData icon;
  final String label;
  final bool isSelected;
  final VoidCallback onTap;
  final int? badgeCount;

  const NavItem({
    required this.icon,
    required this.label,
    required this.isSelected,
    required this.onTap,
    this.badgeCount,
  });

  @override
  Widget build(BuildContext context) {
    final color =
        isSelected
            ? AppTheme.primaryColor
            : AppTheme.textSecondary.withOpacity(0.6);

    return GestureDetector(
      onTap: onTap,
      behavior: HitTestBehavior.opaque,
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 12),
        child: Stack(
          clipBehavior: Clip.none,
          children: [
            Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                AnimatedContainer(
                  duration: const Duration(milliseconds: 200),
                  curve: Curves.easeOut,
                  child: Icon(icon, color: color, size: isSelected ? 30 : 25),
                ),
                const SizedBox(height: 4),
                AnimatedDefaultTextStyle(
                  style: TextStyle(
                    color: color,
                    fontSize: isSelected ? 11.5 : 11,
                    fontWeight:
                        isSelected ? FontWeight.bold : FontWeight.normal,
                  ),
                  duration: const Duration(milliseconds: 200),
                  child: Text(label),
                ),
              ],
            ),
            if (badgeCount != null && badgeCount! > 0)
              Positioned(
                right: 0,
                top: -2,
                child: Container(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 6,
                    vertical: 2,
                  ),
                  decoration: BoxDecoration(
                    color: AppTheme.errorColor,
                    borderRadius: BorderRadius.circular(10),
                    boxShadow: [
                      BoxShadow(
                        color: AppTheme.errorColor.withOpacity(0.3),
                        blurRadius: 4,
                        offset: const Offset(0, 2),
                      ),
                    ],
                  ),
                  constraints: const BoxConstraints(
                    minHeight: 18,
                    minWidth: 18,
                  ),
                  child: Text(
                    badgeCount! > 99 ? '99+' : badgeCount.toString(),
                    style: const TextStyle(
                      color: AppTheme.surfaceColor,
                      fontSize: 10,
                      fontWeight: FontWeight.bold,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ),
              ),
          ],
        ),
      ),
    );
  }
}
