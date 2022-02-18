> :warning: **geliver-vscode is now support `.proto` files**: This repository will be deprecated and will no longer updates anymore. https://github.com/Oskang09/geliver-vscode

# Development Server

This dev server is for third party grpc server to integrate for working with `geliver`.


# Overview

Basically it's just using reflection loop through method to *GET* all request and response to `geliver` and *POST* request from `geliver` to your server and call using reflection also.

# Frontend UI

For frontend ui open in browser at `https://oskang09.github.io/geliver/`. Currently browser is available for geliver devserver, vscode extension will natively work with `.proto`.

# Example Server Setup

It's fully compatible with pure go-grpc server, other might work but not tested yet.

```go
package main

import (
	devserver "github.com/oscrud/go-geliver-devserver"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	devserver.Start(portNo, server, handler, &devserver.Options{
		RequestMarshaler: func(name string, rType reflect.Type) []byte {
			rValue := reflect.New(rType.Elem())
			message := rValue.Interface().(proto.Message)
			marshaler := protojson.MarshalOptions{EmitUnpopulated: true}
			str, _ := marshaler.Marshal(message)
			return []byte(str)
		},
	})
}
```

![image](https://user-images.githubusercontent.com/15674107/117440375-8ce0cd80-af66-11eb-9c1f-b203b12b173c.png)

# Known Issues

* Beacuse of using reflection to access directly, will by pass all grpc server configuration etc `UnaryInterceptor`.
