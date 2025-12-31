import 'package:app/core/api_config.dart';
import 'package:app/core/base_api_client.dart';

class BackendApiClient extends BaseApiClient {
  BackendApiClient({bool? enableLogging})
    : super(
        baseUrl: ApiConfig.baseUrl,
        timeout: const Duration(seconds: 10),
        defaultHeaders: {'Accept': 'application/json'},
        enableLogging:
            enableLogging ?? ApiConfig.backendEnableLogging == 'true',
      );
}
