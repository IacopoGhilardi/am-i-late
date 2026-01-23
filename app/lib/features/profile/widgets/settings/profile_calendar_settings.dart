import 'package:app/core/services/settings_service.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_item.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_section.dart';
import 'package:app/shared/theme/app_theme.dart';
import 'package:app/shared/widgets/custom_switch.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class ProfileCalendarSettings extends StatefulWidget {
  const ProfileCalendarSettings({super.key});

  @override
  State<ProfileCalendarSettings> createState() =>
      _ProfileCalendarSettingsState();
}

class _ProfileCalendarSettingsState extends State<ProfileCalendarSettings> {
  final _settings = SettingsService();

  bool _calendarSynced = false;
  bool _autoPlanningEnabled = true;
  bool _smartRemindersEnabled = true;

  @override
  void initState() {
    super.initState();
    _loadSettings();
  }

  void _loadSettings() {
    setState(() {
      _calendarSynced =
          _settings.get<bool>('calendar_synced', defaultValue: false)!;
      _autoPlanningEnabled =
          _settings.get<bool>('auto_planning_enabled', defaultValue: true)!;
      _smartRemindersEnabled =
          _settings.get<bool>('smart_reminders_enabled', defaultValue: true)!;
    });
  }

  @override
  Widget build(BuildContext context) {
    return ProfileSettingsSection(
      title: 'Calendario e integrazione',
      icon: LucideIcons.calendarDays,
      children: [
        ProfileSettingItem(
          label: 'Calendario Google',
          icon: LucideIcons.calendarCheck,
          description: _calendarSynced ? 'Sincronizzato' : 'Non sincronizzato',
          trailing: Row(
            mainAxisSize: MainAxisSize.min,
            children: [
              Icon(
                _calendarSynced ? LucideIcons.checkCircle : LucideIcons.xCircle,
                color:
                    _calendarSynced
                        ? AppTheme.successColor
                        : AppTheme.errorColor,
                size: 20,
              ),
              const SizedBox(width: 8),
              const Icon(
                Icons.chevron_right,
                color: Color(0xFF6C757D),
                size: 20,
              ),
            ],
          ),
          onTap: () {
            // TODO: Apri schermata sincronizzazione
            setState(() => _calendarSynced = !_calendarSynced);
            _settings.set('calendar_synced', _calendarSynced);
          },
        ),
        ProfileSettingItem(
          label: 'Pianificazione automatica',
          icon: LucideIcons.sparkles,
          description: 'Crea destinazioni da eventi',
          trailing: CustomSwitch(
            value: _autoPlanningEnabled,
            onChanged: (value) {
              setState(() => _autoPlanningEnabled = value);
              _settings.set('auto_planning_enabled', value);
            },
          ),
          onTap: () {},
        ),
        ProfileSettingItem(
          label: 'Promemoria intelligenti',
          icon: LucideIcons.bellRing,
          trailing: CustomSwitch(
            value: _smartRemindersEnabled,
            onChanged: (value) {
              setState(() => _smartRemindersEnabled = value);
              _settings.set('smart_reminders_enabled', value);
            },
          ),
          onTap: () {},
        ),
      ],
    );
  }
}
