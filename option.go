package human

// DefaultTagName defaults the tag name option
const DefaultTagName = "human"

// DefaultListSymbol defaults the list symbols
const DefaultListSymbol = "*"

// DefaultIndent ,,,
const DefaultIndent = 2

// Option defines the function type of Encoder options
type Option func(*Encoder) error

var defaultOptions = []Option{
	OptionTagName(DefaultTagName),
	OptionListSymbols(DefaultListSymbol),
	OptionIndent(DefaultIndent),
}

// OptionTagName specifies the tag name
func OptionTagName(tagName string) Option {
	return func(e *Encoder) error {
		if tagName == "" {
			return ErrInvalidTagName
		}
		e.tagName = tagName
		return nil
	}
}

// OptionListSymbols specifies the list symbols
func OptionListSymbols(listSymbols ...string) Option {
	return func(e *Encoder) error {
		if len(listSymbols) == 0 {
			return ErrListSymbolsEmpty
		}
		e.listSymbols = listSymbols
		return nil
	}
}

// OptionIndent specifies the used indentation
func OptionIndent(indent uint) Option {
	return func(e *Encoder) error {
		e.indent = indent
		return nil
	}
}
