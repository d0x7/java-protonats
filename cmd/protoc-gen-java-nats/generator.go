package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"log"
)

var (
	emptyPB = JavaImport{Package: "com.google.protobuf", Class: "Empty"}
)

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	if len(file.Services) == 0 {
		return
	}

	generatedFile := new(JavaGeneratedFile)
	for _, service := range file.Services {
		generateService(generatedFile, service)
	}
	generatedFile.Finalize(gen, file)
}

func generateService(g *JavaGeneratedFile, service *protogen.Service) {
	generateServer(g, service)
	generateClient(g, service)
}

func generateServer(g *JavaGeneratedFile, service *protogen.Service) {

}

func generateClient(g *JavaGeneratedFile, service *protogen.Service) {
	g.P()
	tab := "    "
	g.P(tab, "public interface I", service.Desc.Name(), " {")
	tab = "        "

	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			// TODO: Skipping currently unsupported streaming methods for now
			g.P("// ", method.Desc.Name(), " is a streaming method and is currently not supported")
			continue
		}

		inputImport := ToJavaImport(method.Input)
		outputImport := ToJavaImport(method.Output)

		log.Printf("Method %s input: %s\n", method.Desc.FullName().Name(), inputImport.FQDN())
		log.Printf("Method %s output: %s\n", method.Desc.FullName().Name(), outputImport.FQDN())

		g.P(tab, outputImport, " ", method.Desc.Name(), "(", inputImport, " req) throws Exception;")
	}
	g.P("}")
}
