import 'dart:convert';
import 'package:http/http.dart' as http;

abstract class BaseApiClient {
  final String baseUrl;
  final Duration timeout;
  final Map<String, String> defaultHeaders;
  final bool enableLogging;

  BaseApiClient({
    required this.baseUrl,
    this.timeout = const Duration(seconds: 30),
    Map<String, String>? defaultHeaders,
    this.enableLogging = false,
  }) : defaultHeaders = defaultHeaders ?? {};

  Future<dynamic> get(
    String endpoint, {
    Map<String, String>? headers,
    Map<String, dynamic>? queryParameters,
  }) async {
    return _request(
      'GET',
      endpoint,
      headers: headers,
      queryParameters: queryParameters,
    );
  }

  Future<dynamic> post(
    String endpoint,
    dynamic body, {
    Map<String, String>? headers,
  }) async {
    return _request('POST', endpoint, body: body, headers: headers);
  }

  Future<dynamic> put(
    String endpoint,
    dynamic body, {
    Map<String, String>? headers,
  }) async {
    return _request('PUT', endpoint, body: body, headers: headers);
  }

  Future<dynamic> delete(
    String endpoint, {
    Map<String, String>? headers,
  }) async {
    return _request('DELETE', endpoint, headers: headers);
  }

  // Template method - può essere override dalle sottoclassi
  Future<dynamic> _request(
    String method,
    String endpoint, {
    dynamic body,
    Map<String, String>? headers,
    Map<String, dynamic>? queryParameters,
  }) async {
    final uri = _buildUri(endpoint, queryParameters);
    final mergedHeaders = _buildHeaders(headers);

    if (enableLogging) {
      _log('$method $uri');
      if (body != null) _log('Body: $body');
    }

    http.Response response;

    try {
      response = await _executeRequest(
        method,
        uri,
        mergedHeaders,
        body,
      ).timeout(timeout);

      return handleResponse(response);
    } catch (e) {
      if (enableLogging) _log('Error: $e');
      return handleError(e);
    }
  }

  Uri _buildUri(String endpoint, Map<String, dynamic>? queryParameters) {
    final url = '$baseUrl$endpoint';

    if (queryParameters != null && queryParameters.isNotEmpty) {
      return Uri.parse(url).replace(
        queryParameters: queryParameters.map(
          (key, value) => MapEntry(key, value.toString()),
        ),
      );
    }

    return Uri.parse(url);
  }

  Map<String, String> _buildHeaders(Map<String, String>? headers) {
    return {...defaultHeaders, ...?headers};
  }

  Future<http.Response> _executeRequest(
    String method,
    Uri uri,
    Map<String, String> headers,
    dynamic body,
  ) {
    switch (method) {
      case 'GET':
        return http.get(uri, headers: headers);
      case 'POST':
        return http.post(
          uri,
          headers: headers,
          body: body != null ? jsonEncode(body) : null,
        );
      case 'PUT':
        return http.put(
          uri,
          headers: headers,
          body: body != null ? jsonEncode(body) : null,
        );
      case 'DELETE':
        return http.delete(uri, headers: headers);
      default:
        throw UnsupportedError('HTTP method $method not supported');
    }
  }

  // Template methods - possono essere override
  dynamic handleResponse(http.Response response) {
    if (enableLogging) {
      _log('Status: ${response.statusCode}');
      _log('Response: ${response.body}');
    }

    if (response.statusCode >= 200 && response.statusCode < 300) {
      if (response.body.isEmpty) return null;
      return jsonDecode(response.body);
    }

    throw ApiException(statusCode: response.statusCode, message: response.body);
  }

  dynamic handleError(Object error) {
    throw error;
  }

  void _log(String message) {
    if (enableLogging) {
      print('[${runtimeType.toString()}] $message');
    }
  }
}

class ApiException implements Exception {
  final int statusCode;
  final String message;

  ApiException({required this.statusCode, required this.message});

  @override
  String toString() => 'ApiException($statusCode): $message';
}
