import 'package:shared_preferences/shared_preferences.dart';

class SettingsService {
  static final SettingsService _instance = SettingsService._internal();
  factory SettingsService() => _instance;
  SettingsService._internal();

  SharedPreferences? _prefs;

  Future<void> init() async {
    _prefs = await SharedPreferences.getInstance();
  }

  Map<String, dynamic> getAll() {
    return _prefs?.getKeys().fold<Map<String, dynamic>>({}, (map, key) {
          map[key] = _prefs?.get(key);
          return map;
        }) ??
        {};
  }

  T? get<T>(String key, {T? defaultValue}) {
    if (T == bool) return (_prefs?.getBool(key) ?? defaultValue) as T?;
    if (T == String) return (_prefs?.getString(key) ?? defaultValue) as T?;
    if (T == int) return (_prefs?.getInt(key) ?? defaultValue) as T?;
    if (T == double) return (_prefs?.getDouble(key) ?? defaultValue) as T?;
    throw Exception('Tipo non supportato');
  }

  Future<void> set<T>(String key, T value) async {
    if (value is bool)
      await _prefs?.setBool(key, value);
    else if (value is String)
      await _prefs?.setString(key, value);
    else if (value is int)
      await _prefs?.setInt(key, value);
    else if (value is double)
      await _prefs?.setDouble(key, value);
    else
      throw Exception('Tipo non supportato');
  }
}
