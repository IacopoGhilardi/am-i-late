import 'package:app/features/profile/widgets/settings/profile_setting_item.dart';
import 'package:app/features/profile/widgets/settings/profile_setting_section.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class ProfileOtherSettings extends StatelessWidget {
  const ProfileOtherSettings({super.key});

  @override
  Widget build(BuildContext context) {
    return ProfileSettingsSection(
      title: 'Altro',
      icon: LucideIcons.badgeInfo,
      children: [
        ProfileSettingItem(
          label: 'Centro assistenza',
          icon: LucideIcons.helpCircle,
          onTap: () {
            // TODO: Apri centro assistenza
          },
        ),
        ProfileSettingItem(
          label: 'Termini e condizioni',
          icon: LucideIcons.folder,
          onTap: () {
            // TODO: Apri termini
          },
        ),
        ProfileSettingItem(
          label: 'Informazioni app',
          icon: LucideIcons.info,
          description: 'Versione 1.0.0',
          onTap: () {
            // TODO: Mostra info app
          },
        ),
      ],
    );
  }
}
