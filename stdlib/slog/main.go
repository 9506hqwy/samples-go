package main

import (
	"context"
	"log"
	"log/slog"
	"os"
)

func main() {
	// OUTPUT:
	// 2026/01/07 10:56:43 INFO Hello, World! key1=value1
	slog.Info("Hello, World!", "key1", "value1")

	// OUTPUT:
	// {"time":"2026-01-05T10:57:10.083392094Z","level":"INFO","msg":"Hello, World!","key1":"value1"}
	logger1 := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger1.Info("Hello, World!", "key1", "value1")

	// OUTPUT:
	// {"time":"2026-01-05T10:58:51.316979948Z","level":"INFO","msg":"Hello, World!","key2":"value2","key1":"value1"}
	logger2 := logger1.With("key2", "value2")
	logger2.Info("Hello, World!", "key1", "value1")

	// OUTPUT:
	// {"time":"2026-01-05T11:31:38.101505853Z","level":"INFO","msg":"Hello, World!","key2":"value2","group":{"key3":"value3"}}
	logger3 := logger2.WithGroup("group")
	logger3.Info("Hello, World!", "key3", "value3")

	// OUTPUT:
	// 2026/01/07 11:30:13 INFO Hello, World! key1=value1 key4=value4
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "key4", "value4")
	logger4 := slog.New(&ContextLogHandler{Handler: slog.Default().Handler()})
	logger4.InfoContext(ctx, "Hello, World!", "key1", "value1")

	// OUTPUT:
	// 2026/01/07 11:21:07 ERROR 1
	// 2026/01/07 11:21:07 WARN 1
	// 2026/01/07 11:21:07 INFO 1
	// 2026/01/07 11:21:07 ERROR 2
	// 2026/01/07 11:21:07 WARN 2
	slog.Error("1")
	slog.Warn("1")
	slog.Info("1")
	slog.Debug("1")
	slog.SetLogLoggerLevel(slog.LevelWarn)
	slog.Error("2")
	slog.Warn("2")
	slog.Info("2")
	slog.Debug("2")
	slog.SetLogLoggerLevel(slog.LevelInfo)

	// OUTPUT:
	// 2026/01/07 11:50:11 INFO Hello, World! value1=Value{}
	// 2026/01/07 11:50:11 INFO Hello, World! value2={}
	slog.Info("Hello, World!", "value1", Value1{})
	slog.Info("Hello, World!", "value2", Value2{})

	// OUTPUT:
	// 2026/01/07 11:57:02.250911 main.go:63: INFO Hello, World!
	log.SetFlags(log.Flags() | log.Lmicroseconds | log.LUTC | log.Lshortfile)
	slog.Info("Hello, World!")

	// OUTPUT:
	// 2026/01/07 11:13:45.992398 main.go:66: INFO Hello, World! key1="[1 2 3]"
	list := []int{1, 2, 3}
	slog.Info("Hello, World!", "key1", list)

	// OUTPUT:
	// 2026/01/07 11:14:42.842182 main.go:71: INFO Hello, World! key1="map[a:1 b:2]"
	dict := map[string]int{"a": 1, "b": 2}
	slog.Info("Hello, World!", "key1", dict)
}

type ContextLogHandler struct {
	slog.Handler
}

func (h *ContextLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if v := ctx.Value("key4"); v != nil {
		r.AddAttrs(slog.Attr{Key: "key4", Value: slog.AnyValue(v)})
	}

	return h.Handler.Handle(ctx, r)
}

type Value1 struct {
}

func (v Value1) LogValue() slog.Value {
	return slog.StringValue("Value{}")
}

type Value2 struct {
}
