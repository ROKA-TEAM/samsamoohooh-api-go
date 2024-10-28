# 프레젠터 (Presenter)

- API 모델을 도메인으로 변환을 담당합니다.

```go
// presenter_user_request.go
type UserUpdateByMeRequest struct {
 Name       string `json:"name"  validate:"min=0,max=15,omitempty"`
 Resolution string `json:"resolution" validate:"min=0,max=22,omitempty"`
}

func (r UserUpdateByMeRequest) ToDomain() *domain.User {
 return &domain.User{
  Name:       r.Name,
  Resolution: r.Resolution,
 }
}
```

- 도메인을 API 모델로 변환하여 클라이언트에게 최적화된 형식으로 제공합니다.

```go
// presenter_user_response.go
type UserGetByMeResponse struct {
 ID         int    `json:"id"`
 Name       string `json:"name"`
 Resolution string `json:"resolution"`
}

func NewUserGetByMeResponse(user *domain.User) *UserGetByMeResponse {
 return &UserGetByMeResponse{
  ID:         user.ID,
  Name:       user.Name,
  Resolution: user.Resolution,
 }
}
```

- 이로써, 응답 데이터를 최적화 할 수 있고, 도메인을 캡슐화 하여 보안에도 강화되고, API가 변경되도 Presneter만 변경하면 되는 유연성을 얻을 수 있었습니다.
