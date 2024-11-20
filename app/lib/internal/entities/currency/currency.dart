final class Currency {
  const Currency({
    required this.code,
    required this.name,
  });

  final String code;
  final String name;

  factory Currency.fromMap(Map<String, dynamic> map) {
    return Currency(
      code: map["code"] as String,
      name: map["name"] as String,
    );
  }
}
