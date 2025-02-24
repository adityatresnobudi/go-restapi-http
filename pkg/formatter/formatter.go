package formatter

func StringOrNil(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}
