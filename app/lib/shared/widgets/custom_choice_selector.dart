import 'package:flutter/material.dart';
import 'package:app/shared/theme/app_theme.dart';

class CustomChoiceSelector extends StatelessWidget {
  final String title;
  final String currentValue;
  final List<ChoiceOption> options;
  final ValueChanged<String> onSelected;
  final bool multipleChoice; // Per future estensioni

  const CustomChoiceSelector({
    super.key,
    required this.title,
    required this.currentValue,
    required this.options,
    required this.onSelected,
    this.multipleChoice = false,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => _showDialog(context),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          Text(
            _getCurrentLabel(),
            style: const TextStyle(
              color: AppTheme.textSecondary,
              fontSize: 14,
              fontWeight: FontWeight.w600,
            ),
          ),
          const SizedBox(width: 8),
          const Icon(
            Icons.chevron_right,
            color: AppTheme.textSecondary,
            size: 20,
          ),
        ],
      ),
    );
  }

  String _getCurrentLabel() {
    final option = options.firstWhere(
      (opt) => opt.value == currentValue,
      orElse: () => options.first,
    );
    return option.label;
  }

  void _showDialog(BuildContext context) {
    showDialog(
      context: context,
      builder:
          (context) => AlertDialog(
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(20),
            ),
            title: Text(
              title,
              style: const TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.w700,
                color: AppTheme.textPrimary,
              ),
            ),
            contentPadding: const EdgeInsets.symmetric(vertical: 20),
            content: Column(
              mainAxisSize: MainAxisSize.min,
              children:
                  options
                      .map((option) => _buildOption(context, option))
                      .toList(),
            ),
          ),
    );
  }

  Widget _buildOption(BuildContext context, ChoiceOption option) {
    final isSelected = option.value == currentValue;

    return InkWell(
      onTap: () {
        onSelected(option.value);
        Navigator.pop(context);
      },
      child: Container(
        padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
        child: Row(
          children: [
            // Icona opzione (opzionale)
            if (option.icon != null) ...[
              Icon(
                option.icon,
                color:
                    isSelected ? AppTheme.primaryColor : AppTheme.textSecondary,
                size: 24,
              ),
              const SizedBox(width: 16),
            ],

            // Label
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    option.label,
                    style: TextStyle(
                      fontSize: 16,
                      fontWeight:
                          isSelected ? FontWeight.w600 : FontWeight.w500,
                      color:
                          isSelected
                              ? AppTheme.primaryColor
                              : AppTheme.textPrimary,
                    ),
                  ),
                  if (option.description != null) ...[
                    const SizedBox(height: 4),
                    Text(
                      option.description!,
                      style: const TextStyle(
                        fontSize: 13,
                        color: AppTheme.textSecondary,
                      ),
                    ),
                  ],
                ],
              ),
            ),

            // Check indicator
            if (isSelected)
              const Icon(
                Icons.check_circle,
                color: AppTheme.primaryColor,
                size: 24,
              )
            else
              Container(
                width: 24,
                height: 24,
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  border: Border.all(
                    color: AppTheme.textSecondary.withOpacity(0.3),
                    width: 2,
                  ),
                ),
              ),
          ],
        ),
      ),
    );
  }
}

// Classe per le opzioni
class ChoiceOption {
  final String value;
  final String label;
  final String? description;
  final IconData? icon;

  const ChoiceOption({
    required this.value,
    required this.label,
    this.description,
    this.icon,
  });
}
