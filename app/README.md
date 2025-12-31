# 📁 Struttura Progetto Flutter

## Schema Generale
```
lib/
├── main.dart
├── app.dart
│
├── core/                          # Infrastruttura condivisa
│   ├── config/                    # Configurazioni (API, routes)
│   ├── constants/                 # Costanti app-wide
│   ├── network/                   # HTTP client, gestione rete
│   └── utils/                     # Helper functions, extensions
│
├── features/                      # Funzionalità dell'app
│   ├── feature_1/
│   │   ├── data/
│   │   │   ├── models/           # Strutture dati (JSON ↔ Dart)
│   │   │   ├── repositories/     # Interfaccia verso i dati
│   │   │   └── services/         # Chiamate API
│   │   └── presentation/
│   │       ├── providers/        # State management
│   │       ├── screens/          # Pagine complete
│   │       └── widgets/          # Componenti UI specifici
│   │
│   └── feature_2/
│       └── ...
│
└── shared/                        # Componenti riutilizzabili
    ├── widgets/                   # Widget custom globali
    └── themes/                    # Temi, stili
```

## 🎯 Regole Base

### 1. Una Feature = Una Cartella
```
features/
  ├── auth/           → Login, registrazione, logout
  ├── destinations/   → CRUD destinazioni
  ├── trips/          → CRUD viaggi
  └── profile/        → Profilo utente
```

### 2. Dentro ogni Feature
```
feature_name/
  ├── data/              → Come prendo/salvo i dati
  │   ├── models/        → Classe Dart ↔ JSON
  │   ├── services/      → Chiamate HTTP
  │   └── repositories/  → Orchestrazione servizi
  │
  └── presentation/      → Come mostro i dati
      ├── providers/     → Logica e stato
      ├── screens/       → Schermate complete
      └── widgets/       → Pezzi di UI riutilizzabili
```

### 3. Flusso Dati
```
Screen → Provider → Repository → Service → API Backend
  ↓         ↓           ↓            ↓
 UI    State Mgmt   Logica      HTTP Call
```

## 📝 Esempio Pratico

### Voglio aggiungere "Notifiche"
```
1. Creo cartella:
   features/notifications/

2. Creo il model:
   data/models/notification_model.dart
   
3. Creo il service:
   data/services/notification_service.dart
   
4. Creo il repository:
   data/repositories/notification_repository.dart
   
5. Creo il provider:
   presentation/providers/notification_provider.dart
   
6. Creo la screen:
   presentation/screens/notifications_screen.dart
```

## 🔧 Core vs Shared

| `core/`                     | `shared/`                    |
|-----------------------------|------------------------------|
| API client                  | Bottoni custom               |
| Configurazioni              | TextField custom             |
| Gestione errori             | Card riutilizzabili          |
| Utility functions           | Temi e colori                |
| NON dipende da feature      | Usato da più feature         |

## ✅ Checklist Nuova Feature

- [ ] Creo cartella `features/nome_feature/`
- [ ] Aggiungo `data/models/` per i dati
- [ ] Aggiungo `data/services/` per le API
- [ ] Aggiungo `data/repositories/` per orchestrare
- [ ] Aggiungo `presentation/providers/` per lo stato
- [ ] Aggiungo `presentation/screens/` per le pagine
- [ ] Registro provider in `app.dart`
- [ ] Aggiungo route in `core/config/routes.dart`

## 🎨 Naming Convention
```dart
// File
destination_list_screen.dart        // snake_case + suffisso
user_model.dart
auth_provider.dart

// Classi
DestinationListScreen              // PascalCase + suffisso
UserModel
AuthProvider

// Variabili
isLoading                          // camelCase
_privateVariable                   // underscore per private
```

## 🚫 Cosa NON Fare

❌ Mettere logica business nelle screen  
❌ Chiamare API direttamente dalle screen  
❌ Creare dipendenze nelle classi (usa Dependency Injection)  
❌ Duplicare codice (spostalo in `core/` o `shared/`)  
❌ Fare feature troppo grandi (dividi in sotto-feature)

## 💡 Quando Modificare Cosa

| Se cambia...              | Modifico...           |
|---------------------------|-----------------------|
| Struttura dati backend    | `models/`             |
| Endpoint API              | `services/`           |
| Logica business           | `repositories/`       |
| Come mostro i dati        | `providers/`          |
| Layout pagina             | `screens/`            |
| Componente riutilizzabile | `widgets/` o `shared/`|

## 🎓 Ricorda

1. **Una feature = indipendente dalle altre**
2. **Presentation non sa nulla di API**
3. **Data non sa nulla di UI**
4. **Shared = usato ovunque**
5. **Core = fondamenta dell'app**

---

**Domanda rapida:** Dove metto X?
- È infrastruttura? → `core/`
- È UI riutilizzabile? → `shared/`
- È specifico di una funzionalità? → `features/nome_feature/`