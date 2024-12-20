package trace

var (
	Flags = FlagNone
)

type flag uint8

const (
	FlagNone = flag(1 << iota >> 1)
	FlagHookCode
	FlagOnCodeHookDisassemble
	FlagHookMemValid
	FlagOnValidMemHookRead
	FlagHookMemInvalid
	FlagOnInvalidMemHookRead
)
