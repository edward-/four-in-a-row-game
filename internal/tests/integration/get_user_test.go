package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/edward-/four-in-a-row-game/internal/infrastructure/repository/postgres/model"
	"github.com/edward-/four-in-a-row-game/internal/tests/utils"
	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	"github.com/google/uuid"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetUser", func() {
	ctx := context.Background()
	ctx = utils.LoadCtx(ctx)
	var user *model.User

	BeforeEach(func() {
		utils.DeleteData()

		db := contextPkg.DatabaseFromCtx(ctx)

		// create users
		user = &model.User{
			NickName: "user1",
			Email:    "user1@mail.com",
		}

		db.Create(&user)
	})

	AfterEach(func() {
		utils.DeleteData()
	})

	Describe("when get user", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			path := fmt.Sprintf("/v1/users/%s", user.Id)
			req := httptest.NewRequest(http.MethodGet, path, nil)
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should get board successfuly", func() {
			It("status code should be 200", func() {
				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("body should return an user", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["id"]).NotTo(BeNil())
				Expect(jsonResponse["nickName"]).To(Equal(user.NickName))
				Expect(jsonResponse["email"]).To(Equal(user.Email))
			})
		})
	})

	Describe("when user id does not exist", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			path := fmt.Sprintf("/v1/users/%s", uuid.NewString())
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
				Expect(jsonResponse["message"]).To(Equal("error getting user: user not found"))
			})
		})
	})
})
