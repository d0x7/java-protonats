package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"path"
	"strings"
	"unicode"
)

func ToJavaImport(message *protogen.Message) {
	protoFile := getProtoFile(message)
	pkg, class, subClass := getJavaImport(protoFile, message.Desc)
	fmt.Printf("Java Import: %s.%s.%s\n", pkg, class, subClass)
	JavaImport{Package: JavaImportPackage(pkg), Class: class, SubClass: subClass}

}

func getJavaImport(file *protogen.File, desc protoreflect.MessageDescriptor) (pkg, class, subClass string) {
	pkg = file.Proto.GetOptions().GetJavaPackage()
	if pkg == "" {
		if protoPkg := file.Proto.GetPackage(); protoPkg != "" {
			pkg = protoPkg
			//fmt.Printf("Normal pkg != empty - setting javaPackage of %s to %s\n", file.Desc.FullName(), javaPackage)
		} else {
			pkg = "proto"
			//fmt.Printf("Last resort - setting javaPackage of %s to proto\n", file.Desc.FullName())
		}
	} else {
		//fmt.Printf("Java Package set - setting javaPackage of %s to %s\n", file.Desc.FullName(), javaPackage)
	}

	javaClassname := file.GeneratedFilenamePrefix
	if file.Proto.GetOptions().GetJavaOuterClassname() != "" {
		fmt.Printf("Outer != nil - Changing javaClassname of %s from %s to %s\n", file.Desc.FullName(), javaClassname, file.Proto.GetOptions().GetJavaOuterClassname())
		javaClassname = file.Proto.GetOptions().GetJavaOuterClassname()
	} else {
		fmt.Printf("For %s no outer class name existing", file.Desc.FullName())
	}

	// Honor java_multiple_files option
	//if file.Proto.GetOptions().GetJavaMultipleFiles() {
	//	fmt.Printf("Multiple Files - Changing javaClassname of %s from %s to %s\n", file.Desc.FullName(), javaClassname, path.Base(file.GeneratedFilenamePrefix))
	//	javaClassname = path.Base(file.GeneratedFilenamePrefix)
	//}

	//desc.

	//part := lastPart(desc.FullName().Name())

	if file.Proto.GetOptions().GetJavaMultipleFiles() {
		class = string(desc.Name())
	} else {
		class = toPascalCase(path.Base(javaClassname))
		subClass = string(desc.Name())
	}

	fmt.Printf("File")

	fmt.Printf("For file %s(%s), returning javaPackage: %s, javaClassname: %s\n", file.Desc.Name(), desc.Name(), javaPackage, javaClassname)
	return
}

func lastPart(s string) string {
	idx := strings.LastIndex(s, ".")
	if idx == -1 {
		return s // Return the full string if separator is not found
	}
	return s[idx+1:]
}

func toPascalCase(s string) string {
	separators := []rune{'_', '-'}
	words := []string{}
	start := 0

	for i, r := range s {
		if containsRune(separators, r) {
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

func containsRune(slice []rune, r rune) bool {
	for _, v := range slice {
		if v == r {
			return true
		}
	}
	return false
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
