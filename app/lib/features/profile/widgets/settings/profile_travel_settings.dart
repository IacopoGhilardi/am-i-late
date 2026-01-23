import 'package:app/core/services/settings_service.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_item.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_section.dart';
import 'package:app/shared/widgets/custom_choice_selector.dart';
import 'package:app/shared/widgets/custom_switch.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class ProfileTravelSettings extends StatefulWidget {
  const ProfileTravelSettings({super.key});

  @override
  State<ProfileTravelSettings> createState() => _ProfileTravelSettingsState();
}

class _ProfileTravelSettingsState extends State<ProfileTravelSettings> {
  final _settings = SettingsService();

  bool _avoidPaymentRoads = true;
  String _securityDelay = '10';
  String _defaultTransport = 'auto';

  @override
  void initState() {
    super.initState();
    _loadSettings();
  }

  void _loadSettings() {
    setState(() {
      _avoidPaymentRoads =
          _settings.get<bool>('avoid_payment_roads', defaultValue: true)!;

      _securityDelay =
          _settings.get<String>('security_delay', defaultValue: '10')!;

      _defaultTransport =
          _settings.get<String>('default_transport', defaultValue: 'auto')!;
    });
  }

  @override
  Widget build(BuildContext context) {
    return ProfileSettingsSection(
      title: 'preferenze viaggio'.toUpperCase(),
      icon: LucideIcons.car,
      children: [
        ProfileSettingItem(
          label: 'Evida pedaggi',
          onTap: () {},
          icon: LucideIcons.circleDollarSign,
          trailing: CustomSwitch(
            value: _avoidPaymentRoads,
            onChanged: (value) {
              setState(() => _avoidPaymentRoads = value);
              _settings.set('aovid_payment_roads', value);
            },
          ),
        ),
        ProfileSettingItem(
          label: 'Mezzo predefinito',
          icon: LucideIcons.car,
          trailing: CustomChoiceSelector(
            title: 'Seleziona mezzo',
            currentValue: _defaultTransport,
            options: const [
              ChoiceOption(value: 'car', label: 'Auto', icon: LucideIcons.car),
              ChoiceOption(
                value: 'transit',
                label: 'Mezzi pubblici',
                icon: LucideIcons.train,
              ),
              ChoiceOption(
                value: 'bike',
                label: 'Bicicletta',
                icon: LucideIcons.bike,
              ),
              ChoiceOption(
                value: 'walk',
                label: 'A piedi',
                icon: LucideIcons.footprints,
              ),
            ],
            onSelected: (value) {
              setState(() => _defaultTransport = value);
              _settings.set('default_transport', value);
            },
          ),
          onTap: () {},
        ),
        ProfileSettingItem(
          label: 'Buffer di sicurezza',
          description: 'Tempo extra per imprevisti',
          icon: LucideIcons.clock,
          trailing: CustomChoiceSelector(
            title: '0 min',
            currentValue: _securityDelay,
            options: const [
              ChoiceOption(value: '0', label: '0 min'),
              ChoiceOption(value: '10', label: '10 min'),
              ChoiceOption(value: '20', label: '20 min'),
              ChoiceOption(value: '30', label: '30 min'),
            ],
            onSelected: (value) {
              setState(() => _securityDelay = value);
              _settings.set('security_delay', value);
            },
          ),
          onTap: () {},
        ),
      ],
    );
  }
}
