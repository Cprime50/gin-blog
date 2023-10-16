package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Cprime50/gin-blog/config"
	db "github.com/Cprime50/gin-blog/db/sqlc"
	"github.com/Cprime50/gin-blog/docs"
	"github.com/Cprime50/gin-blog/logger"
	"github.com/Cprime50/gin-blog/search"
)

type Server struct {
	config config.Config
	router *gin.Engine
	store  db.Store
	search search.Searcher
	log    logger.Logger
}

func NewServer(config config.Config, store db.Store, tsHandler search.Searcher, log logger.Logger) *Server {
	var engine *gin.Engine
	if config.Environment == "test" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("test environment detected")
		engine = gin.New()
	} else {
		engine = gin.Default()
	}
	server := &Server{
		config: config,
		router: engine,
		store:  store,
		search: tsHandler,
		log:    log,
	}
	return server
}

func (s *Server) MountHandlers() {
	api := s.router.Group("/api")
	api.POST("/users", s.RegisterUser)
	api.POST("/users/login", s.LoginUser)

	user := api.Group("/user")
	user.Use(AuthMiddleware())
	user.GET("", s.GetCurrentUser)
	user.PUT("", s.UpdateUser)

	profiles := api.Group("/profiles")
	profiles.Use(AuthMiddleware())
	profiles.GET("/:username", s.GetProfile)
	profiles.POST("/:username/follow", s.FollowUser)
	profiles.DELETE("/:username/follow", s.UnfollowUser)

	articles := api.Group("/articles")
	articles.GET("/search", s.SearchArticles)
	articles.GET("", s.ListArticles)
	articles.GET("/:slug", s.GetArticle)
	articles.GET("/:slug/comments", s.GetComments)
	articles.Use(AuthMiddleware())
	articles.POST("", s.CreateArticle)
	articles.GET("/feed", s.FeedArticles)
	articles.PUT("/:slug", s.UpdateArticle)
	articles.DELETE("/:slug", s.DeleteArticle)
	articles.POST("/:slug/comments", s.AddComment)
	articles.DELETE("/:slug/comments/:id", s.DeleteComment)
	articles.POST("/:slug/favorite", s.FavoriteArticle)
	articles.DELETE("/:slug/favorite", s.UnfavoriteArticle)

	tags := api.Group("/tags")
	tags.GET("", s.GetTags)
}

func (s *Server) MountSwaggerHandlers() {
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Title = "ThelsBlog API"
	docs.SwaggerInfo.Description = "thelsblog API Documentation"
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) findUniqueSlug(c *gin.Context, title string) (string, error) {
	var (
		found      bool
		uniqueSlug string
	)
	for !found {
		uniqueSlug = createUniqueSlug(title)
		articleID, err := NullableID(s.store.GetArticleIDBySlug(c, uniqueSlug))
		if err != nil {
			return "", err
		}
		if articleID == "" {
			found = true
		}
	}
	return uniqueSlug, nil
}
