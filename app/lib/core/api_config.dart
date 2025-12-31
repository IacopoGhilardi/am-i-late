import 'package:flutter_dotenv/flutter_dotenv.dart';

class ApiConfig {
  // Google maps client
  static String get googleMapsBaseUrl =>
      dotenv.env['GOOGLE_MAPS_BASE_URL'] ?? '';
  static String get googleMapsApiKey => dotenv.env['GOOGLE_MAPS_API_KEY'] ?? '';
  static String get googleMapsEnableLogging =>
      dotenv.env['GOOGLE_MAPS_API_ENABLE_LOGGING'] ?? '';

  // Backend
  static String get baseUrl =>
      dotenv.env['API_BASE_URL'] ?? 'http://localhost:8080/api';
  static String get backendEnableLogging =>
      dotenv.env['BACKEND_ENABLE_LOGGING'] ?? '';

  // Environment
  static String get appName => dotenv.env['APP_NAME'] ?? 'App';
  static String get environment => dotenv.env['ENVIRONMENT'] ?? 'development';
  static bool get isDevelopment => environment == 'development';
  static bool get isProduction => environment == 'production';
}
