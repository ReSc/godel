package reflect

import (
	"os"
	"path/filepath"
	"strings"
)

func (m *Model) String() string {
	return "model " + m.Name + "\n"
}

func (p *Package) GetPackageDirectory() string {
	gopath := os.Getenv("GOPATH")
	dir := filepath.Join(gopath, "src", p.GetPackagePath())
	return dir
}

func (p *Package) String() string {
	return "package " + p.Name + "\n"
}

func (p *Type) BaseName() string {
	return strings.TrimPrefix(p.Name, "*")
}

func (f *Field) FieldType() *Type {
	if f.Type.Package.HasType(f.DataType) {
		fieldType := f.Type.Package.GetType(f.DataType)
		return fieldType
	}
	return nil
}

func (f *Field) IsContainer() bool {
	fieldType := f.FieldType()
	return fieldType != nil && fieldType.Meta.IsContainer
}

// visitor pattern implementation
func (this *Model) doAccept(v Visitor) {
	v.VisitModel(this)
	for _, key := range this.Packages.SortedKeys() {
		p := this.Packages[key]
		p.Accept(v)
	}
}

func (this *Package) doAccept(v Visitor) {
	v.VisitPackage(this)
	for _, key := range this.Types.SortedKeys() {
		t := this.Types[key]
		t.Accept(v)
	}
}

func (this *Type) doAccept(v Visitor) {
	v.VisitType(this)

	for _, key := range this.Fields.SortedKeys() {
		f := this.Fields[key]
		f.Accept(v)
	}

	for _, key := range this.Methods.SortedKeys() {
		f := this.Methods[key]
		f.Accept(v)
	}
}

func (this *Field) doAccept(v Visitor) {
	v.VisitField(this)
}

func (this *Method) doAccept(v Visitor) {
	v.VisitMethod(this)
}
