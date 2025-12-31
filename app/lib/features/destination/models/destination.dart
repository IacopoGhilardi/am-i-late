class Destination {
  final String? id;
  final String? name;
  final String? formattedAddress;
  final double? latitude;
  final double? longitude;
  final String? googleApiId;
  final String transportMode;

  Destination({
    this.id,
    this.name,
    required this.formattedAddress,
    this.latitude,
    this.longitude,
    this.googleApiId,
    required this.transportMode,
  });

  factory Destination.fromJson(Map<String, dynamic> json) {
    return Destination(
      id: json['id'],
      formattedAddress: json['formatted_address'] ?? '',
      latitude: (json['latitude'] as num).toDouble(),
      longitude: (json['longitude'] as num).toDouble(),
      name: json['name'] ?? '',
      googleApiId: json['google_place_id'],
      transportMode: json['transport_mode'],
    );
  }

  Map<String, dynamic> toJson() => {
    'formatted_address': formattedAddress,
    'latitude': latitude,
    'longitude': longitude,
  };
}
