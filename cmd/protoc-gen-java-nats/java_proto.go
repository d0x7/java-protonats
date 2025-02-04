package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"path"
	"strings"
	"unicode"
)

func ToJavaImport(message *protogen.Message) JavaImport {
	protoFile := getProtoFile(message)
	pkg, class, subClass := getJavaImport(protoFile, message.Desc)
	return JavaImport{Package: JavaImportPackage(pkg), Class: class, SubClass: subClass}
}

func getJavaImport(file *protogen.File, desc protoreflect.MessageDescriptor) (pkg, class, subClass string) {
	pkg = file.Proto.GetOptions().GetJavaPackage()
	if pkg == "" {
		if protoPkg := file.Proto.GetPackage(); protoPkg != "" {
			pkg = protoPkg
		} else {
			pkg = "proto"
		}
	}

	javaClassname := file.GeneratedFilenamePrefix
	if file.Proto.GetOptions().GetJavaOuterClassname() != "" {
		javaClassname = file.Proto.GetOptions().GetJavaOuterClassname()
	}

	if file.Proto.GetOptions().GetJavaMultipleFiles() {
		class = string(desc.Name())
	} else {
		class = toPascalCase(path.Base(javaClassname))
		subClass = string(desc.Name())
	}

	return
}

func toPascalCase(s string) string {
	var words []string
	start := 0

	for i, r := range s {
		if r == '_' {
			if start < i {
				words = append(words, capitalize(s[start:i]))
			}
			start = i + 1
		}
	}
	if start < len(s) {
		words = append(words, capitalize(s[start:]))
	}

	return strings.Join(words, "")
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
