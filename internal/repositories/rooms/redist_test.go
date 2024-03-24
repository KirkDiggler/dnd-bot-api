package rooms

import (
	"context"
	"encoding/json"
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

func (s *suiteRedis) SetupTest() {
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

func (s *suiteRedis) TearDownTest() {
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

func (s *suiteRedis) TestUpdateRoom() {

	existingRoom := &entities.Room{
		ID:          "123",
		Name:        "not test",
		Description: "test",
	}

	existingJson, _ := json.Marshal(existingRoom)

	expectedRoom := &entities.Room{
		ID:          "123",
		Name:        "test",
		Description: "test",
	}

	s.mockRedisClient.ExpectGet(getRoomKey("123")).SetVal(string(existingJson))

	s.mockRedisClient.ExpectSet(getRoomKey("123"), gomock.Any(), 0).SetVal("OK")

	input := &UpdateInput{
		Room: &entities.Room{
			ID:          "123",
			Name:        "test",
			Description: "not test",
		},
		InputMask: []string{"name"},
	}

	output, err := s.fixture.UpdateRoom(s.ctx, input)
	s.NoError(err)
	s.Equal(expectedRoom, output.Room)
}

func TestRedis(t *testing.T) {
	suite.Run(t, new(suiteRedis))
}
