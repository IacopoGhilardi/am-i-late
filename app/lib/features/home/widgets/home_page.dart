import 'package:app/shared/theme/app_theme.dart';
import 'package:app/shared/widgets/page_title.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:intl/intl.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final bool _hasActiveEvents = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppTheme.backgroundColor,
      body: SingleChildScrollView(
        child: Column(
          children: [
            SafeArea(
              child: PageTitle(
                title: 'Bentornato! 👋',
                subtitle: _getGreetingSubtitle(),
              ),
            ),

            const SizedBox(height: 20),

            _hasActiveEvents ? _buildActiveEvents() : _buildEmptyState(),
          ],
        ),
      ),
    );
  }

  String _getGreetingSubtitle() {
    final hour = DateTime.now().hour;
    if (hour < 12) {
      return 'Buongiorno! Ecco i tuoi appuntamenti';
    } else if (hour < 18) {
      return 'Buon pomeriggio! Ecco i tuoi appuntamenti';
    } else {
      return 'Buonasera! Ecco i tuoi appuntamenti';
    }
  }

  Widget _buildEmptyState() {
    return Container(
      margin: const EdgeInsets.all(20),
      padding: const EdgeInsets.all(40),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(20),
        border: Border.all(color: const Color(0xFFE5E7EB), width: 2),
      ),
      child: Column(
        children: [
          Container(
            width: 100,
            height: 100,
            decoration: BoxDecoration(
              gradient: LinearGradient(
                colors: [
                  AppTheme.primaryColor.withOpacity(0.2),
                  AppTheme.secondaryColor.withOpacity(0.2),
                ],
              ),
              shape: BoxShape.circle,
            ),
            child: const Icon(
              LucideIcons.calendarCheck,
              size: 50,
              color: AppTheme.primaryColor,
            ),
          ),
          const SizedBox(height: 24),
          const Text(
            'Nessun appuntamento attivo',
            style: TextStyle(
              fontSize: 20,
              fontWeight: FontWeight.w700,
              color: AppTheme.textPrimary,
            ),
          ),
          const SizedBox(height: 8),
          const Text(
            'Vai al calendario per attivare\nil monitoraggio dei tuoi eventi',
            textAlign: TextAlign.center,
            style: TextStyle(
              fontSize: 14,
              color: AppTheme.textSecondary,
              height: 1.5,
            ),
          ),
          const SizedBox(height: 24),
          _buildInfoCards(),
        ],
      ),
    );
  }

  Widget _buildInfoCards() {
    return Row(
      children: [
        Expanded(
          child: _buildInfoCard(
            icon: LucideIcons.bell,
            title: 'Notifiche',
            subtitle: 'In tempo reale',
            color: AppTheme.primaryColor,
          ),
        ),
        const SizedBox(width: 12),
        Expanded(
          child: _buildInfoCard(
            icon: LucideIcons.navigation,
            title: 'Traffico',
            subtitle: 'Monitorato',
            color: AppTheme.secondaryColor,
          ),
        ),
      ],
    );
  }

  Widget _buildInfoCard({
    required IconData icon,
    required String title,
    required String subtitle,
    required Color color,
  }) {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: color.withOpacity(0.1),
        borderRadius: BorderRadius.circular(16),
        border: Border.all(color: color.withOpacity(0.2), width: 2),
      ),
      child: Column(
        children: [
          Icon(icon, size: 24, color: color),
          const SizedBox(height: 8),
          Text(
            title,
            style: TextStyle(
              fontSize: 12,
              fontWeight: FontWeight.w700,
              color: color,
            ),
          ),
          const SizedBox(height: 2),
          Text(
            subtitle,
            style: TextStyle(fontSize: 10, color: AppTheme.textSecondary),
          ),
        ],
      ),
    );
  }

  Widget _buildActiveEvents() {
    final mockEvents = [
      {
        'title': 'Riunione team',
        'location': 'Via Roma 123, Milano',
        'arrivalTime': DateTime.now().add(const Duration(hours: 2)),
        'departureTime': DateTime.now().add(const Duration(minutes: 78)),
        'status': 'ok', // 'ok', 'warning', 'danger'
        'travelTime': '42 min',
        'traffic': 'moderate',
      },
      {
        'title': 'Dentista',
        'location': 'Piazza Duomo 7, Milano',
        'arrivalTime': DateTime.now().add(const Duration(hours: 5)),
        'departureTime': DateTime.now().add(
          const Duration(hours: 4, minutes: 30),
        ),
        'status': 'ok',
        'travelTime': '15 min',
        'traffic': 'light',
      },
    ];
    return Column(
      children: [
        _buildStatusCard(mockEvents[0]),
        const SizedBox(height: 20),
        ...mockEvents.map((event) => _buildEventCard(event)),
      ],
    );
  }

  Widget _buildStatusCard(Map<String, dynamic> event) {
    final status = event['status'] as String;
    Color backgroundColor;
    Color textColor = Colors.white;
    String message;
    IconData icon;

    switch (status) {
      case 'danger':
        backgroundColor = const Color(0xFFFF6B6B);
        message = '⚠️ SEI IN RITARDO!';
        icon = LucideIcons.alertCircle;
        break;
      case 'warning':
        backgroundColor = const Color(0xFFFFE66D);
        message = 'Parti tra 12 minuti!';
        icon = LucideIcons.clock;
        break;
      default:
        backgroundColor = const Color(0xFF4ECDC4);
        message = '✓ Tutto ok, hai tempo';
        icon = LucideIcons.checkCircle;
    }

    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 20),
      padding: const EdgeInsets.all(24),
      decoration: BoxDecoration(
        gradient: LinearGradient(
          colors: [backgroundColor, backgroundColor.withOpacity(0.8)],
        ),
        borderRadius: BorderRadius.circular(20),
        boxShadow: [
          BoxShadow(
            color: backgroundColor.withOpacity(0.3),
            blurRadius: 12,
            offset: const Offset(0, 6),
          ),
        ],
      ),
      child: Row(
        children: [
          Icon(icon, color: textColor, size: 48),
          const SizedBox(width: 16),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  message,
                  style: TextStyle(
                    color: textColor,
                    fontSize: 20,
                    fontWeight: FontWeight.w700,
                  ),
                ),
                const SizedBox(height: 4),
                Text(
                  event['title'],
                  style: TextStyle(
                    color: textColor.withOpacity(0.9),
                    fontSize: 14,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildEventCard(Map<String, dynamic> event) {
    return Container(
      margin: const EdgeInsets.fromLTRB(20, 0, 20, 16),
      padding: const EdgeInsets.all(20),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(20),
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
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          // Header evento
          Row(
            children: [
              Container(
                width: 48,
                height: 48,
                decoration: BoxDecoration(
                  gradient: const LinearGradient(
                    colors: [AppTheme.primaryColor, AppTheme.secondaryColor],
                  ),
                  borderRadius: BorderRadius.circular(12),
                ),
                child: const Icon(
                  LucideIcons.briefcase,
                  color: Colors.white,
                  size: 24,
                ),
              ),
              const SizedBox(width: 12),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      event['title'],
                      style: const TextStyle(
                        fontSize: 18,
                        fontWeight: FontWeight.w700,
                        color: AppTheme.textPrimary,
                      ),
                    ),
                    const SizedBox(height: 4),
                    Row(
                      children: [
                        const Icon(
                          LucideIcons.clock,
                          size: 14,
                          color: AppTheme.textSecondary,
                        ),
                        const SizedBox(width: 4),
                        Text(
                          DateFormat.Hm().format(event['arrivalTime']),
                          style: const TextStyle(
                            fontSize: 14,
                            color: AppTheme.textSecondary,
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ],
          ),

          const SizedBox(height: 16),

          // Countdown
          Container(
            padding: const EdgeInsets.all(16),
            decoration: BoxDecoration(
              color: AppTheme.primaryColor.withOpacity(0.1),
              borderRadius: BorderRadius.circular(12),
            ),
            child: Column(
              children: [
                const Text(
                  'TEMPO PER PARTIRE',
                  style: TextStyle(
                    fontSize: 11,
                    fontWeight: FontWeight.w700,
                    color: AppTheme.textSecondary,
                    letterSpacing: 1,
                  ),
                ),
                const SizedBox(height: 8),
                Text(
                  _formatCountdown(event['departureTime']),
                  style: const TextStyle(
                    fontSize: 32,
                    fontWeight: FontWeight.w700,
                    color: AppTheme.primaryColor,
                  ),
                ),
              ],
            ),
          ),

          const SizedBox(height: 16),

          // Info rows
          _buildInfoRow(LucideIcons.mapPin, 'Destinazione', event['location']),
          const SizedBox(height: 12),
          _buildInfoRow(LucideIcons.car, 'Durata viaggio', event['travelTime']),
          const SizedBox(height: 12),
          _buildInfoRow(
            LucideIcons.activity,
            'Traffico',
            _getTrafficLabel(event['traffic']),
            valueColor: _getTrafficColor(event['traffic']),
          ),
        ],
      ),
    );
  }

  Widget _buildInfoRow(
    IconData icon,
    String label,
    String value, {
    Color? valueColor,
  }) {
    return Row(
      children: [
        Container(
          width: 32,
          height: 32,
          decoration: BoxDecoration(
            color: AppTheme.backgroundColor,
            borderRadius: BorderRadius.circular(8),
          ),
          child: Icon(icon, size: 16, color: AppTheme.textSecondary),
        ),
        const SizedBox(width: 12),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                label,
                style: const TextStyle(
                  fontSize: 12,
                  color: AppTheme.textSecondary,
                ),
              ),
              const SizedBox(height: 2),
              Text(
                value,
                style: TextStyle(
                  fontSize: 14,
                  fontWeight: FontWeight.w600,
                  color: valueColor ?? AppTheme.textPrimary,
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }

  String _formatCountdown(DateTime departureTime) {
    final now = DateTime.now();
    final difference = departureTime.difference(now);

    if (difference.isNegative) {
      return 'PARTI ORA!';
    }

    final hours = difference.inHours;
    final minutes = difference.inMinutes % 60;
    final seconds = difference.inSeconds % 60;

    if (hours > 0) {
      return '${hours}h ${minutes}m';
    } else {
      return '${minutes}:${seconds.toString().padLeft(2, '0')}';
    }
  }

  String _getTrafficLabel(String traffic) {
    switch (traffic) {
      case 'light':
        return 'Leggero';
      case 'moderate':
        return 'Moderato';
      case 'heavy':
        return 'Intenso';
      default:
        return 'Sconosciuto';
    }
  }

  Color _getTrafficColor(String traffic) {
    switch (traffic) {
      case 'light':
        return const Color(0xFF51CF66);
      case 'moderate':
        return const Color(0xFFFFE66D);
      case 'heavy':
        return const Color(0xFFFF6B6B);
      default:
        return AppTheme.textSecondary;
    }
  }
}
