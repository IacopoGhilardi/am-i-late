// lib/shared/widgets/priority_selector.dart
import 'package:app/shared/widgets/custom_icon_selector.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class PrioritySelector extends StatelessWidget {
  final String initialValue;
  final ValueChanged<String> onChanged;

  const PrioritySelector({
    super.key,
    this.initialValue = 'medium',
    required this.onChanged,
  });

  @override
  Widget build(BuildContext context) {
    return CustomIconSelector(
      options: const [
        IconSelectorOption(
          value: 'low',
          icon: LucideIcons.arrowDown,
          label: 'Bassa',
        ),
        IconSelectorOption(
          value: 'medium',
          icon: LucideIcons.minus,
          label: 'Media',
        ),
        IconSelectorOption(
          value: 'high',
          icon: LucideIcons.arrowUp,
          label: 'Alta',
        ),
        IconSelectorOption(
          value: 'urgent',
          icon: LucideIcons.alertCircle,
          label: 'Urgente',
        ),
      ],
      initialValue: initialValue,
      onChanged: onChanged,
    );
  }
}
