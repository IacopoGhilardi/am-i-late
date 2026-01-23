import 'package:flutter/material.dart';
import 'package:app/shared/theme/app_theme.dart';

class ProfileStatsCardSection extends StatelessWidget {
  final int totalTrips;
  final int punctualityPercentage;
  final String timeSaved;

  const ProfileStatsCardSection({
    super.key,
    required this.totalTrips,
    required this.punctualityPercentage,
    required this.timeSaved,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 20),
      child: Row(
        children: [
          Expanded(
            child: _StatCard(value: totalTrips.toString(), label: 'Viaggi'),
          ),
          const SizedBox(width: 12),
          Expanded(
            child: _StatCard(
              value: '$punctualityPercentage%',
              label: 'Puntualità',
            ),
          ),
          const SizedBox(width: 12),
          Expanded(child: _StatCard(value: timeSaved, label: 'Risparmiati')),
        ],
      ),
    );
  }
}

class _StatCard extends StatelessWidget {
  final String value;
  final String label;

  const _StatCard({required this.value, required this.label});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(16),
        border: Border.all(color: const Color(0xFFE5E7EB), width: 2),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.04),
            blurRadius: 8,
            offset: const Offset(0, 2),
          ),
        ],
      ),
      child: Column(
        children: [
          Text(
            value,
            style: const TextStyle(
              fontSize: 24,
              fontWeight: FontWeight.w700,
              color: AppTheme.textPrimary,
            ),
          ),
          const SizedBox(height: 4),
          Text(
            label,
            style: const TextStyle(
              fontSize: 11,
              fontWeight: FontWeight.w600,
              color: AppTheme.textSecondary,
              letterSpacing: 0.5,
            ),
            textAlign: TextAlign.center,
          ),
        ],
      ),
    );
  }
}
