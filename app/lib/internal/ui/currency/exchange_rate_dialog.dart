part of "screen.dart";

final class _ExchangeRateDialog extends StatefulWidget {
  const _ExchangeRateDialog({
    required this.exchangeRate,
    required this.size,
  });

  final entities.ExchangeRate exchangeRate;
  final Size size;

  @override
  State<_ExchangeRateDialog> createState() => _DialogState();
}

final class _DialogState extends State<_ExchangeRateDialog> {
  bool isCopied = false;

  @override
  Widget build(context) {
    return AlertDialog(
      title: Text(
        widget.exchangeRate.currencyName,
        style: const TextStyle(
          fontWeight: FontWeight.w700,
        ),
      ),
      content: SizedBox(
        width: widget.size.width,
        height: widget.size.height,
        child: ListView(
          children: [
            ListTile(
              contentPadding: EdgeInsets.zero,
              title: const Text(
                "Currency Code",
                style: TextStyle(
                  fontWeight: FontWeight.w700,
                ),
              ),
              subtitle: Text(widget.exchangeRate.currencyCode),
            ),
            ListTile(
              contentPadding: EdgeInsets.zero,
              title: const Text(
                "Date",
                style: TextStyle(
                  fontWeight: FontWeight.w700,
                ),
              ),
              subtitle: Text(widget.exchangeRate.date),
            ),
            ListTile(
              contentPadding: EdgeInsets.zero,
              title: Text(
                "One ${widget.exchangeRate.currencyCode} Equals",
                style: const TextStyle(
                  fontWeight: FontWeight.w700,
                ),
              ),
              subtitle: Text("${widget.exchangeRate.directRate.toString()} KRW"),
              trailing: IconButton(
                icon: isCopied ? const Text("Copied") : const Icon(Icons.copy),
                onPressed: isCopied
                    ? null
                    : () {
                        Clipboard.setData(ClipboardData(text: widget.exchangeRate.directRate.toString()));
                        setState(() {
                          isCopied = true;
                        });
                      },
              ),
            ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: const Text("Close"),
        ),
      ],
    );
  }
}
