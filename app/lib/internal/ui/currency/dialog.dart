part of "screen.dart";

final class _Dialog extends StatefulWidget {
  const _Dialog({
    required this.exchangeRate,
    required this.width,
    required this.height,
  });

  final entities.ExchangeRate exchangeRate;
  final double width;
  final double height;

  @override
  State<_Dialog> createState() => _DialogState();
}

final class _DialogState extends State<_Dialog> {
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
        width: widget.width,
        height: widget.height,
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
