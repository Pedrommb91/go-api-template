package tests

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var containerDB *sql.DB
var emptyDB *sql.DB

func NewPostgresTestContainerWithInitScript() *sql.DB {
	if containerDB != nil {
		return containerDB
	}
	ctx := context.Background()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:13-alpine"),
		postgres.WithInitScripts(filepath.Join(path+"../../../../pkg/database/tests", "init-db.sql")),
		postgres.WithDatabase("db"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Connect to the database
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("could not get mapped port: %v", err)
	}

	connStr := fmt.Sprintf("postgres://user:password@localhost:%s/db?sslmode=disable", port.Port())
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	return DB
}

func NewPostgresTestContainerEmpty() *sql.DB {
	if emptyDB != nil {
		return emptyDB
	}
	ctx := context.Background()

	emptyDB, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:13-alpine"),
		postgres.WithDatabase("db"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Connect to the database
	port, err := emptyDB.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("could not get mapped port: %v", err)
	}

	connStr := fmt.Sprintf("postgres://user:password@localhost:%s/db?sslmode=disable", port.Port())
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	return DB
}
