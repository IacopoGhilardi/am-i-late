import 'package:app/features/profile/models/profile.dart';
import 'package:app/features/profile/widgets/settings/profile_general_settings.dart';
import 'package:app/features/profile/widgets/profile_header.dart';
import 'package:flutter/material.dart';

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key});

  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Column(
          children: [
            ProfileHeader(
              profile: Profile(
                name: 'Mario',
                surname: 'Rossi',
                email: 'email.92@gmail.com',
              ),
            ),
            ProfileGeneralSettings(),
          ],
        ),
      ),
    );
  }
}
