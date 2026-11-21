# tasks

```go
import "github.com/prettylittlelies/cry/ext/tasks"
```

## Index

- [type Func](#type-func)
- [type Loop](#type-loop)
- [type TimeOfDay](#type-timeofday)

## Variables

```go
var ErrRunning = errors.New("tasks: loop already running")
```

## Types

### type Func

```go
type Func func(ctx context.Context) error
```

### type Loop

```go
type Loop struct {
	// contains filtered or unexported fields
}
```

#### func At

```go
func At(times []TimeOfDay, fn Func) *Loop
```

#### func Every

```go
func Every(d time.Duration, fn Func) *Loop
```

#### func (*Loop) After

```go
func (l *Loop) After(fn func(ctx context.Context)) *Loop
```

#### func (*Loop) Before

```go
func (l *Loop) Before(fn func(ctx context.Context)) *Loop
```

#### func (*Loop) Cancel

```go
func (l *Loop) Cancel()
```

#### func (*Loop) Count

```go
func (l *Loop) Count(n int) *Loop
```

#### func (*Loop) Iteration

```go
func (l *Loop) Iteration() int
```

#### func (*Loop) Next

```go
func (l *Loop) Next() time.Time
```

#### func (*Loop) OnError

```go
func (l *Loop) OnError(fn func(ctx context.Context, err error)) *Loop
```

#### func (*Loop) Restart

```go
func (l *Loop) Restart(ctx context.Context) error
```

#### func (*Loop) Running

```go
func (l *Loop) Running() bool
```

#### func (*Loop) SetInterval

```go
func (l *Loop) SetInterval(d time.Duration)
```

#### func (*Loop) Start

```go
func (l *Loop) Start(ctx context.Context) error
```

#### func (*Loop) Stop

```go
func (l *Loop) Stop()
```

#### func (*Loop) Wait

```go
func (l *Loop) Wait()
```

### type TimeOfDay

```go
type TimeOfDay struct {
	Hour     int
	Minute   int
	Second   int
	Location *time.Location
}
```

