# Development Server

This dev server is for third party grpc server to integrate for working with `geliver`.


# Overview

Basically it's just using reflection loop through method to *GET* all request and response to `geliver` and *POST* request from `geliver` to your server and call using reflection also.


# Example

It's fully compatible with pure go-grpc server, other might work but not tested yet.

```go
package main

import (
	gdev "github.com/oscrud/go-geliver-devserver"
)

func main() {
    gdev.Start(portNo, grpcServer, serverHandler, &gdev.Options{
      // Custom Marshaler to make `geliver` to show request with empty values.
			RequestMarshaler: func(name string, rType reflect.Type) []byte {
				rValue := reflect.New(rType.Elem())
				message := rValue.Interface().(proto.Message)
				marshaler := jsonpb.Marshaler{EmitDefaults: true}
				str, _ := marshaler.MarshalToString(message)
				return []byte(str)
			},
		})
}
```

![image](https://user-images.githubusercontent.com/15674107/117440375-8ce0cd80-af66-11eb-9c1f-b203b12b173c.png)
