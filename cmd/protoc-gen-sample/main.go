package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
)

func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		// nolint:wrapcheck
		return nil, err
	}

	var req plugin.CodeGeneratorRequest
	if err = proto.Unmarshal(buf, &req); err != nil {
		// nolint:wrapcheck
		return nil, err
	}

	return &req, nil
}

// nolint:unparam
func processReq(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	files := make(map[string]*descriptor.FileDescriptorProto)

	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}

	for _, fname := range req.FileToGenerate {
		f := files[fname]

		for _, s := range f.GetService() {
			for _, m := range s.GetMethod() {
				log.Println(m.GetName())
				log.Println(m.GetOptions())
				log.Printf("%T", m.GetOptions())
			}
		}
	}

	return nil
}

// func processReq(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
// 	files := make(map[string]*descriptor.FileDescriptorProto)
//
// 	for _, f := range req.ProtoFile {
// 		files[f.GetName()] = f
// 	}
//
// 	var resp plugin.CodeGeneratorResponse
//
// 	for _, fname := range req.FileToGenerate {
// 		// f := files[fname]
// 		out := fname + ".dump"
//
// 		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
// 			Name: proto.String(out),
// 			// Content: proto.String(proto.Marshal(f)),
// 		})
// 	}
//
// 	return &resp
// }

func emitResp(resp *plugin.CodeGeneratorResponse) error {
	// buf, err := proto.Marshal(resp)
	_, err := proto.Marshal(resp)
	if err != nil {
		// nolint:wrapcheck
		return err
	}

	// // _, err = os.Stdout.Write(buf)

	// // nolint:wrapcheck
	// return err
	return nil
}

func run() error {
	req, err := parseReq(os.Stdin)
	if err != nil {
		return err
	}

	resp := processReq(req)

	return emitResp(resp)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
