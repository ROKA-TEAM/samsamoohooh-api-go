# 포트(Port) 레이어

- Hexagonal Architecture에서 착안한 아이디어로, 의존성 역전(DIP)을 실현하는 중요한 역할을 담당합니다.

- 파일들에서 확인할 수 있듯이 본 프로젝트에서 필요한 모든 로직들을 interface로 정의한 것을 확인할 수 있습니다. 이를 통해 마치 전자제품을 아무거나 사도 포트가 같기 때문에 충전을 할 수 있듯이, port 레이어 덕분에 확장성과 유지보수성애 매우 증가합니다.

## 예시 코드

[port_user.go](https://github.com/ROKA-TEAM/samsamoohooh-go-api/blob/pinned/2024_10_28/internal/application/port/port_user.go)

```go
// userRepository interface를 정의한 코드 
type UserRepository interface {
 CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)

 GetByUserID(ctx context.Context, id int) (*domain.User, error)

 GetByUserSub(ctx context.Context, sub string) (*domain.User, error)

 GetGroupsByUserID(ctx context.Context, id int, limit, offset int) ([]*domain.Group, error)

 GetUsers(ctx context.Context, limit, offset int) ([]*domain.User, error)

 UpdateUser(ctx context.Context, id int, user *domain.User) (*domain.User, error)

 DeleteUser(ctx context.Context, id int) error

 IsUserInGroup(ctx context.Context, userID, groupID int) (bool, error)
}


// userService 생성할 때 port.UserRepository(interface)를 주입받는 모습 
type UserService struct {
 userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
 return &UserService{
  userRepository: userRepository,
 }
}

// port.userRepository를 구현하는 객체의 모습 
var _ port.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
 database *mysql.MySQL
}

func NewUserRepository(database *mysql.MySQL) *UserRepository {
 return &UserRepository{
  database: database,
 }
}

...implements!

```

![예시 이미지](/docs/images/port_and_adapter_example.png)

만약 더 자세한 정보를 원한다면 [hexagonal-startup](https://github.com/fullgukbap/hexagonal-startup) 리포지토리를 참고해주세요.
