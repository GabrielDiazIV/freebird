//go:build mage

package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	ProjectName = "freebird"
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

	time.Sleep(5 * time.Second)

	return sh.RunV("migrate", "-source", "file://db/data/migrations", "-database", "postgres://postgres:asdfasdf@localhost:5432/postgres?sslmode=disable", "up")
}

func (Run) Envoy() error {
	if err := sh.RunV("docker", "build", "-t", "freebird/envoy", "-f", "./Dockerfile.envoy", "."); err != nil {
		return err
	}

	sh.RunV("docker", "container", "stop", "freebird_envoy")
	sh.RunV("docker", "container", "rm", "freebird_envoy")

	return sh.RunV("docker", "run", "-p", "9901:9901", "-p", "8080:8080", "--name=freebird_envoy", "freebird/envoy", "-c", "/etc/envoy/envoy.yaml", "-l", "off", "--component-log-level", "upstream:debug,connection:trace")
}

func buildRoot() error {
	return sh.RunV("docker", "build", "--no-cache", "--file", "Dockerfile.root", "--tag", genURL("root"), ".")
}

func (Build) Service(name string) error {
	url := genURL(name)
	return sh.RunV("docker", "build", "--no-cache", "--tag", url, "--file", fmt.Sprintf("app/cmd/%s/Dockerfile", name), ".")
}

func (Run) Image(name string) error {
	svcName := fmt.Sprintf("%s_%s_svc", ProjectName, name)
	sh.RunV("docker", "container", "stop", svcName)
	sh.RunV("docker", "container", "rm", svcName)

	return sh.RunV("docker", "run", "-p", "9090:9090", "--name", svcName, genURL(name))
}

func buildDataSvc() error {
	return Build{}.Service("data")
}

func (Build) All() error {
	mg.Deps(
		buildRoot,
		// buildDataSvc,
	)
	return nil
}
