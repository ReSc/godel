package model

func IsValidTag(tag string) bool {
	if len(tag) < 1 {
		return false
	}
	if !isStartNameByte(tag[0]) {
		return false
	}
	tag = tag[1:]
	for i := range tag {
		if !isTagByte(tag[i]) {
			return false
		}
	}
	return true
}

func IsValidName(name string) bool {
	if len(name) < 1 {
		return false
	}
	if !isStartNameByte(name[0]) {
		return false
	}
	name = name[1:]
	for i := range name {
		if !isNameByte(name[i]) {
			return false
		}
	}
	return true
}

func isStartNameByte(c byte) bool {
	return 'A' <= c && c <= 'Z' ||
		'a' <= c && c <= 'z'
}

func isNameByte(c byte) bool {
	return 'A' <= c && c <= 'Z' ||
		'a' <= c && c <= 'z' ||
		'0' <= c && c <= '9' ||
		c == '_'
}

func isTagByte(c byte) bool {
	return 'A' <= c && c <= 'Z' ||
		'a' <= c && c <= 'z' ||
		'0' <= c && c <= '9' ||
		c == '_' ||
		c == ':' || c == '.' || c == '-'
}
