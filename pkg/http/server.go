package http

import (
	"github.com/TicketsBot/logarchiver/internal"
	"github.com/TicketsBot/logarchiver/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v6"
	"go.uber.org/zap"
)

type Server struct {
	Logger      *zap.Logger
	Config      config.Config
	RemoveQueue internal.RemoveQueue
	router      *gin.Engine
	client      *minio.Client
}

func NewServer(logger *zap.Logger, config config.Config, client *minio.Client) *Server {
	return &Server{
		Logger:      logger,
		Config:      config,
		RemoveQueue: internal.NewRemoveQueue(logger),
		router:      gin.Default(),
		client:      client,
	}
}

func (s *Server) RegisterRoutes() {
	s.router.LoadHTMLGlob("./public/templates/*")

	s.router.POST("/encode", encodeHandler)

	s.router.GET("/", s.ticketGetHandler)
	s.router.POST("/", s.ticketUploadHandler)

	s.router.GET("/guild/status/:id", s.purgeStatusHandler)
	s.router.DELETE("/guild/:id", s.purgeGuildHandler)
}

func (s *Server) Start() {
	if err := s.router.Run(s.Config.Address); err != nil {
		panic(err)
	}
}
