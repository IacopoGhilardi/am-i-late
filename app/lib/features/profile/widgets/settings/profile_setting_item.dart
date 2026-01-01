import 'package:flutter/material.dart';

class ProfileSettingItem extends StatelessWidget {
  final String label;
  final IconData? icon;
  final VoidCallback onTap;
  final Widget? trailing;

  const ProfileSettingItem({
    super.key,
    required this.label,
    this.icon,
    required this.onTap,
    this.trailing,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      behavior: HitTestBehavior.opaque,
      child: Container(
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 16),
        child: Row(
          children: [
            if (icon != null) Icon(icon, size: 24),
            if (icon != null) const SizedBox(width: 12),
            Expanded(child: Text(label, style: TextStyle(fontSize: 16))),
            trailing ?? Icon(Icons.chevron_right),
          ],
        ),
      ),
    );
  }
}
