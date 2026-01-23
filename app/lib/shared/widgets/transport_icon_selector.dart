// lib/shared/widgets/transport_selector.dart
import 'package:app/shared/widgets/custom_icon_selector.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class TransportSelector extends StatelessWidget {
  final String initialValue;
  final ValueChanged<String> onChanged;

  const TransportSelector({
    super.key,
    this.initialValue = 'driving',
    required this.onChanged,
  });

  @override
  Widget build(BuildContext context) {
    return CustomIconSelector(
      options: const [
        IconSelectorOption(
          value: 'driving',
          icon: LucideIcons.car,
          label: 'Auto',
        ),
        IconSelectorOption(
          value: 'walking',
          icon: LucideIcons.footprints,
          label: 'A piedi',
        ),
        IconSelectorOption(
          value: 'transit',
          icon: LucideIcons.bus,
          label: 'Mezzi',
        ),
        IconSelectorOption(
          value: 'bicycling',
          icon: LucideIcons.bike,
          label: 'Bici',
        ),
      ],
      initialValue: initialValue,
      onChanged: onChanged,
    );
  }
}
