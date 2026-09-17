//revive:disable:package-comments
package meter

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// Mock wraps noop.Meter for testing with value tracking and error injection
type Mock struct {
	noop.Meter
	mu sync.RWMutex

	// Counter tracking
	counters map[string]*Int64Counter

	// Error injection
	int64CounterErr       error
	float64HistogramErr   error
	int64HistogramErr     error
	int64UpDownCounterErr error
}

// New creates a new mock meter
func New() *Mock {
	return &Mock{
		counters: make(map[string]*Int64Counter),
	}
}

// Int64Counter creates a mock Int64Counter that tracks values
func (m *Mock) Int64Counter(name string, options ...metric.Int64CounterOption) (metric.Int64Counter, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check for error injection
	if m.int64CounterErr != nil {
		return nil, m.int64CounterErr
	}

	// Return existing counter if already created
	if counter, exists := m.counters[name]; exists {
		return counter, nil
	}

	// Create new counter wrapping noop counter
	noopCounter, _ := m.Meter.Int64Counter(name, options...)
	counter := &Int64Counter{
		Int64Counter: noopCounter,
		name:         name,
	}
	m.counters[name] = counter
	return counter, nil
}

// GetCounter returns the mock counter for inspection
func (m *Mock) GetCounter(name string) *Int64Counter {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.counters[name]
}

// SetInt64CounterError sets an error to be returned by Int64Counter
func (m *Mock) SetInt64CounterError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.int64CounterErr = err
}

// ClearInt64CounterError clears the error injection
func (m *Mock) ClearInt64CounterError() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.int64CounterErr = nil
}

// Int64Counter wraps noop.Int64Counter and tracks values
type Int64Counter struct {
	metric.Int64Counter
	mu    sync.RWMutex
	name  string
	value int64
}

// Add increments the counter and tracks the value
func (c *Int64Counter) Add(ctx context.Context, incr int64, options ...metric.AddOption) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += incr
	c.Int64Counter.Add(ctx, incr, options...)
}

// Value returns the current counter value for inspection
func (c *Int64Counter) Value() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// Float64Histogram creates a noop Float64Histogram with error injection support
func (m *Mock) Float64Histogram(name string, options ...metric.Float64HistogramOption) (metric.Float64Histogram, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check for error injection
	if m.float64HistogramErr != nil {
		return nil, m.float64HistogramErr
	}

	// Return noop histogram
	return m.Meter.Float64Histogram(name, options...)
}

// SetFloat64HistogramError sets an error to be returned by Float64Histogram
func (m *Mock) SetFloat64HistogramError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.float64HistogramErr = err
}

// ClearFloat64HistogramError clears the error injection
func (m *Mock) ClearFloat64HistogramError() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.float64HistogramErr = nil
}

// Int64Histogram creates a noop Int64Histogram with error injection support
func (m *Mock) Int64Histogram(name string, options ...metric.Int64HistogramOption) (metric.Int64Histogram, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check for error injection
	if m.int64HistogramErr != nil {
		return nil, m.int64HistogramErr
	}

	// Return noop histogram
	return m.Meter.Int64Histogram(name, options...)
}

// SetInt64HistogramError sets an error to be returned by Int64Histogram
func (m *Mock) SetInt64HistogramError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.int64HistogramErr = err
}

// ClearInt64HistogramError clears the error injection
func (m *Mock) ClearInt64HistogramError() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.int64HistogramErr = nil
}

// Int64UpDownCounter creates a noop Int64UpDownCounter with error injection support
func (m *Mock) Int64UpDownCounter(name string, options ...metric.Int64UpDownCounterOption) (metric.Int64UpDownCounter, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check for error injection
	if m.int64UpDownCounterErr != nil {
		return nil, m.int64UpDownCounterErr
	}

	// Return noop updown counter
	return m.Meter.Int64UpDownCounter(name, options...)
}

// SetInt64UpDownCounterError sets an error to be returned by Int64UpDownCounter
func (m *Mock) SetInt64UpDownCounterError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.int64UpDownCounterErr = err
}

// ClearInt64UpDownCounterError clears the error injection
func (m *Mock) ClearInt64UpDownCounterError() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.int64UpDownCounterErr = nil
}
