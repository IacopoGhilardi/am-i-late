import 'package:app/shared/theme/app_theme.dart';
import 'package:flutter/material.dart';

class ProfileGeneralSetting extends StatefulWidget {
  const ProfileGeneralSetting({super.key});

  @override
  State<ProfileGeneralSetting> createState() => _ProfileGeneralSettingState();
}

class _ProfileGeneralSettingState extends State<ProfileGeneralSetting> {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      padding: const EdgeInsets.all(30),
      child: SafeArea(
        child: Column(
          children: [
            Row(
              children: [
                Icon(Icons.engineering),
                SizedBox(width: 12),
                Text(
                  'impostazioni generali'.toUpperCase(),
                  style: TextStyle(
                    color: AppTheme.textSecondary,
                    fontWeight: FontWeight.bold,
                    fontSize: 16,
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
