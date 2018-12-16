package generator

import (
	"fmt"
	"strings"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator/minimal"
	"go.larrymyers.com/protoc-gen-twirp_typescript/generator/pbjs"
)

func GetParameters(in *plugin.CodeGeneratorRequest) map[string]string {
	params := make(map[string]string)

	if in.Parameter == nil {
		return params
	}

	pairs := strings.Split(*in.Parameter, ",")

	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		params[kv[0]] = kv[1]
	}

	return params
}

type Generator interface {
	Generate(in *plugin.CodeGeneratorRequest) ([]*plugin.CodeGeneratorResponse_File, error)
}

func NewGenerator(p map[string]string) (Generator, error) {
	version, ok := p["version"]
	if !ok {
		version = "v5"
	}
	prettierPath, ok := p["prettier_path"]
	if !ok {
		prettierPath = "prettier"
	}

	if version != "v5" && version != "v6" {
		return nil, fmt.Errorf("version is %s, must be v5 or v6", version)
	}

	lib, ok := p["library"]
	if ok && lib == "pbjs" {
		return pbjs.NewGenerator(version, prettierPath), nil
	}

	return minimal.NewGenerator(version, p, prettierPath), nil
}
