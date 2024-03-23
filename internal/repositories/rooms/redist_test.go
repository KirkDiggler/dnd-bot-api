package rooms

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/KirkDiggler/dnd-bot-api/internal/entities"
	"github.com/go-redis/redismock/v9"

	mock_uuider "github.com/KirkDiggler/dnd-bot-api/internal/common/mock"

	"github.com/stretchr/testify/suite"
)

type suiteRedis struct {
	suite.Suite

	ctx context.Context

	mockController *gomock.Controller
	mockUUID       *mock_uuider.MockUUIDGenerator

	mockRedisClient redismock.ClientMock

	fixture *Redis
}

func (s *suiteRedis) SetupSuite() {
	s.ctx = context.Background()
	s.mockController = gomock.NewController(s.T())
	s.mockUUID = mock_uuider.NewMockUUIDGenerator(s.mockController)

	db, mock := redismock.NewClientMock()
	s.mockRedisClient = mock

	s.fixture = &Redis{
		client:        db,
		uuidGenerator: s.mockUUID,
	}
}

func (s *suiteRedis) TearDownSuite() {
	s.mockController.Finish()
}

func (s *suiteRedis) TestCreateRoom() {
	s.mockUUID.EXPECT().New().Return("123")

	room := &entities.Room{
		Name:        "test",
		Description: "test",
	}

	s.mockRedisClient.ExpectSet(getRoomKey("123"), room, 0).SetVal("OK")
	s.mockRedisClient.ExpectSAdd(getRoomListKey(), "room:123").SetVal(int64(1))

	err := s.fixture.CreateRoom(s.ctx, room)
	s.NoError(err)
}

func TestRedis(t *testing.T) {
	suite.Run(t, new(suiteRedis))
}