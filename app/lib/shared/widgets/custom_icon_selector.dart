import 'package:app/shared/theme/app_theme.dart';
import 'package:flutter/material.dart';

class IconSelectorOption {
  final String value;
  final IconData icon;
  final String label;

  const IconSelectorOption({
    required this.value,
    required this.icon,
    required this.label,
  });
}

class CustomIconSelector extends StatefulWidget {
  final List<IconSelectorOption> options;
  final String initialValue;
  final ValueChanged<String> onChanged;

  const CustomIconSelector({
    super.key,
    required this.options,
    required this.initialValue,
    required this.onChanged,
  });

  @override
  State<CustomIconSelector> createState() => _CustomIconSelectorState();
}

class _CustomIconSelectorState extends State<CustomIconSelector> {
  late String _selectedValue;

  @override
  void initState() {
    super.initState();
    _selectedValue = widget.initialValue;
  }

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        for (var option in widget.options) ...[
          Expanded(child: _buildOption(option)),
          if (option != widget.options.last) const SizedBox(width: 12),
        ],
      ],
    );
  }

  Widget _buildOption(IconSelectorOption option) {
    final isSelected = _selectedValue == option.value;

    return GestureDetector(
      onTap: () {
        setState(() => _selectedValue = option.value);
        widget.onChanged(option.value);
      },
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 16),
        decoration: BoxDecoration(
          gradient:
              isSelected
                  ? const LinearGradient(
                    colors: [AppTheme.primaryColor, AppTheme.secondaryColor],
                  )
                  : null,
          color: isSelected ? null : Colors.white,
          borderRadius: BorderRadius.circular(16),
          border: Border.all(
            color: isSelected ? Colors.transparent : const Color(0xFFE5E7EB),
            width: 2,
          ),
        ),
        child: Column(
          children: [
            Icon(
              option.icon,
              color: isSelected ? Colors.white : AppTheme.textSecondary,
              size: 24,
            ),
            const SizedBox(height: 8),
            Text(
              option.label,
              style: TextStyle(
                fontSize: 12,
                fontWeight: FontWeight.w600,
                color: isSelected ? Colors.white : AppTheme.textSecondary,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
