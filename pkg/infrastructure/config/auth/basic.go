package auth

type BasicExtractor struct {
	baseExtractor
}

func (e BasicExtractor) ExtractToken(getter Getter) (string, error) {
	return e.baseExtractor.ExtractToken(getter, "basic ")
}
