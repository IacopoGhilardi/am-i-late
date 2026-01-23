import 'package:app/features/destination/models/destination.dart';
import 'package:app/features/destination/widgets/destination_detail.dart';
import 'package:app/features/destination/widgets/destination_item.dart';
import 'package:app/shared/theme/app_theme.dart';
import 'package:app/shared/widgets/custom_floating_button.dart';
import 'package:app/shared/widgets/page_title.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class DestinationPage extends StatefulWidget {
  const DestinationPage({super.key});

  @override
  State<DestinationPage> createState() => _DestinationPageState();
}

class _DestinationPageState extends State<DestinationPage> {
  final destinations = [
    Destination(
      id: 'ciao',
      name: 'Nome_1',
      formattedAddress: 'via Francesco de sanctis 5',
      transportMode: 'car',
    ),
    Destination(
      id: 'uuid2',
      name: 'Nome_2',
      formattedAddress: 'via Francesco de sanctis 6',
      transportMode: 'car',
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: SingleChildScrollView(
          child: Column(
            children: [
              PageTitle(
                title: 'Le mie destinazioni',
                subtitle: 'Gestisci i tuoi luoghi preferiti',
              ),
              SizedBox(height: 20),
              for (var destination in destinations)
                DestinationItem(destination: destination),
            ],
          ),
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      floatingActionButton: Padding(
        padding: EdgeInsets.symmetric(horizontal: 20),
        child: CustomFloatingButton(
          onTap:
              () => {
                showModalBottomSheet(
                  context: context,
                  isScrollControlled: true,
                  backgroundColor: Colors.transparent,
                  builder:
                      (context) =>
                          const DestinationDetail(), // ← Nessuna destinazione = creazione
                ).then((result) {
                  if (result == true) {
                    setState(() {
                      // TODO: Ricarica lista destinazioni
                    });
                  }
                }),
              },
          widget: Container(
            width: double.infinity,
            padding: const EdgeInsets.all(12),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Icon(LucideIcons.plus, color: Colors.white, size: 24),
                SizedBox(width: 10),
                Text(
                  "Aggiungi destinazione",
                  style: TextStyle(
                    color: AppTheme.backgroundColor,
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
