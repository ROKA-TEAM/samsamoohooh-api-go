# 리포지토리(Repository) Layer

- MySQL과 Redis와 같은 데이터베이스와의 상호작용을 추상화하는 역활을 담당합니다.

```go
// repository_user.go
type UserRepository struct {
 database *mysql.MySQL
}

func NewUserRepository(database *mysql.MySQL) *UserRepository {
 return &UserRepository{
  database: database,
 }
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
 createdUser, err := r.database.User.
  Create().
  SetName(user.Name).
  SetResolution(user.Resolution).
  SetRole(entuser.Role(user.Role)).
  SetSocial(entuser.Social(user.Social)).
  SetSocialSub(user.SocialSub).
  Save(ctx)

 if err != nil {
  return nil, utils.Wrap(err)
 }

 return utils.ConvertDomainUser(createdUser), nil
}

```
