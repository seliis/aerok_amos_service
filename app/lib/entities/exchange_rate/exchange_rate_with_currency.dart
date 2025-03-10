final class ExchangeRateWithCurrency {
  const ExchangeRateWithCurrency({
    required this.id,
    required this.code,
    required this.name,
    required this.base,
    required this.date,
    required this.rate,
  });

  final String id;
  final String code;
  final String name;
  final int base;
  final String date;
  final double rate;

  factory ExchangeRateWithCurrency.fromJson(Map<String, dynamic> json) {
    return ExchangeRateWithCurrency(
      id: json["id"] as String,
      code: json["code"] as String,
      name: json["name"] as String,
      base: json["base"] as int,
      date: json["date"] as String,
      rate: json["rate"] as double,
    );
  }
}
