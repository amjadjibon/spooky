package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/urfave/cli/v2"
)

/*
spooky postgres --user=baseuser --pass=d4xLtYSmcw --host=localhost --port=5432 --db_name=postgres --sslmode=disable ping
*/

func connectPG(user, password, host, port, dbName, sslmode string) (*pg.DB, error) {
	opt, err := pg.ParseURL(
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			user, password, host, port, dbName, sslmode))
	if err != nil {
		return nil, err
	}
	return pg.Connect(opt), nil
}

func getPGFlags(ctx *cli.Context) error {

	return nil
}

func pingDB(user, password, host, port, dbName, sslmode string) error {
	db, err := connectPG(user, password, host, port, dbName, sslmode)
	if err != nil {
		return err
	}

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		return err
	}

	fmt.Println("pong")

	return nil
}

func PingPostgres(ctx *cli.Context) error {
	user := ctx.String("user")
	pass := ctx.String("pass")
	host := ctx.String("host")
	port := ctx.String("port")
	dbName := ctx.String("db_name")
	sslmode := ctx.String("sslmode")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "5432"
	}

	if dbName == "" {
		dbName = "postgres"
	}

	if sslmode == "" {
		sslmode = "disable"
	}

	if err := pingDB(
		user,
		pass,
		host,
		port,
		dbName,
		sslmode); err != nil {
		return err
	}
	return nil
}

func getVersion(user, password, host, port, dbName, sslmode string) error {
	db, err := connectPG(user, password, host, port, dbName, sslmode)
	if err != nil {
		return err
	}

	ctx := context.Background()

	var version string
	_, err = db.QueryOneContext(ctx, pg.Scan(&version), "SELECT version()")
	if err != nil {
		return err
	}

	fmt.Println(version)

	return nil
}

func PostgresVersion(ctx *cli.Context) error {
	user := ctx.String("user")
	pass := ctx.String("pass")
	host := ctx.String("host")
	port := ctx.String("port")
	dbName := ctx.String("db_name")
	sslmode := ctx.String("sslmode")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "5432"
	}

	if dbName == "" {
		dbName = "postgres"
	}

	if sslmode == "" {
		sslmode = "disable"
	}

	err := getVersion(user, pass, host, port, dbName, sslmode)
	if err != nil {
		return err
	}

	return nil
}