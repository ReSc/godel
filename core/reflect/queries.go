package reflect

import (
	"github.com/ReSc/fmt"
	"path/filepath"
	"strings"
)

func (m *Model) GetPackage(name string) *Package {
	if p, ok := m.Packages[name]; ok {
		return p
	}
	fmt.Panic("Model %v does not contain package %v", m.Name, name)
	return nil
}

func (m *Model) GetBasePath() string {
	return filepath.Join(m.Path, m.Name)
}

func (p *Package) GetBasePath() string {
	return p.Model.GetBasePath()
}
func (p *Package) GetPackagePath() string {
	return filepath.Join(p.Model.GetBasePath(), p.Name)
}

func (p *Package) GetType(name string) *Type {
	name = strings.TrimPrefix(name, "*")
	// local types
	if t, ok := p.Types[name]; ok {
		return t
	}

	// remote types
	if parts := strings.Split(name, "."); len(parts) == 2 {
		if pkg, ok := p.Model.Packages[parts[0]]; ok {
			if t, ok := pkg.Types[name]; ok {
				return t
			}
		}
	}

	// global types
	if global, ok := p.Model.Packages["global"]; ok {
		if t, ok := global.Types[name]; ok {
			return t
		}
	}

	fmt.Panic("Type %v not found in package %v", name, p.Name)
	return nil
}

func (p *Package) HasType(name string) bool {
	name = strings.TrimPrefix(name, "*")
	_, ok := p.Types[name]
	return ok
}

func (t *Type) GetField(name string) *Field {
	if f, ok := t.Fields[name]; ok {
		return f
	}
	fmt.Panic("Unknown field %v on type %v", name, t.Name)
	return nil
}
