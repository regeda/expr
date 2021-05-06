package memory

// Memory types supported by the virtual machine.
const (
	TypeNil Type = iota
	TypeBytes
	TypeInt64
	TypeBool
	TypeVector
)

var typeNames = [...]string{
	"nil",
	"bytes",
	"int64",
	"bool",
	"vector",
}

// Type represents a data type.
type Type uint8

// String returns a name of the type.
func (t Type) String() string {
	if int(t) < len(typeNames) {
		return typeNames[t]
	}
	return "unknown"
}
