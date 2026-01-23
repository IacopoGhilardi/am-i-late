import 'package:app/features/destination/models/destination.dart';
import 'package:app/features/destination/widgets/destination_detail.dart';
import 'package:app/shared/theme/app_theme.dart';
import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';

class DestinationItem extends StatelessWidget {
  final Destination destination;
  final VoidCallback? onDelete;
  final VoidCallback? onUpdate; // Callback dopo modifica

  const DestinationItem({
    super.key,
    required this.destination,
    this.onDelete,
    this.onUpdate,
  });

  @override
  Widget build(BuildContext context) {
    return Dismissible(
      key: Key(destination.id ?? ''),
      direction: DismissDirection.endToStart,
      background: Container(
        alignment: Alignment.centerRight,
        padding: const EdgeInsets.only(right: 20),
        margin: const EdgeInsets.symmetric(horizontal: 20, vertical: 8),
        decoration: BoxDecoration(
          color: AppTheme.errorColor,
          borderRadius: BorderRadius.circular(20),
        ),
        child: const Icon(LucideIcons.trash2, color: Colors.white, size: 28),
      ),
      confirmDismiss: (direction) => _confirmDelete(context),
      onDismissed: (direction) => onDelete?.call(),
      child: GestureDetector(
        onTap: () => _openEditSheet(context), // ← Apre bottom sheet
        child: Container(
          margin: const EdgeInsets.symmetric(horizontal: 20, vertical: 8),
          padding: const EdgeInsets.all(16),
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
          child: Row(
            children: [
              // Icona categoria
              Container(
                width: 48,
                height: 48,
                decoration: BoxDecoration(
                  gradient: LinearGradient(
                    begin: Alignment.topLeft,
                    end: Alignment.bottomRight,
                    colors: [
                      _getCategoryColor().withOpacity(0.2),
                      _getCategoryColor().withOpacity(0.1),
                    ],
                  ),
                  borderRadius: BorderRadius.circular(12),
                ),
                child: Icon(
                  _getCategoryIcon(),
                  color: _getCategoryColor(),
                  size: 24,
                ),
              ),

              const SizedBox(width: 16),

              // Info destinazione
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      destination.name ?? 'Destinazione senza nome',
                      style: const TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.w700,
                        color: AppTheme.textPrimary,
                      ),
                    ),
                    const SizedBox(height: 4),
                    Row(
                      children: [
                        const Icon(
                          LucideIcons.mapPin,
                          size: 14,
                          color: AppTheme.textSecondary,
                        ),
                        const SizedBox(width: 4),
                        Expanded(
                          child: Text(
                            destination.formattedAddress ??
                                'Indirizzo non disponibile',
                            style: const TextStyle(
                              fontSize: 13,
                              color: AppTheme.textSecondary,
                            ),
                            maxLines: 1,
                            overflow: TextOverflow.ellipsis,
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),

              // Menu button
              GestureDetector(
                onTap: () => _showMenu(context),
                behavior: HitTestBehavior.opaque,
                child: Padding(
                  padding: const EdgeInsets.all(8),
                  child: const Icon(
                    LucideIcons.moreVertical,
                    color: AppTheme.textSecondary,
                    size: 20,
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  // Apri bottom sheet per modifica
  void _openEditSheet(BuildContext context) {
    showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      backgroundColor: Colors.transparent,
      builder: (context) => DestinationDetail(destination: destination),
    ).then((result) {
      // Se ha salvato con successo, chiama callback
      if (result == true) {
        onUpdate?.call();
      }
    });
  }

  void _showMenu(BuildContext context) {
    showModalBottomSheet(
      context: context,
      shape: const RoundedRectangleBorder(
        borderRadius: BorderRadius.vertical(top: Radius.circular(20)),
      ),
      builder:
          (context) => Container(
            padding: const EdgeInsets.all(20),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                Container(
                  width: 40,
                  height: 4,
                  decoration: BoxDecoration(
                    color: AppTheme.textSecondary.withOpacity(0.3),
                    borderRadius: BorderRadius.circular(2),
                  ),
                  margin: const EdgeInsets.only(bottom: 20),
                ),
                ListTile(
                  leading: const Icon(LucideIcons.pencil),
                  title: const Text('Modifica'),
                  onTap: () {
                    Navigator.pop(context);
                    _openEditSheet(context);
                  },
                ),
                ListTile(
                  leading: const Icon(
                    LucideIcons.trash2,
                    color: AppTheme.errorColor,
                  ),
                  title: const Text(
                    'Elimina',
                    style: TextStyle(color: AppTheme.errorColor),
                  ),
                  onTap: () async {
                    Navigator.pop(context);
                    final confirmed = await _confirmDelete(context);
                    if (confirmed == true) {
                      onDelete?.call();
                    }
                  },
                ),
              ],
            ),
          ),
    );
  }

  Future<bool?> _confirmDelete(BuildContext context) {
    return showDialog<bool>(
      context: context,
      builder:
          (context) => AlertDialog(
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(20),
            ),
            title: const Text('Elimina destinazione'),
            content: Text('Vuoi eliminare "${destination.name}"?'),
            actions: [
              TextButton(
                onPressed: () => Navigator.pop(context, false),
                child: const Text('Annulla'),
              ),
              TextButton(
                onPressed: () => Navigator.pop(context, true),
                style: TextButton.styleFrom(
                  foregroundColor: AppTheme.errorColor,
                ),
                child: const Text('Elimina'),
              ),
            ],
          ),
    );
  }

  IconData _getCategoryIcon() {
    final name = destination.name?.toLowerCase() ?? '';

    if (name.contains('ufficio') || name.contains('lavoro')) {
      return LucideIcons.briefcase;
    } else if (name.contains('casa')) {
      return LucideIcons.home;
    } else if (name.contains('palestra')) {
      return LucideIcons.dumbbell;
    } else if (name.contains('supermercato')) {
      return LucideIcons.shoppingCart;
    } else if (name.contains('ristorante') || name.contains('bar')) {
      return LucideIcons.utensils;
    } else if (name.contains('cinema')) {
      return LucideIcons.film;
    } else if (name.contains('stazione') || name.contains('aeroporto')) {
      return LucideIcons.train;
    }

    return LucideIcons.mapPin;
  }

  Color _getCategoryColor() {
    final name = destination.name?.toLowerCase() ?? '';

    if (name.contains('ufficio') || name.contains('lavoro')) {
      return const Color(0xFF8B5CF6);
    } else if (name.contains('casa')) {
      return const Color(0xFF2196F3);
    } else if (name.contains('palestra')) {
      return const Color(0xFFFF9800);
    } else if (name.contains('supermercato')) {
      return const Color(0xFF00C853);
    } else if (name.contains('ristorante')) {
      return const Color(0xFFFF5252);
    }

    return AppTheme.textSecondary;
  }
}
