package api

import (
	db "backend-master-class/db/sqlc"
	"backend-master-class/token"
	"backend-master-class/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Marker
	router     *gin.Engine
}

func NewServer(c util.Config, store db.Store) (*Server, error) {
	tk, err := token.NewJWTMaker(c.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	s := &Server{store: store, tokenMaker: tk, config: c}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	s.setUpRouter()

	return s, nil
}

func (s *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users", s.CreateUser)
	router.POST("/users/login", s.LoginUser)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	authRoutes.POST("/accounts", s.CreateAccount)
	authRoutes.GET("/accounts/:id", s.GetAccount)
	authRoutes.GET("/accounts", s.ListAccount)

	authRoutes.POST("/transfers", s.CreateTransfer)

	s.router = router
}

func (s *Server) Start(addres string) error {
	return s.router.Run(addres)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
