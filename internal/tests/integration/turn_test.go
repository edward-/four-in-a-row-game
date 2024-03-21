package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

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

var _ = Describe("Turn", func() {
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

	Describe("when it's the first turn", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			turn := fmt.Sprintf(`{
				"userId": "%s",
				"dropItIn": %d
			}`, user2.Id, 1)

			path := fmt.Sprintf("/v1/games/%s/turn", game.Id)
			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(turn))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should complete turn successfuly", func() {
			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("body should return an id", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["resolution"]).To(Equal("Next"))
				Expect(jsonResponse["user_id"]).To(Equal(user1.Id))
			})
		})
	})

	Describe("when user id does not exist", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			turn := fmt.Sprintf(`{
				"userId": "%s",
				"dropItIn": %d
			}`, uuid.NewString(), 1)

			path := fmt.Sprintf("/v1/games/%s/turn", game.Id)
			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(turn))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should fail", func() {
			It("status code should be 500", func() {
				Expect(response.Code).To(Equal(http.StatusInternalServerError))
			})

			It("body should return message", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["message"]).To(Equal("could not do next move"))
			})
		})
	})

	Describe("when next turn", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			turn := fmt.Sprintf(`{
				"userId": "%s",
				"dropItIn": %d
			}`, user1.Id, 1)

			path := fmt.Sprintf("/v1/games/%s/turn", game.Id)
			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(turn))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)

			req = httptest.NewRequest(http.MethodPost, path, strings.NewReader(turn))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should fail if the same user send turn", func() {
			It("status code should be 500", func() {
				Expect(response.Code).To(Equal(http.StatusInternalServerError))
			})

			It("body should return message", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["message"]).To(Equal("could not do next move"))
			})
		})
	})

	Describe("when column is not in range", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			turn := fmt.Sprintf(`{
				"userId": "%s",
				"dropItIn": %d
			}`, user1.Id, 10)

			path := fmt.Sprintf("/v1/games/%s/turn", game.Id)
			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(turn))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should fail", func() {
			It("status code should be 500", func() {
				Expect(response.Code).To(Equal(http.StatusInternalServerError))
			})

			It("body should return message", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["message"]).To(Equal("could not do next move"))
			})
		})
	})
})
