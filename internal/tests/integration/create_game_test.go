package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres/model"
	"github.com/edward-/four-in-a-row-game/internal/tests/utils"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateGame", func() {
	ctx := context.Background()
	ctx = utils.LoadCtx(ctx)
	var user1 *model.User
	var user2 *model.User

	BeforeEach(func() {
		utils.DeleteData()

		db := contextPkg.DatabaseFromCtx(ctx)

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
	})

	AfterEach(func() {
		utils.DeleteData()
	})

	Describe("when create game", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			games := fmt.Sprintf(`{
				"userId1": "%s",
				"userId2": "%s"
			}`, user1.Id, user2.Id)

			req := httptest.NewRequest(http.MethodPost, "/v1/games", strings.NewReader(games))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should create an user successfuly", func() {
			It("status code should be 201", func() {
				Expect(response.Code).To(Equal(http.StatusCreated))
			})

			It("body should return an id", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["id"]).ToNot(BeNil())
			})
		})
	})

	Describe("when create game whit user id does not exist", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			games := fmt.Sprintf(`{
				"userId1": "%s",
				"userId2": "%s"
			}`, uuid.NewString(), user2.Id)

			req := httptest.NewRequest(http.MethodPost, "/v1/games", strings.NewReader(games))
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
				Expect(jsonResponse["message"]).To(Equal("could not create the game"))
			})
		})
	})

	Describe("when create game whit body request is wrong", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			games := fmt.Sprintf(`{
				"userId2": "%s"
			}`, user2.Id)

			req := httptest.NewRequest(http.MethodPost, "/v1/games", strings.NewReader(games))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should fail", func() {
			It("status code should be 400", func() {
				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})

			It("body should return message", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["message"]).To(Equal("body invalid"))
			})
		})
	})
})
