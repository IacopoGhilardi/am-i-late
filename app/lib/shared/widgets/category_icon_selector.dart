import 'package:app/shared/widgets/custom_icon_selector.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class CategoryIconSelector extends StatelessWidget {
  final String initialValue;
  final ValueChanged<String> onChanged;

  const CategoryIconSelector({
    super.key,
    this.initialValue = 'mapPin',
    required this.onChanged,
  });

  @override
  Widget build(BuildContext context) {
    return CustomIconSelector(
      options: const [
        IconSelectorOption(
          value: 'home',
          icon: LucideIcons.home,
          label: 'Casa',
        ),
        IconSelectorOption(
          value: 'briefcase',
          icon: LucideIcons.briefcase,
          label: 'Lavoro',
        ),
        IconSelectorOption(
          value: 'dumbbell',
          icon: LucideIcons.dumbbell,
          label: 'Palestra',
        ),
        IconSelectorOption(
          value: 'shoppingCart',
          icon: LucideIcons.shoppingCart,
          label: 'Shopping',
        ),
      ],
      initialValue: initialValue,
      onChanged: onChanged,
    );
  }
}
