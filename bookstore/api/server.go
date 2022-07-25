package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/MichaelGenchev/restAPIBookstore/db/sqlc"
)


type Server struct {
	store *db.Store
	router  *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/authors", server.listAuthors)
	router.POST("/authors", server.createAuthor)
	router.GET("/authors/:id", server.getAuthor)
	router.PUT("/authors/:id", server.updateAuthor)
	router.DELETE("/authors/:id", server.deleteAuthor)

	router.GET("/books", server.listBooks)
	router.POST("/books", server.createBook)
	router.GET("/books/:id", server.getBook)
	router.PUT("/books/:id", server.updateBook)
	router.DELETE("/books/:id", server.deleteBook)

	router.GET("/categories", server.listCategories)
	router.POST("/categories", server.createCategory)
	router.GET("/categories/:id", server.getCategory)
	router.PUT("/categories/:id", server.updateCategory)
	router.DELETE("/categories/:id", server.deleteCategory)


	server.router = router	
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error ) gin.H {
	return gin.H{"error": err.Error()}
} 