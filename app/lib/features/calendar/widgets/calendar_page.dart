import 'package:app/features/calendar/widgets/calendar_add.dart';
import 'package:app/features/calendar/widgets/calendar_week.dart';
import 'package:app/shared/theme/app_theme.dart';
import 'package:app/shared/widgets/page_title.dart';
import 'package:flutter/material.dart';

class CalendarPage extends StatefulWidget {
  const CalendarPage({super.key});

  @override
  State<CalendarPage> createState() => _CalendarPageState();
}

class _CalendarPageState extends State<CalendarPage> {
  DateTime _selectedDate = DateTime.now();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppTheme.backgroundColor,
      body: Column(
        children: [
          const SafeArea(
            child: PageTitle(
              title: 'Appuntamenti',
              subtitle: 'Gestisci i tuoi impegni',
            ),
          ),

          CalendarWeek(
            initialDate: _selectedDate,
            onDateSelected: (date) {
              setState(() => _selectedDate = date);
            },
          ),

          Expanded(child: _buildEmptyState()),
        ],
      ),

      floatingActionButton: FloatingActionButton(
        onPressed: () {
          showModalBottomSheet(
            context: context,
            isScrollControlled: true,
            backgroundColor: Colors.transparent,
            builder: (context) => CalendarAdd(selectedDate: _selectedDate),
          );
        },
        backgroundColor: AppTheme.primaryColor,
        child: const Icon(Icons.add, color: Colors.white),
      ),
    );
  }

  Widget _buildEmptyState() {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(40),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: const [
            Icon(Icons.event_note, size: 80, color: AppTheme.textSecondary),
            SizedBox(height: 20),
            Text(
              'Nessun appuntamento',
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.w700,
                color: AppTheme.textPrimary,
              ),
            ),
            SizedBox(height: 8),
            Text(
              'Tocca + per aggiungerne uno',
              style: TextStyle(fontSize: 14, color: AppTheme.textSecondary),
            ),
          ],
        ),
      ),
    );
  }
}
