package reflect

import (
	"github.com/ReSc/fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func (m *Model) Initialize() {
	for _, p := range m.Packages {
		p.Model = m
		p.Initialize()
	}
}

func (p *Package) Initialize() {
	for _, i := range p.Imports {
		i.Package = p
		i.Initialize()
	}

	for _, t := range p.Types {
		t.Package = p
		t.Initialize()
	}
}

func (i *Import) Initialize() {
	path, name := filepath.Split(i.Path)
	if i.Alias == "" {
		i.Alias = name
	}
	if path == "../" {
		path = filepath.Join(i.Package.GetBasePath(), name)
		i.Path = strings.Replace(path, "\\", "/", -1)
	}
}

func (t *Type) Initialize() {
	for _, f := range t.Fields {
		f.Type = t
	}
	t.InitializeMetaType()
}

var metaTypeRegex = regexp.MustCompile(`^(?P<type>\w+)\(\s*(?P<args>[*]?\w+)?(?:\s*,\s*[Kk]ey\s*:\s*(?P<key>\w+))?\s*\)$`)

func (t *Type) InitializeMetaType() {
	typeAndArgs := metaTypeRegex.FindAllStringSubmatch(t.MetaType, 3)
	if len(typeAndArgs) == 0 {
		fmt.Panic("Invalid metatype for type %v", t.Name)
	}
	metaType := typeAndArgs[0][1]
	typeArg := typeAndArgs[0][2]
	keyName := typeAndArgs[0][3]
	switch metaType {
	case "primitive":
		t.Meta = &MetaType{
			Name: t.Name,
		}
	case "struct":
		t.Meta = &MetaType{
			Name: "struct",
		}
	case "set":
		t.Meta = &MetaType{
			IsContainer: true,
			Name:        "map[" + typeArg + "]bool",
			KeyType:     typeArg,
			KeyName:     keyName,
			ElementType: "bool",
		}
	case "map":
		keyType := t.Package.GetType(typeArg).GetField(keyName).DataType
		t.Meta = &MetaType{
			IsContainer: true,
			Name:        "map[" + keyType + "]" + typeArg,
			KeyType:     keyType,
			KeyName:     keyName,
			ElementType: typeArg,
		}
	case "list":
		t.Meta = &MetaType{
			IsContainer: true,
			Name:        "[]" + typeArg,
			KeyType:     "int",
			ElementType: typeArg,
		}
	case "listmap":
		keyType := t.Package.GetType(typeArg).GetField(keyName).DataType
		t.Meta = &MetaType{
			IsContainer: true,
			Name:        "struct",
			KeyType:     keyType,
			KeyName:     keyName,
			ElementType: typeArg,
		}
	case "enum":
		t.Meta = &MetaType{
			Name:        typeArg,
			KeyType:     "string",
			KeyName:     "Name",
			ElementType: typeArg,
		}
	default:
		t.Meta = &MetaType{
			Name: "struct",
		}
	}
	t.Meta.Type = metaType
	t.Meta.ElementName = "item"
	t.Meta.ElementTypeName = strings.TrimPrefix(t.Meta.ElementType, "*")
}
