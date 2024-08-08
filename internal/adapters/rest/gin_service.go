package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mukappalambda/my-trader/internal/adapters/rest/types"
)

type SchemaRegistryServer struct {
	Engine *gin.Engine
}

// TODO: replace the in-memory database later
var schemas = make([]types.Schema, 0)

func NewSchemaRegistryServer(ginMode string) (*SchemaRegistryServer, error) {
	if ginMode != "debug" && ginMode != "release" {
		return nil, fmt.Errorf("got invalid gin mode: %q (want 'debug' or 'release')", ginMode)
	}
	if ginMode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	err := engine.SetTrustedProxies(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to set trusted proxies: %v", err)
	}
	engine.Use(cors.Default())
	setupRoutes(engine)
	engine.Use(HandlerNotFound())
	return &SchemaRegistryServer{
		Engine: engine,
	}, nil
}

func setupRoutes(engine *gin.Engine) {
	schemaRouterGroup := engine.Group("/schemas")
	{
		schemaRouterGroup.GET("", GetAllSchemas())
		schemaRouterGroup.POST("", NewSchema())
		schemaRouterGroup.GET("/:id", GetSchemaById())
	}
}

func HandlerNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	}
}

func GetAllSchemas() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, schemas)
	}
}

func NewSchema() gin.HandlerFunc {
	return func(c *gin.Context) {
		var schema types.Schema
		if err := c.BindJSON(&schema); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"reason": "invalid schema",
			})
			return
		}
		(&schema).SetId(fmt.Sprintf("%d", len(schemas)+1))
		schemas = append(schemas, schema)
		c.JSON(http.StatusCreated, schema)
	}
}

func GetSchemaById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}

func (srv *SchemaRegistryServer) Run(addr ...string) error {
	return srv.Engine.Run(addr...)
}
