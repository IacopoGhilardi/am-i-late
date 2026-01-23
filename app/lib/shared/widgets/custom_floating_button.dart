import 'package:app/shared/theme/app_theme.dart';
import 'package:flutter/material.dart';

class CustomFloatingButton extends StatelessWidget {
  final Function onTap;
  final Widget widget;

  const CustomFloatingButton({
    super.key,
    required this.onTap,
    required this.widget,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        gradient: const LinearGradient(
          begin: Alignment.topLeft,
          end: Alignment.bottomRight,
          colors: [AppTheme.primaryColor, AppTheme.secondaryColor],
        ),
        borderRadius: BorderRadius.circular(25),
      ),
      child: Material(
        color: Colors.transparent,
        borderRadius: BorderRadius.circular(16),
        child: InkWell(
          borderRadius: BorderRadius.circular(16),
          onTap: () {
            onTap();
          },
          child: widget,
        ),
      ),
    );
  }
}
