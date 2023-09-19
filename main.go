package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arfan21/mkpmobile/middleware"
	"github.com/arfan21/mkpmobile/migrations"
	"github.com/arfan21/mkpmobile/terminal"
	"github.com/arfan21/mkpmobile/users"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	app := cli.NewApp()
	app.Name = "mkp"
	app.Description = "mkp-mobile"
	app.Commands = []*cli.Command{
		{
			Name: "serve",
			Action: func(c *cli.Context) error {

				db, err := pgxpool.New(context.Background(), dsn)
				if err != nil {
					panic(err)
				}

				defer db.Close()

				if err := db.Ping(context.Background()); err != nil {
					panic(err)
				}

				userRepo := users.NewRepository(db)
				userService := users.NewService(userRepo)
				userController := users.NewController(userService)

				terminalRepo := terminal.NewRepository(db)
				terminalService := terminal.NewService(terminalRepo)
				terminalController := terminal.NewController(terminalService)

				r := mux.NewRouter()

				r.HandleFunc("/users", userController.RegisterUser).Methods("POST")
				r.HandleFunc("/login", userController.Login).Methods("POST")
				r.Handle("/terminals", middleware.Auth(http.HandlerFunc(terminalController.RegisterTerminal))).Methods("POST")

				port := "8080"
				if os.Getenv("PORT") != "" {
					port = os.Getenv("PORT")
				}

				log.Println("Server running on http://localhost:" + port)
				http.ListenAndServe(":"+port, r)

				return nil
			},
		},
		{
			Name: "migrate",
			Action: func(c *cli.Context) error {
				conn, err := sql.Open("pgx", dsn)
				if err != nil {
					return fmt.Errorf("failed to connect to database: %w", err)
				}

				defer conn.Close()

				migration, err := migrations.New(conn)
				if err != nil {
					return fmt.Errorf("failed to create migration: %w", err)
				}

				switch c.Args().First() {
				case "up":
					if err := migration.Up(c.Context); err != nil {
						return fmt.Errorf("failed to migrate up: %w", err)
					}
				case "down":
					if err := migration.Down(c.Context); err != nil {
						return fmt.Errorf("failed to migrate down: %w", err)
					}
				case "fresh":
					if err := migration.Fresh(c.Context); err != nil {
						return fmt.Errorf("failed to migrate fresh: %w", err)
					}
				default:
					if err := migration.Up(c.Context); err != nil {
						return fmt.Errorf("failed to migrate up: %w", err)
					}
				}
				log.Println("Migration succeed")
				return nil
			},
		},
	}

	return app
}

func main() {

	if err := App().Run(os.Args); err != nil {
		log.Fatal("Failed to run app", err)
	}

}
