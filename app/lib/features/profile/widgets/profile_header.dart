import 'package:app/features/profile/models/profile.dart';
import 'package:app/shared/theme/app_theme.dart';
import 'package:flutter/material.dart';

class ProfileHeader extends StatefulWidget {
  final Profile profile;
  const ProfileHeader({super.key, required this.profile});

  @override
  State<ProfileHeader> createState() => _ProfileHeaderState();
}

class _ProfileHeaderState extends State<ProfileHeader> {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      padding: const EdgeInsets.all(30),
      child: Column(
        children: [
          CircleAvatar(
            radius: 75,
            backgroundImage: NetworkImage(
              'https://media.tenor.com/kRYvat8fCCMAAAAe/mr-bean-bean.png',
            ),
          ),
          SizedBox(height: 20),
          Text(
            '${widget.profile.name} ${widget.profile.surname}',
            style: TextStyle(fontSize: 28, fontWeight: FontWeight.bold),
          ),
          SizedBox(height: 3),
          Text(
            '${widget.profile.email}',
            style: TextStyle(fontSize: 16, color: AppTheme.textSecondary),
          ),
        ],
      ),
    );
  }
}
