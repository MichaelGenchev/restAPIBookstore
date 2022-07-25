package api

import (
	"database/sql"
	"net/http"

	db "github.com/MichaelGenchev/restAPIBookstore/db/sqlc"
	"github.com/gin-gonic/gin"
)


type CreateBookRequest struct {
	Title      string `json:"title" bindning:"required"`
	Author 		int64 `json:"author" binding:"required"`
	Category 	int64 `json:"category" binding:"required"`
	Price 		float64 `json:"price" binding:"required"`
}

type BookResponse struct {
	ID        int64  `json:"id"`
	Title      string `json:"title"`
	Author 		db.Author `json:"author"`
	Category 	db.Category `json:"category"`
	Price 		float64 `json:"price"`

}


func (server *Server) createBook(ctx *gin.Context){

	var bookRequest CreateBookRequest

	if err := ctx.ShouldBindJSON(&bookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBookParams{
		Title: bookRequest.Title,
		Author: bookRequest.Author,
		Category: bookRequest.Category,
		Price: bookRequest.Price,
	}
	
	book, err := server.store.CreateBook(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	var response BookResponse

	thisBookAuthor , err := server.store.GetAuthor(ctx, book.Author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	
	response.Author = thisBookAuthor

	thisBookCategory, err := server.store.GetCategory(ctx, book.Category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	response.Category = thisBookCategory

	response.ID = book.ID
	response.Title = book.Title
	response.Price = book.Price
	
	ctx.JSON(http.StatusOK, response)

}

type listBooksResponse struct {
	Books []listBookModel `json:"books"`

}
type listBookModel struct {
	ID       int64   `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Author db.Author `json:"author"`
	Category db.Category `json:"category"`
}
func (server *Server) listBooks(ctx *gin.Context) {

	books , err := server.store.ListBooks(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	var booksResponse listBooksResponse

	for _, book := range books {
		var listBookModel listBookModel

		listBookModel.ID = book.ID
		listBookModel.Title = book.Title
		listBookModel.Price = book.Price

		thisBookAuthor, err := server.store.GetAuthor(ctx, book.Author)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		listBookModel.Author = thisBookAuthor

		thisBookCategory, err := server.store.GetCategory(ctx, book.Category)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		listBookModel.Category = thisBookCategory

		booksResponse.Books = append(booksResponse.Books, listBookModel)
	}

	ctx.JSON(http.StatusOK, booksResponse)
}


type getBookRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}

type BookResponseModel struct {
	ID 		 int64	 `json:"id"`
	Title    string  `json:"title"`
	Author   db.Author   `json:"author"`
	Category db.Category   `json:"category"`
	Price    float64 `json:"price"`
}

func (server *Server) getBook(ctx *gin.Context) {

	var request getBookRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetBook(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 		
	}

	var responseBook BookResponseModel

	responseBook.ID = book.ID
	responseBook.Title = book.Title
	responseBook.Price = book.Price

	thisBookAuthor , err := server.store.GetAuthor(ctx, book.Author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	responseBook.Author = thisBookAuthor

	thisBookCategory, err := server.store.GetCategory(ctx, book.Category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	responseBook.Category = thisBookCategory

	ctx.JSON(http.StatusOK, responseBook)
}

type updateBookRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}
type updateBookChangesBody struct {
	Price 	  float64 `json:"price"`
}
func (server *Server) updateBook(ctx *gin.Context){

	var request updateBookRequest

	var changes updateBookChangesBody
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&changes); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateBookParams {
		ID: request.ID,
		Price: changes.Price,
	}

	updatedBook, err := server.store.UpdateBook(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var responseBook BookResponse

	responseBook.ID = updatedBook.ID
	responseBook.Title = updatedBook.Title
	responseBook.Price = updatedBook.Price

	thisBookAuthor , err := server.store.GetAuthor(ctx, updatedBook.Author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	responseBook.Author = thisBookAuthor

	thisBookCategory, err := server.store.GetCategory(ctx, updatedBook.Category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	responseBook.Category = thisBookCategory

	ctx.JSON(http.StatusOK, responseBook)
}

type deleteBookRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}

func (server *Server) deleteBook(ctx *gin.Context) {

	var request deleteBookRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	deletedBook , err := server.store.DeleteBook(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	var responseBook BookResponse
	responseBook.ID = deletedBook.ID
	responseBook.Title = deletedBook.Title
	responseBook.Price = deletedBook.Price

	thisBookAuthor , err := server.store.GetAuthor(ctx, deletedBook.Author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	responseBook.Author = thisBookAuthor

	thisBookCategory, err := server.store.GetCategory(ctx, deletedBook.Category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	responseBook.Category = thisBookCategory

	ctx.JSON(http.StatusOK, responseBook)


}