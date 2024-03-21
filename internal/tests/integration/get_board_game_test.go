package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/edward-/four-in-a-row-game/internal/domain/entity"
	vo "github.com/edward-/four-in-a-row-game/internal/domain/value_object"
	modelCache "github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/cache/model"
	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres/model"
	"github.com/edward-/four-in-a-row-game/internal/tests/utils"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	"github.com/google/uuid"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetBoardGame", func() {
	ctx := context.Background()
	ctx = utils.LoadCtx(ctx)
	var user1 *model.User
	var user2 *model.User
	var game *model.Game
	var board *modelCache.Board

	BeforeEach(func() {
		utils.DeleteData()

		db := contextPkg.DatabaseFromCtx(ctx)
		cache := contextPkg.CacheFromCtx(ctx)

		// create users
		user1 = &model.User{
			NickName: "user1",
			Email:    "user1@mail.com",
		}
		user2 = &model.User{
			NickName: "user2",
			Email:    "user2@mail.com",
		}
		db.Create(&user1)
		db.Create(&user2)

		// create game
		game = &model.Game{
			UserId1: user1.Id,
			UserId2: user2.Id,
		}
		db.Create(&game)

		// create board in cache
		board = modelCache.BoardToModel(entity.NewBoard())
		boardBytes, _ := json.Marshal(&board)
		cache.Set(ctx, game.Id, boardBytes, vo.ActiveGame)
	})

	AfterEach(func() {
		utils.DeleteData()
	})

	Describe("when get board game", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			path := fmt.Sprintf("/v1/games/%s/board", game.Id)
			req := httptest.NewRequest(http.MethodGet, path, nil)
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should get board successfuly", func() {
			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("body should return an array of 2D", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)

				col := jsonResponse["squares"].([]any)
				Expect(len(col)).To(Equal(7))

				for _, c := range col {
					row := c.([]any)
					Expect(len(row)).To(Equal(6))
				}
			})
		})
	})

	Describe("when game id does not exist", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			path := fmt.Sprintf("/v1/games/%s/board", uuid.NewString())
			req := httptest.NewRequest(http.MethodGet, path, nil)
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should fail", func() {
			It("status code should be 500", func() {
				Expect(response.Code).To(Equal(http.StatusInternalServerError))
			})

			It("body should return message", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["message"]).To(Equal("could not get the board"))
			})
		})
	})
})
