package integration_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/edward-/four-in-a-row-game/internal/tests/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateUser", func() {
	ctx := context.Background()
	ctx = utils.LoadCtx(ctx)

	BeforeEach(func() {
		utils.DeleteData()
	})

	AfterEach(func() {
		utils.DeleteData()
	})

	Describe("when create user", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			user := `{
				"nickName": "evega",
				"email": "edward.vega@live.com"
			}`

			req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(user))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should create an user successfully", func() {
			It("status code should be 201", func() {
				Expect(response.Code).To(Equal(http.StatusCreated))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})

			It("body should return an id", func() {
				var jsonResponse map[string]any
				json.Unmarshal(response.Body.Bytes(), &jsonResponse)
				Expect(jsonResponse["id"]).ToNot(BeNil())
			})
		})
	})

	Describe("when create user whitout email", func() {
		var response *httptest.ResponseRecorder

		JustBeforeEach(func() {
			user := `{
				"nickName": "user_test"
			}`

			req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(user))
			req.Header.Add("Content-Type", "application/json")
			response = utils.ExecuteRequest(ctx, req)
		})

		Context("should fail", func() {
			It("status code should be 400", func() {
				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})

			It("body should not be nil", func() {
				Expect(response.Body).ToNot(BeNil())
			})
		})
	})
})
