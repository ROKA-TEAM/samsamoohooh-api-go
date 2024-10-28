# Storage

- 실지적으로 데이터를 저장하는 데이터베이스를 포함하는 파일입니다.

삼삼오오 서비스는 mysql, redis를 사용하기에 두개의 데이터베이스를 초기화 및 서비스 작성 코드를 내포하고 있습니다.

```go
// mysql.go

type MySQL struct {
 *ent.Client
}

func NewMySQL(config *config.Config) (*MySQL, error) {
 dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
  config.Database.User,
  config.Database.Password,
  config.Database.Host,
  config.Database.Port,
  config.Database.Database,
 )

 client, err := ent.Open("mysql", dsn)
 if err != nil {
  return nil, err
 }

 return &MySQL{client}, nil
}

func (d *MySQL) AutoMigration(ctx context.Context) error {
 err := d.Client.Schema.Create(ctx)
 if err != nil {
  return err
 }

 return nil
}

func (d *MySQL) Close() error {
 return d.Client.Close()
}

```

```go
// redis.go
type Redis struct {
 client *redis.Client
}

func NewRedis(config *config.Config) (*Redis, error) {
 client := redis.NewClient(&redis.Options{
  Addr:     config.Redis.Addr,
  Password: config.Redis.Password,
  DB:       config.Redis.DB,
 })

 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()

 _, err := client.Ping(ctx).Result()
 if err != nil {
  return nil, err
 }

 return &Redis{client}, nil
}

func (r *Redis) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
 err := r.client.Set(ctx, key, value, ttl).Err()
 if err != nil {
  return exception.New(
   err,
   exception.ErrRedisSetFailed,
   exception.StatusInternalServerError,
   "failed to set redis",
   exception.WithData(
    exception.Map{
     "key":   key,
     "value": value,
    },
   ),
  )
 }

 return nil
}
```
