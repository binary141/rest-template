package db

import (
	"embed"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Connect() *sqlx.DB {
	host := getEnv("POSTGRES_HOST", "db")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "dev")
	password := getEnv("POSTGRES_PASSWORD", "secret")
	dbname := getEnv("POSTGRES_DB", "app")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	database, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("connected to database")
	return database
}

func RunMigrations(database *sqlx.DB) {
	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatalf("failed to create schema_migrations table: %v", err)
	}

	entries, err := migrations.ReadDir("migrations")
	if err != nil {
		log.Fatalf("failed to read migrations directory: %v", err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	for _, file := range files {
		var count int
		err := database.Get(&count, "SELECT COUNT(*) FROM schema_migrations WHERE version = $1", file)
		if err != nil {
			log.Fatalf("failed to check migration %s: %v", file, err)
		}
		if count > 0 {
			continue
		}

		content, err := migrations.ReadFile("migrations/" + file)
		if err != nil {
			log.Fatalf("failed to read migration %s: %v", file, err)
		}

		tx := database.MustBegin()
		tx.MustExec(string(content))
		tx.MustExec("INSERT INTO schema_migrations (version) VALUES ($1)", file)
		if err := tx.Commit(); err != nil {
			log.Fatalf("failed to apply migration %s: %v", file, err)
		}

		log.Printf("applied migration: %s", file)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
