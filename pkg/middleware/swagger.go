package middleware

import (
	"context"
	"net/http"

	contextPkg "github.com/edward-/four-in-a-row-game/pkg/context"
	routerGin "github.com/edward-/four-in-a-row-game/pkg/routers/gin"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
)

func SwaggerValidation(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		loader := &openapi3.Loader{Context: c, IsExternalRefsAllowed: true}
		doc, _ := loader.LoadFromFile(".../My-OpenAPIv3-API.yml")
		log := contextPkg.LoggerFromCtx(ctx)

		err := doc.Validate(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Error(err, "swagger validation error")
			return
		}
		// router, _ := gorillamux.NewRouter(doc)
		router, err := routerGin.NewRouter(doc)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Error(err, "error creating swagger router")
			return
		}

		// Find route
		route, pathParams, err := router.FindRoute(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Error(err, "route not found")
			return
		}

		// Validate request
		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    c.Request,
			PathParams: pathParams,
			Route:      route,
		}
		err = openapi3filter.ValidateRequest(ctx, requestValidationInput)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.Error(err, "swagger validation error")
			return
		}

		c.Next()
	}
}
