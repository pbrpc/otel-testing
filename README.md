# otel-testing

`otel-testing` provides inspectable OpenTelemetry test doubles for metrics and
traces.

## Installation

```bash
go get github.com/pbrpc/otel-testing
```

## Packages

| Package        | Helpers |
| -------------- | ------- |
| `mocks/meter`  | A `metric.Meter` with inspectable integer counters and configurable instrument-construction errors |
| `mocks/tracer` | An in-memory trace provider with access to recorded spans and events |

## Metrics

`meter.New` creates a meter that records values added to its integer counters:

```go
metrics := meter.New()
counter, err := metrics.Int64Counter("requests")
if err != nil {
	t.Fatal(err)
}

counter.Add(t.Context(), 2)
if got := metrics.GetCounter("requests").Value(); got != 2 {
	t.Fatalf("requests = %d, want 2", got)
}
```

The meter can return configured errors from integer counters, float and integer
histograms, and integer up-down counters. Each error has matching `Set` and
`Clear` methods.

## Traces

`tracer.New` creates an in-memory provider and a test context containing an
active span:

```go
traces, ctx := tracer.New(t)
defer traces.Shutdown(t)

trace.SpanFromContext(ctx).AddEvent("connected")
events := traces.EndSpanAndGetEvents(t)
```

`Tracer` creates additional tracers from the same provider, and `GetSpans`
returns the spans recorded by that provider.
