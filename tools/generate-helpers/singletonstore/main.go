package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	. "github.com/dave/jennifer/jen"
	"github.com/spf13/cobra"
	"github.com/stackrox/stackrox/pkg/utils"
	"github.com/stackrox/stackrox/tools/generate-helpers/common/packagenames"
	"github.com/stackrox/stackrox/tools/generate-helpers/singletonstore/operations"
)

const (
	bucketNameVariable = "bucketName"
)

func newFile() *File {
	f := NewFile("store")
	f.HeaderComment("Code generated by singletonstore generator. DO NOT EDIT.")
	return f
}

func generateInterface(f *File, interfaces ...Code) {
	f.Type().Id("Store").Interface(
		interfaces...,
	)
}

func generateNew(f *File, props *operations.GeneratorProperties) {
	f.Func().Id("New").Params(Id("db").Op("*").Qual(packagenames.BBolt, "DB")).
		Id("Store").Block(
		Return(Op("&").Id("store").Values(Dict{
			Id("underlying"): Qual(packagenames.SingletonStore, "New").Call(
				Id("db"),
				Id(bucketNameVariable),
				Func().Params().Qual(packagenames.GogoProto, "Message").Block(
					Return(New(Qual(props.Pkg, props.Object))),
				),
				Lit(props.HumanName),
			),
		})),
	)
}

func generateStruct(f *File) {
	f.Type().Id("store").Struct(
		Id("underlying").Qual(packagenames.SingletonStore, "SingletonStore"),
	)
}

func generateMethods(f *File, implementations ...Code) {
	for _, method := range implementations {
		f.Add(method)
		f.Line()
	}
}

func generate(props *operations.GeneratorProperties) error {
	f := newFile()
	f.ImportAlias(packagenames.Ops, "ops")
	f.Var().Defs(Id(bucketNameVariable).Op("=").Index().Byte().Parens(Lit(props.BucketName)))

	var implementations []Code
	var interfaces []Code
	getInterface, getMethod := operations.GenerateGet(props)
	interfaces = append(interfaces, getInterface)
	implementations = append(implementations, getMethod)
	if props.AddInsteadOfUpsert {
		addInterface, addMethod := operations.GenerateAdd(props)
		interfaces = append(interfaces, addInterface)
		implementations = append(implementations, addMethod)
	} else {
		upsertInterface, upsertMethod := operations.GenerateUpsert(props)
		interfaces = append(interfaces, upsertInterface)
		implementations = append(implementations, upsertMethod)
	}

	generateInterface(f, interfaces...)
	generateNew(f, props)
	generateStruct(f)
	generateMethods(f, implementations...)

	if err := f.Save("store.go"); err != nil {
		return err
	}
	if props.GenerateMockStore {
		if props.MockgenWrapperExecutablePath == "" {
			return errors.New("mockgen-wrapper path not specified")
		}
		cmd := exec.Command(props.MockgenWrapperExecutablePath)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOFILE=store.go")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("couldn't exec mockgen-wrapper: %w", err)
		}
	}

	return nil
}

func main() {
	c := &cobra.Command{
		Use: "generate a type-safe implementation for a singleton store",
	}

	props := operations.GeneratorProperties{}
	c.Flags().StringVar(&props.Pkg, "package", "github.com/stackrox/stackrox/generated/storage", "the package of the object being stored")

	c.Flags().StringVar(&props.Object, "object", "", "the (Go) name of the object being stored")
	utils.Must(c.MarkFlagRequired("object"))

	c.Flags().StringVar(&props.HumanName, "human-name", "", "the human-friendly name of the object (defaults to Go name)")

	c.Flags().StringVar(&props.BucketName, "bucket", "", "the name of the bucket")
	utils.Must(c.MarkFlagRequired("bucket"))

	c.Flags().BoolVar(&props.AddInsteadOfUpsert, "add-instead-of-upsert", false, "if the flag is set an Add method will be generated and no Upsert method will be generated")

	c.Flags().BoolVar(&props.GenerateMockStore, "generate-mock-store", false, "whether to generate a mock for the store")
	c.Flags().StringVar(&props.MockgenWrapperExecutablePath, "mockgen-executable-path", "", "path to mockgen-wrapper executable")

	c.RunE = func(*cobra.Command, []string) error {
		if props.HumanName == "" {
			props.HumanName = props.Object
		}
		return generate(&props)
	}

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
