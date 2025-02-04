package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"path/filepath"
)

var protoFiles map[protoreflect.FileDescriptor]*protogen.File

func getProtoFile(method *protogen.Message) *protogen.File {
	file, ok := protoFiles[method.Desc.ParentFile()]
	if !ok {
		panic(fmt.Sprintf("File not found for method %v", method.Desc.FullName()))
	}
	return file
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "print the version and exit")
	flag.BoolVar(&showVersion, "v", false, "print the version and exit")
	flag.Parse()
	if showVersion {
		fmt.Printf("%s %s\n", filepath.Base(os.Args[0]), version)
		return
	}

	var (
		flags flag.FlagSet
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL) | uint64(pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)
		protoFiles = make(map[protoreflect.FileDescriptor]*protogen.File)
		for _, file := range gen.Files {
			protoFiles[file.Desc] = file
		}

		for _, f := range gen.Files {
			if f.Generate {
				generateFile(gen, f)
			}
		}
		return nil
	})
}
