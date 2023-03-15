# Waterslog

[Watermill](watermill.io)'s log adapter for standard GO [slog](golang.org/x/exp/slog).

## Usage example

    l  := slog.New(slog.HandlerOptions{Level: slog.LevelDebug}.NewJSONHandler(stdout))
    
    logger  := waterslog.New(l)
    logger.Info("hello", watermill.LogFields{"foo": "bar"})
   
   
Output:
> {"time":"2023-03-15T20:07:50.877803+03:00","level":"INFO","msg":"hello","foo":"bar"}

