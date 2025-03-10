final class Currency {
  const Currency({required this.code, required this.name, required this.base});

  final String code;
  final String name;
  final int base;

  factory Currency.fromJson(Map<String, dynamic> json) {
    return Currency(
      code: json["code"] as String,
      name: json["name"] as String,
      base: json["base"] as int,
    );
  }
}
