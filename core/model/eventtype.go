package model

type (
	EventType int
)

const (
	Signal EventType = iota
	ChildSet
	ChildDel
	TagSet
	TagDel
	AttrSet
	AttrDel
)

func (c EventType) Value() int {
	return int(c)
}

func (c EventType) String() string {
	switch c {
	case Signal:
		return "Signal"
	case ChildSet:
		return "ChildSet"
	case ChildDel:
		return "ChildDel"
	case TagSet:
		return "TagSet"
	case TagDel:
		return "TagDel"
	case AttrSet:
		return "AttrSet"
	case AttrDel:
		return "AttrDel"
	}
	return "UNDEFINED"
}
