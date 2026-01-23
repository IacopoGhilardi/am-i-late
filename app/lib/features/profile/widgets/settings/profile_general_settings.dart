import 'package:app/core/services/settings_service.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_item.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_section.dart';
import 'package:app/shared/widgets/custom_choice_selector.dart';
import 'package:app/shared/widgets/custom_switch.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

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
    return ProfileSettingsSection(
      title: 'Impostazioni generali',
      icon: LucideIcons.settings,
      children: [
        ProfileSettingItem(
          label: 'Notifiche push',
          icon: LucideIcons.bell,
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

        ProfileSettingItem(
          label: 'Posizione sempre attiva',
          icon: LucideIcons.mapPin,
          trailing: CustomSwitch(
            value: _locationAlwaysOn,
            onChanged: (value) {
              setState(() => _locationAlwaysOn = value);
              _settings.set('location_always_on', value);
            },
          ),
          onTap: () {},
        ),

        ProfileSettingItem(
          label: 'Tema scuro',
          icon: LucideIcons.moon,
          description: 'Modalità automatica',
          trailing: CustomChoiceSelector(
            title: 'Seleziona tema',
            currentValue: _themeMode,
            options: const [
              ChoiceOption(
                value: 'auto',
                label: 'Automatico',
                description: 'Segue le impostazioni di sistema',
                icon: LucideIcons.sunDim,
              ),
              ChoiceOption(
                value: 'light',
                label: 'Chiaro',
                description: 'Tema sempre chiaro',
                icon: LucideIcons.sunDim,
              ),
              ChoiceOption(
                value: 'dark',
                label: 'Scuro',
                description: 'Tema sempre scuro',
                icon: LucideIcons.moon,
              ),
            ],
            onSelected: (value) {
              setState(() => _themeMode = value);
              _settings.set('theme_mode', value);
            },
          ),
          onTap: () {},
        ),
      ],
    );
  }
}
