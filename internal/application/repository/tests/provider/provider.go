package provider

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"fmt"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/repository"

	"samsamoohooh-go-api/internal/infra/storage/mysql"
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultTimeout = 5 * time.Second
	mysqlImage     = "mysql:8.0.36"
	mysqlPort      = "3306/tcp"
	mysqlPassword  = "1234"
	mysqlDatabase  = "test"
)

type Provider struct {
	engine testcontainers.Container
	*mysql.MySQL
	*sql.Driver
}

// NewProvider creates a new provider instance for testing
func NewProvider(ctx context.Context) (*Provider, error) {
	provider := &Provider{}
	if err := provider.onEngine(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialize engine: %w", err)
	}

	if err := provider.setConnect(ctx); err != nil {
		return nil, fmt.Errorf("failed to set connection: %w", err)
	}

	return provider, nil
}

func (p *Provider) GetUserRepository() port.UserRepository {
	return repository.NewUserRepository(p.MySQL)
}

// Setup initializes the database schema
func (p *Provider) Setup(ctx context.Context) error {
	if err := p.Client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	return nil
}

// TruncateTables removes all data from tables while keeping the schema
func (p *Provider) TruncateTables(ctx context.Context) error {
	// ent framework에서 제공하는 모든 테이블 목록
	tables := []string{"comments", "posts", "topics", "tasks", "user_groups", "groups", "users"}

	for _, table := range tables {

		if _, err := p.Driver.ExecContext(ctx, fmt.Sprintf("DROP TABLE `%s`", table)); err != nil {
			return fmt.Errorf("failed to truncate table %s: %w", table, err)
		}
	}

	return nil
}

// Shutdown closes connections and terminates the container
func (p *Provider) Shutdown(ctx context.Context) error {
	var errs []error

	if p.MySQL != nil {
		if err := p.MySQL.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close MySQL connection: %w", err))
		}
	}

	if p.engine != nil {
		if err := p.engine.Terminate(ctx); err != nil {
			errs = append(errs, fmt.Errorf("failed to terminate container: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("shutdown errors: %v", errs)
	}
	return nil
}

func (p *Provider) onEngine(ctx context.Context) error {
	req := testcontainers.ContainerRequest{
		Image:        mysqlImage,
		ExposedPorts: []string{mysqlPort},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": mysqlPassword,
			"MYSQL_DATABASE":      mysqlDatabase,
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}
	p.engine = mysqlContainer

	if err := p.setConnect(ctx); err != nil {
		return fmt.Errorf("failed to set connection: %w", err)
	}

	return nil
}

func (p *Provider) setConnect(ctx context.Context) error {
	host, err := p.engine.Host(ctx)
	if err != nil {
		return fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := p.engine.MappedPort(ctx, mysqlPort)
	if err != nil {
		return fmt.Errorf("failed to get mapped port: %w", err)
	}

	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?parseTime=True",
		mysqlPassword, host, port.Port(), mysqlDatabase)

	driver, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	p.Driver = driver
	client := ent.NewClient(ent.Driver(driver))
	p.MySQL = &mysql.MySQL{
		Client: client,
	}

	return nil
}
