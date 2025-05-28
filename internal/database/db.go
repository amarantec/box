package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func OpenConnection(ctx context.Context, dbConfig string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("create connection Pool: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timed out, trying to connect to the database: %w", ctx.Err())
		default:
			Conn, err = pgxpool.NewWithConfig(ctx, cfg)
			if err == nil {
				if err := Conn.Ping(ctx); err == nil {
					return Conn, nil
				}
				log.Printf("database not yet available, trying again in 2 seconds... [%v]\n", err)
			} else {
				log.Printf("error trying to connect to the database, trying again in 2 seconds... [%v]\n", err)
			}
			time.Sleep(2 * time.Second)

		}
	}
}
