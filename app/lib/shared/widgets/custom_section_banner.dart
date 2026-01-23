import 'package:app/shared/theme/app_theme.dart';
import 'package:flutter/widgets.dart';

class CustomSectionBanner extends StatelessWidget {
  final Widget widget;
  const CustomSectionBanner({super.key, required this.widget});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(20),
      decoration: BoxDecoration(
        gradient: LinearGradient(
          begin: Alignment.topLeft,
          end: Alignment.bottomRight,
          colors: [AppTheme.primaryColor, AppTheme.secondaryColor],
        ),
        borderRadius: BorderRadius.circular(10),
      ),
      child: widget,
    );
  }
}
