final class ExchangeRate {
  const ExchangeRate({
    required this.currencyCode,
    required this.currencyName,
    required this.date,
    required this.directRate,
  });

  final String currencyCode;
  final String currencyName;
  final String date;
  final double directRate;

  factory ExchangeRate.fromMap(Map<String, dynamic> map) {
    return ExchangeRate(
      currencyCode: map["currency_code"] as String,
      currencyName: map["currency_name"] as String,
      date: map["date"] as String,
      directRate: map["direct_rate"] as double,
    );
  }
}
