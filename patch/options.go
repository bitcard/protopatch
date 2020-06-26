package patch

import (
	patch_go "github.com/alta/protopatch/patch/go"
	"github.com/alta/protopatch/patch/go/enum"
	"github.com/alta/protopatch/patch/go/field"
	"github.com/alta/protopatch/patch/go/message"
	"github.com/alta/protopatch/patch/go/oneof"
	"github.com/alta/protopatch/patch/go/value"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// EnumOptions returns patch_go.Options if present, otherwise nil.
func EnumOptions(e *protogen.Enum) *patch_go.Options {
	return Options(e.Desc, enum.E_Options)
}

// ValueOptions returns patch_go.Options if present, otherwise nil.
func ValueOptions(v *protogen.EnumValue) *patch_go.Options {
	return Options(v.Desc, value.E_Options)
}

// MessageOptions returns patch_go.Options if present, otherwise nil.
func MessageOptions(m *protogen.Message) *patch_go.Options {
	return Options(m.Desc, message.E_Options)
}

// FieldOptions returns patch_go.Options if present on e, otherwise nil.
func FieldOptions(f *protogen.Field) *patch_go.Options {
	// First try (go.field.options)
	if opts := Options(f.Desc, field.E_Options); opts != nil {
		return opts
	}
	// Then try (go.field.options)
	if opts := Options(f.Desc, patch_go.E_Options); opts != nil {
		return opts
	}
	return nil
}

// OneofOptions returns patch_go.Options if present on e, otherwise nil.
func OneofOptions(o *protogen.Oneof) *patch_go.Options {
	return Options(o.Desc, oneof.E_Options)
}

// Options returns patch_go.Options if present on desc, otherwise nil.
func Options(desc protoreflect.Descriptor, xt protoreflect.ExtensionType) *patch_go.Options {
	o := desc.Options()
	if proto.HasExtension(o, xt) {
		return proto.GetExtension(o, xt).(*patch_go.Options)
	}
	return nil
}
