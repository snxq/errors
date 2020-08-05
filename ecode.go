package errors

// Ecode for define err code.
type Ecode int64

// Int32 returns int32.
func (e Ecode) Int32() int32 { return int32(e) }

// Int64 returns
func (e Ecode) Int64() int64 { return int64(e) }
