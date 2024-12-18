# protoc-gen-java-nats

This is a protoc plugin that generates Java server and client code for NATS microservices.

Prior experience with Protobuf is greatly recommended, especially to understand how the package and imports work.

## Installation

As this plugin is written in Go, you need to have Go installed to be able to install this plugin.
You also already need to have the protoc compiler along the Java protobuf plugin installed on your system.
After that, you can go ahead and install this plugin using the following command:

```shell
go install xiam.li/java-nats/cmd/protoc-gen-java-nats@latest
```

To check if the installation was successful, you can run:

```shell
protoc-gen-java-nats -v
```

## Usage

Upon installation, you should create a protobuf file that contains a service, similar to how gRPC servers work.
An example protobuf file might look like this:

```protobuf
syntax = "proto3";
package your.package;
option go_package = "github.com/user/repo/pb;pb";
option java_multiple_files = true;
option java_package = "me.username.repo.service";

service HelloWorldService {
    rpc HelloWorld(HelloWorldRequest) returns (HelloWorldResponse);
}

message HelloWorldRequest {
    string name = 1;
}

message HelloWorldResponse {
    string message = 1;
}
```

To generate the Java code for this service, run the following command.
This command expects your proto file in a directory called `pb` in your project.

```shell
protoc -I pb --java_out=pb --java-nats_out=pb --go-nats_opt=paths=source_relative pb/hello_world.proto
```

This obviously requires the protoc compiler to be installed on your system
and also having the Java protobuf plugin installed, so that besides the code
regarding NATS can be generated, the messages and everything else can also be generated.

### Streaming

Streaming is not yet supported, but is planned for the future.
It'll probably be implemented along with better timeout handling,
that will come with keepalive messages and therefore also allow streaming.
