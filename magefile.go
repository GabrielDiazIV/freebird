//go:build mage

package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	ProjectName = "Freebird"
)

type (
	Run      mg.Namespace
	Test     mg.Namespace
	Build    mg.Namespace
	Util     mg.Namespace
	Generate mg.Namespace
)

func genURL(name string) string {
	return fmt.Sprintf("%s/%ssvc", ProjectName, name)
}

func (Generate) SQL() error {
	return sh.RunV("sqlc", "generate", "-f", "db/sqlc.yaml")
}

func (Generate) GRPC(name string) error {
	return sh.RunV("protoc", "--go_out=.", "--go-grpc_out=.", fmt.Sprintf("./app/api/proto/%s.proto", name))
}

func (Run) Service(name string) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	servicePath := fmt.Sprintf("app/cmd/%s/main.go", name)
	return sh.RunV("go", "run", servicePath)
}

func (Run) DataDB() error {
	if err := sh.RunV("docker", "build", "-t", "freebird/datadb", "-f", "./db/data/Dockerfile", "./db/data/"); err != nil {
		return err
	}

	sh.RunV("docker", "container", "stop", "freebird_data_db")
	sh.RunV("docker", "container", "rm", "freebird_data_db")

	if err := sh.RunV("docker", "run", "-p", "5432:5432", "-d", "--name=freebird_data_db", "freebird/datadb"); err != nil {
		return err
	}

	return sh.RunV("migrate", "-source", "file://db/data/migrations", "-database", "postgres://postgres:asdfasdf@localhost:5432/postgres?sslmode=disable", "up")
}

func buildRoot() error {
	return sh.RunV("docker", "build", "--no-cache", "--file", "Dockerfile.root", "--tag", genURL("root"), ".")
}

func (Build) Service(name string) error {
	mg.Deps(buildRoot)
	url := genURL(name)
	return sh.RunV("docker", "build", "--no-cache", "--tag", url, "--file", fmt.Sprintf("app/cmd/%s/Dockerfile", name), ".")
}

func (Build) All() error {
	mg.Deps(
		buildRoot,
	)
	return nil
}
