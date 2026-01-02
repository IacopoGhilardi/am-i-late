import 'package:flutter/material.dart';

class ProfileSettingItem extends StatelessWidget {
  final String label;
  final String? description;
  final IconData? icon;
  final VoidCallback onTap;
  final Widget? trailing;

  const ProfileSettingItem({
    super.key,
    required this.label,
    this.description,
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
            if (icon != null) ...[
              Icon(icon, size: 24, color: const Color(0xFF2C2C2C)),
              const SizedBox(width: 12),
            ],

            // Label + Description
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    label,
                    style: const TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.w600,
                      color: Color(0xFF2C2C2C),
                    ),
                  ),
                  if (description != null) ...[
                    const SizedBox(height: 2),
                    Text(
                      description!,
                      style: const TextStyle(
                        fontSize: 13,
                        color: Color(0xFF6C757D),
                      ),
                    ),
                  ],
                ],
              ),
            ),

            // Trailing (switch, arrow, text, ecc.)
            if (trailing != null)
              trailing!
            else
              const Icon(
                Icons.chevron_right,
                color: Color(0xFF6C757D),
                size: 20,
              ),
          ],
        ),
      ),
    );
  }
}
