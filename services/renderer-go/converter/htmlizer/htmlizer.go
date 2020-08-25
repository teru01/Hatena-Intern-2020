package htmlizer

type LineConverter interface {
	convert(src string) (string, error)
}

type WholeConverter LineConverter
