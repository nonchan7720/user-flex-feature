package auth

type BearerExtractor struct {
	baseExtractor
}

func (e BearerExtractor) ExtractToken(getter Getter) (string, error) {
	return e.baseExtractor.ExtractToken(getter, "bearer ")
}
