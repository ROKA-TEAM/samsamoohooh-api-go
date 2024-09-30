package tests

import (
	"github.com/stretchr/testify/suite"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/repository"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/repository/tests/provider"
	"samsamoohooh-go-api/internal/core/port"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	Provider       *provider.Provider
	UserRepository port.UserRepository
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	s.Provider = provider.Provide()
	s.UserRepository = repository.NewUserRepository(s.Provider.Database)
}

func (s *UserRepositoryTestSuite) SetupTest() {
	err := s.Provider.Reset()
	s.Require().NoError(err)
}

//func (s *UserRepositoryTestSuite) TestGetUser() {
//	// 사용자 생성
//	user := &domain.User{Name: "Test User", Name: "test@example.com"}
//	err := s.repo.Create(user)
//	s.Require().NoError(err)
//
//	// 사용자 조회
//	fetchedUser, err := s.repo.GetByID(user.ID)
//	s.Require().NoError(err)
//	s.Equal(user.Name, fetchedUser.Name)
//	s.Equal(user.Email, fetchedUser.Email)
//}
