import 'package:app/core/api_config.dart';
import 'package:app/core/base_api_client.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

class GoogleApiClient extends BaseApiClient {
  final String apiKey;

  GoogleApiClient({String? apiKey, bool? enableLogging})
    : apiKey = ApiConfig.googleMapsApiKey,
      super(
        baseUrl: ApiConfig.googleMapsBaseUrl,
        timeout: const Duration(seconds: 15),
        defaultHeaders: {'Accept': 'application/json'},
        enableLogging:
            enableLogging ?? ApiConfig.googleMapsEnableLogging == 'true',
      );
}
