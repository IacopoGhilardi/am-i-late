import 'package:flutter/material.dart';
import 'package:app/shared/theme/app_theme.dart';

class ProfileSettingsSection extends StatelessWidget {
  final String title;
  final IconData icon;
  final List<Widget> children;

  const ProfileSettingsSection({
    super.key,
    required this.title,
    required this.icon,
    required this.children,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      padding: const EdgeInsets.all(25),
      child: Column(
        children: [
          // Title
          Row(
            children: [
              Icon(icon, size: 26, color: AppTheme.textSecondary),
              const SizedBox(width: 12),
              Text(
                title.toUpperCase(),
                style: const TextStyle(
                  color: AppTheme.textSecondary,
                  fontWeight: FontWeight.bold,
                  fontSize: 16,
                  letterSpacing: 1,
                ),
              ),
            ],
          ),
          const SizedBox(height: 16),

          // Container with children
          Container(
            decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(20),
              border: Border.all(
                color: AppTheme.textSecondary.withOpacity(0.2),
                width: 2,
              ),
            ),
            child: Column(children: _buildChildrenWithDividers()),
          ),
        ],
      ),
    );
  }

  List<Widget> _buildChildrenWithDividers() {
    final List<Widget> widgetsWithDividers = [];

    for (int i = 0; i < children.length; i++) {
      widgetsWithDividers.add(children[i]);

      if (i < children.length - 1) {
        widgetsWithDividers.add(
          const Divider(height: 1, color: Color(0xFFF3F4F6)),
        );
      }
    }

    return widgetsWithDividers;
  }
}
