package api

import (
	"fmt"

	db "github.com/Zinoshin/simplebank/db/sqlc"
	"github.com/Zinoshin/simplebank/token"
	"github.com/Zinoshin/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServe는 새로운 HTTP 서버를 생성하고 라우팅을 설정.
func NewServer(store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	authRoutes.POST("/transfers", server.createTransfer)
}

// Start runs the HTTP server on a specific address.
// http 서버를 실행하기 위한 코드
// 스타트 기능 추가 - 주소를 입력하고 에러를 반환 인풋주소에서 http서버를 실행해 api요청을 듣는것.
// 서버 실행
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// 오류메시지
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
