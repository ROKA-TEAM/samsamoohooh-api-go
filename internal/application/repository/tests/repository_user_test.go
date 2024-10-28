package tests

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/repository/tests/provider"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositorySuite struct {
	*provider.Provider
	suite.Suite
}

// 전체 테스트 하기 전에 한 번만 실행된다.
func (s *UserRepositorySuite) SetupSuite() {
	// provider를 초기화 한다.
	// 가상 conatiner 생성
	// MySQL 연결
	p, err := provider.NewProvider(context.Background())
	if err != nil {
		panic(err)
	}

	s.Provider = p
}

// 각 테스트 실행 전에 한 번만 실행된다.
func (s *UserRepositorySuite) SetupTest() {
	// database auto migration
	if err := s.AutoMigration(context.Background()); err != nil {
		panic(err)
	}
}

// 각 테스트가 끝난 후 실행후에 실행된다.
func (s *UserRepositorySuite) TearDownTest() {
	// database truncate
	// 가상 container 종료
	if err := s.TruncateTables(context.Background()); err != nil {
		panic(err)
	}
}

// 모든 테스트가 끝난 후에 실행된다.
func (s *UserRepositorySuite) TearDownSuite() {
	// 가상 container 종료
	if err := s.Provider.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}

func TestRunUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

func (s *UserRepositorySuite) TestCreateUser() {
	testCases := []struct {
		name    string
		user    *domain.User
		wantErr bool
	}{
		{
			name: "성공_기본_유저_생성",
			user: &domain.User{
				Name:       "Test User",
				Resolution: "My Resolution",
				Role:       domain.UserRoleUser,
				Social:     domain.UserSocialGoogle,
				SocialSub:  "google123",
			},
			wantErr: false,
		},
		{
			name: "성공_어드민_유저_생성",
			user: &domain.User{
				Name:       "Admin User",
				Resolution: "Admin Resolution",
				Role:       domain.UserRoleAdmin,
				Social:     domain.UserSocialKaKao,
				SocialSub:  "kakao123",
			},
			wantErr: false,
		},
		{
			name: "성공_매니저_유저_생성",
			user: &domain.User{
				Name:       "Manager User",
				Resolution: "Manager Resolution",
				Role:       domain.UserRoleManager,
				Social:     domain.UserSocialApple,
				SocialSub:  "apple123",
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// When
			createdUser, err := s.GetUserRepository().CreateUser(context.Background(), tc.user)

			// Then
			if tc.wantErr {
				assert.Error(s.T(), err)
				return
			}

			assert.NoError(s.T(), err)
			assert.NotNil(s.T(), createdUser)
			assert.NotZero(s.T(), createdUser.ID)
			assert.Equal(s.T(), tc.user.Name, createdUser.Name)
			assert.Equal(s.T(), tc.user.Resolution, createdUser.Resolution)
			assert.Equal(s.T(), tc.user.Role, createdUser.Role)
			assert.Equal(s.T(), tc.user.Social, createdUser.Social)
			assert.Equal(s.T(), tc.user.SocialSub, createdUser.SocialSub)
			assert.NotZero(s.T(), createdUser.CreatedAt)
			assert.NotZero(s.T(), createdUser.UpdatedAt)
			assert.Zero(s.T(), createdUser.DeletedAt)

			// Verify the user was actually persisted
			persistedUser, err := s.GetUserRepository().GetByUserID(context.Background(), createdUser.ID)
			assert.NoError(s.T(), err)
			assert.Equal(s.T(), createdUser, persistedUser)
		})
	}
}

func (s *UserRepositorySuite) TestGetByUserID() {
	// 먼저 테스트를 위한 사용자 데이터를 생성합니다.
	testUser := &domain.User{
		Name:       "Test User",
		Resolution: "My Resolution",
		Role:       domain.UserRoleUser,
		Social:     domain.UserSocialGoogle,
		SocialSub:  "google123",
	}

	createdUser, _ := s.GetUserRepository().CreateUser(context.Background(), testUser)

	testCases := []struct {
		name    string
		userID  int
		wantErr bool
	}{
		{
			name:    "성공_유저_ID로_조회",
			userID:  createdUser.ID,
			wantErr: false,
		},
		{
			name:    "실패_존재하지_않는_ID",
			userID:  9999, // 존재하지 않는 ID
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// When
			gotUser, err := s.GetUserRepository().GetByUserID(context.Background(), tc.userID)

			// Then
			if tc.wantErr {
				assert.Error(s.T(), err)
				assert.Nil(s.T(), gotUser)
				return
			}

			assert.NoError(s.T(), err)
			assert.NotNil(s.T(), gotUser)
			assert.Equal(s.T(), createdUser.ID, gotUser.ID)
			assert.Equal(s.T(), createdUser.Name, gotUser.Name)
			assert.Equal(s.T(), createdUser.Resolution, gotUser.Resolution)
			assert.Equal(s.T(), createdUser.Role, gotUser.Role)
			assert.Equal(s.T(), createdUser.Social, gotUser.Social)
			assert.Equal(s.T(), createdUser.SocialSub, gotUser.SocialSub)
		})
	}
}

func (s *UserRepositorySuite) TestUpdateUser() {
	// 먼저 테스트를 위한 사용자 데이터를 생성합니다.
	testUser := &domain.User{
		Name:       "Test User",
		Resolution: "My Resolution",
		Role:       domain.UserRoleUser,
		Social:     domain.UserSocialGoogle,
		SocialSub:  "google123",
	}

	createdUser, _ := s.GetUserRepository().CreateUser(context.Background(), testUser)

	testCases := []struct {
		name        string
		userID      int
		updatedData *domain.User
		wantErr     bool
	}{
		{
			name:   "성공_유저_정보_업데이트",
			userID: createdUser.ID,
			updatedData: &domain.User{
				Name:       "Updated User",
				Resolution: "Updated Resolution",
				Role:       domain.UserRoleAdmin,
				Social:     domain.UserSocialKaKao,
				SocialSub:  "kakao123",
			},
			wantErr: false,
		},
		{
			name:   "실패_존재하지_않는_ID_업데이트",
			userID: 9999, // 존재하지 않는 ID
			updatedData: &domain.User{
				Name: "New User",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// When
			updatedUser, err := s.GetUserRepository().UpdateUser(context.Background(), tc.userID, tc.updatedData)

			// Then
			if tc.wantErr {
				assert.Error(s.T(), err)
				assert.Nil(s.T(), updatedUser)
				return
			}

			assert.NoError(s.T(), err)
			assert.NotNil(s.T(), updatedUser)
			assert.Equal(s.T(), tc.updatedData.Name, updatedUser.Name)
			assert.Equal(s.T(), tc.updatedData.Resolution, updatedUser.Resolution)
			assert.Equal(s.T(), tc.updatedData.Role, updatedUser.Role)
			assert.Equal(s.T(), tc.updatedData.Social, updatedUser.Social)
			assert.Equal(s.T(), tc.updatedData.SocialSub, updatedUser.SocialSub)
		})
	}
}

func (s *UserRepositorySuite) TestDeleteUser() {
	// 먼저 테스트를 위한 사용자 데이터를 생성합니다.
	testUser := &domain.User{
		Name:       "Test User",
		Resolution: "My Resolution",
		Role:       domain.UserRoleUser,
		Social:     domain.UserSocialGoogle,
		SocialSub:  "google123",
	}

	createdUser, _ := s.GetUserRepository().CreateUser(context.Background(), testUser)

	testCases := []struct {
		name    string
		userID  int
		wantErr bool
	}{
		{
			name:    "성공_유저_삭제",
			userID:  createdUser.ID,
			wantErr: false,
		},
		{
			name:    "실패_존재하지_않는_ID_삭제",
			userID:  9999, // 존재하지 않는 ID
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// When
			err := s.GetUserRepository().DeleteUser(context.Background(), tc.userID)

			// Then
			if tc.wantErr {
				assert.Error(s.T(), err)
				return
			}

			assert.NoError(s.T(), err)

			// Verify the user was actually deleted
			gotUser, err := s.GetUserRepository().GetByUserID(context.Background(), tc.userID)
			assert.Error(s.T(), err)   // 에러가 발생해야 함
			assert.Nil(s.T(), gotUser) // 삭제된 사용자는 nil 이어야 함
		})
	}
}
