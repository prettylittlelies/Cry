# snowflake

```go
import "github.com/prettylittlelies/cry/snowflake"
```

## Index

- [type ID](#type-id)

## Constants

```go
const Epoch = 1420070400000
```

## Types

### type ID

```go
type ID uint64
```

#### func FromTime

```go
func FromTime(t time.Time) ID
```

#### func Parse

```go
func Parse(s string) (ID, error)
```

#### func (ID) IsZero

```go
func (id ID) IsZero() bool
```

#### func (ID) MarshalText

```go
func (id ID) MarshalText() ([]byte, error)
```

#### func (ID) String

```go
func (id ID) String() string
```

#### func (ID) Time

```go
func (id ID) Time() time.Time
```

#### func (*ID) UnmarshalJSON

```go
func (id *ID) UnmarshalJSON(b []byte) error
```

#### func (*ID) UnmarshalText

```go
func (id *ID) UnmarshalText(b []byte) error
```

