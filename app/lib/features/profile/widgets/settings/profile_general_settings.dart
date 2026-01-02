import 'package:app/core/services/settings_service.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_item.dart';
import 'package:app/shared/theme/app_theme.dart';
import 'package:app/shared/widgets/custom_switch.dart';
import 'package:flutter/material.dart';

class ProfileGeneralSettings extends StatefulWidget {
  const ProfileGeneralSettings({super.key});

  @override
  State<ProfileGeneralSettings> createState() => _ProfileGeneralSettingsState();
}

class _ProfileGeneralSettingsState extends State<ProfileGeneralSettings> {
  final _settings = SettingsService();

  bool _notificationEnabled = true;
  bool _locationAlwaysOn = true;
  String _themeMode = 'auto';

  @override
  void initState() {
    super.initState();
    _loadSettings();
  }

  void _loadSettings() {
    setState(() {
      _notificationEnabled =
          _settings.get<bool>('notification_enabled', defaultValue: true)!;

      _locationAlwaysOn =
          _settings.get<bool>('location_always_on', defaultValue: true)!;

      _themeMode = _settings.get<String>('theme_mode', defaultValue: 'auto')!;
    });
  }

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
            SizedBox(height: 16),
            Container(
              decoration: BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.circular(20),
                border: Border.all(
                  color: AppTheme.textSecondary.withOpacity(0.2),
                  width: 2,
                ),
              ),
              child: Column(
                children: [
                  ProfileSettingItem(
                    label: 'Notifiche push',
                    icon: Icons.notifications_rounded,
                    description: 'Ricevi avvisi in tempo reale',
                    trailing: CustomSwitch(
                      value: _notificationEnabled,
                      onChanged: (value) {
                        setState(() => _notificationEnabled = value);
                        _settings.set('notifications_enabled', value);
                      },
                    ),
                    onTap: () {},
                  ),
                  Divider(height: 1, color: Color(0xFFF3F4F6)),

                  ProfileSettingItem(
                    label: 'Posizione sempre attiva',
                    icon: Icons.location_city,
                    trailing: CustomSwitch(
                      value: _locationAlwaysOn,
                      onChanged: (value) {
                        setState(() => _locationAlwaysOn = value);
                        _settings.set('location_always_on', value);
                      },
                    ),
                    onTap: () {},
                  ),
                  Divider(height: 1, color: Color(0xFFF3F4F6)),

                  ProfileSettingItem(
                    label: 'Tema scuro',
                    icon: Icons.dark_mode_rounded,
                    trailing: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Text(
                          _themeMode,
                          style: TextStyle(
                            color: Color(0xFF6C757D),
                            fontSize: 14,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                        const SizedBox(width: 8),
                        Icon(
                          Icons.chevron_right,
                          color: Color(0xFF6C757D),
                          size: 20,
                        ),
                      ],
                    ),
                    onTap: () {
                      _showThemeDialog();
                    },
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

  void _showThemeDialog() {
    showDialog(
      context: context,
      builder:
          (context) => AlertDialog(
            title: Text('Seleziona tema'),
            content: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                _buildThemeOption('Auto'),
                _buildThemeOption('Chiaro'),
                _buildThemeOption('Scuro'),
              ],
            ),
          ),
    );
  }

  Widget _buildThemeOption(String theme) {
    return RadioListTile<String>(
      title: Text(theme),
      value: theme.toLowerCase(),
      groupValue: _themeMode,
      onChanged: (value) {
        setState(() => _themeMode = value!);
        _settings.set('theme_mode', value!);
        Navigator.pop(context);
      },
    );
  }
}
