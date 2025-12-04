package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
<<<<<<< HEAD
	defaultMaxPoolSize  = 1
=======
	defaultMaxPoolSize  = 20
>>>>>>> new
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	GormDB *gorm.DB
}

func New(url string, opts ...Option) (*Postgres, error) {

	pg := &Postgres{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, option := range opts {
		option(pg)
	}

	var err error = nil
	for ; pg.connAttempts > 0; pg.connAttempts-- {
<<<<<<< HEAD
		log.Println(url)
=======
>>>>>>> new
		pg.GormDB, err = gorm.Open(postgres.New(postgres.Config{
			DSN: url,
		}), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)
		time.Sleep(pg.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("error gorm.Open: %w", err)
	}

	sqlDB, err := pg.GormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(pg.maxPoolSize / 10)
	sqlDB.SetMaxOpenConns(pg.maxPoolSize)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	sqlDB, err := p.GormDB.DB()
	if err != nil {
		sqlDB.Close()
	}
}

func (p *Postgres) AutoMigrate(models ...interface{}) error {
	if p.GormDB == nil {
		return fmt.Errorf("database not initialized")
	}
	return p.GormDB.AutoMigrate(models...)
}
