package main

import (
	"sort"
)

type Tags []string

func (tags *Tags) Add(tag string) bool {
	for i, t := range *tags {
		if t == tag {
			return false
		}
		if t > tag {
			tags.insertAt(i, tag)
			return true
		}
	}
	tags.append(tag)
	return true
}

func (tags *Tags) Contains(tag string) bool {
	return tags.indexOf(tag) >= 0
}

func (tags *Tags) Del(tag string) bool {
	if i := tags.indexOf(tag); i >= 0 {
		tags.removeAt(i)
		return true
	}
	return false
}

func (tags *Tags) append(tag string) {
	*tags = append(*tags, tag)
}

func (tags *Tags) indexOf(tag string) int {
	for i, t := range *tags {
		if t == tag {
			return i
		}
	}
	return -1
}

func (tags *Tags) insertAt(index int, tag string) {
	t := *tags
	t = append(t, "")
	copy(t[index+1:], t[index:])
	t[index] = tag
	*tags = t
}

func (tags *Tags) removeAt(index int) string {
	t := *tags

	old := t[index]
	copy(t[index:], t[index+1:])
	t[len(t)-1] = ""
	t = t[:len(t)-1]

	*tags = t
	return old
}

func (tags Tags) Sort() {
	sort.Stable(tags)
}

func (tags Tags) Len() int {
	return len(tags)
}

func (tags Tags) Less(i, j int) bool {
	return tags[i] < tags[j]
}

func (tags Tags) Swap(i, j int) {
	tags[j], tags[i] = tags[i], tags[j]
}
