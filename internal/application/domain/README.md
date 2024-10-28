# 도메인 (Domain) 레이

- 도메인 레이어는 핵심 비즈니스 로직과 엔티티를 정의하며, 데이터베이스 연관관계 또한 고려해 작성합니다.
- 또한 의존성이 전혀 없으며, 순수한 상태로 존재해야 합니다.

## 예시 코드

[domain_user.go](https://github.com/ROKA-TEAM/samsamoohooh-go-api/blob/pinned/2024_10_28/internal/application/domain/domain_user.go)

```go
type User struct {
 ID         int
 Name       string
 Resolution string
 Role       UserRoleType
 Social     UserSocialType
 SocialSub  string
 CreatedAt  time.Time
 UpdatedAt  time.Time
 DeletedAt  time.Time

 // relation
 Groups   []*Group
 Topics   []*Topic
 Posts    []*Post
 Comments []*Comment
}
```
