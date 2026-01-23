import 'package:app/features/profile/widgets/settings/profile_logout_section.dart';
import 'package:app/features/profile/widgets/settings/profile_stats_card_section.dart';
import 'package:flutter/material.dart';
import 'package:app/features/profile/models/profile.dart';
import 'package:app/features/profile/widgets/profile_header.dart';
import 'package:app/features/profile/widgets/settings/profile_general_settings.dart';
import 'package:app/features/profile/widgets/settings/profile_travel_settings.dart';
import 'package:app/features/profile/widgets/settings/profile_calendar_settings.dart';
import 'package:app/features/profile/widgets/settings/profile_privacy_settings.dart';
import 'package:app/features/profile/widgets/settings/profile_other_settings.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: SingleChildScrollView(
          child: Column(
            children: [
              // Header con avatar, nome, email
              ProfileHeader(
                profile: Profile(
                  name: 'Mario',
                  surname: 'Rossi',
                  email: 'email.92@gmail.com',
                ),
              ),

              // Statistiche
              const ProfileStatsCardSection(
                totalTrips: 247,
                punctualityPercentage: 94,
                timeSaved: '12h',
              ),

              // Sezioni impostazioni
              const ProfileGeneralSettings(),
              const ProfileTravelSettings(),
              const ProfileCalendarSettings(),
              const ProfilePrivacySettings(),
              const ProfileOtherSettings(),

              // Bottone logout
              ProfileLogoutSection(
                onLogout: () {
                  // TODO: Implementa logout
                  print('Logout');
                },
              ),

              const SizedBox(height: 20),
            ],
          ),
        ),
      ),
    );
  }
}
