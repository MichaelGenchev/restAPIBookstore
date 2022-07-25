package api

import (
	"database/sql"

	"net/http"

	db "github.com/MichaelGenchev/restAPIBookstore/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateAuthorRequest struct {
	Name      string `json:"name" bindning:"required"`
	Biography string `json:"biography" binding:"required"`
}

func (server *Server) createAuthor(ctx *gin.Context){

	var req CreateAuthorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAuthorParams{
		Name: req.Name,
		Biography: req.Biography,
	}
	
	account, err := server.store.CreateAuthor(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, account)
}


type getAccountRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}
type AuthorResponseModel struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
	Books     []db.Book `json:"books"`
}

func (server *Server) getAuthor(ctx *gin.Context){
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := server.store.GetAuthor(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 		
	}
	var returningAuthor AuthorResponseModel
	returningAuthor.ID = author.ID
	returningAuthor.Name = author.Name
	returningAuthor.Biography = author.Biography
	returningAuthor.Books = []db.Book{}

	booksByThisAuthor, err := server.store.GetBookByAuthor(ctx, returningAuthor.ID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusOK, returningAuthor)
			
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	returningAuthor.Books = booksByThisAuthor

	ctx.JSON(http.StatusOK, returningAuthor)
}



type listAuthorsResponse struct {
	Authors []listAuthorsModel `json:"authors"`
}
type listAuthorsModel struct {
	ID    int64		`json:"id"`
	Name  string 	`json:"name"`
	Biography string `json:"biography"`
	Books []db.Book  `json:"books"`
}

func (server *Server) listAuthors(ctx *gin.Context) {
	
	authors, err := server.store.ListAuthors(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var authorsResponse listAuthorsResponse

	for _, author := range authors {
		var listAuthorModel listAuthorsModel
		
		listAuthorModel.ID = author.ID
		listAuthorModel.Name = author.Name
		listAuthorModel.Biography = author.Biography
		
		booksByThisAuthor, err := server.store.GetBookByAuthor(ctx, author.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return 
		}

		listAuthorModel.Books = booksByThisAuthor

		authorsResponse.Authors = append(authorsResponse.Authors, listAuthorModel)
	}

	ctx.JSON(http.StatusOK, authorsResponse)
}


type updateAuthorRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}

type updateAuthorChangesBody struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

func (server *Server) updateAuthor(ctx *gin.Context) {

	var request updateAuthorRequest

	var changes updateAuthorChangesBody
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&changes); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAuthorParams {
		ID: request.ID,
		Name: changes.Name,
		Biography: changes.Biography,
	}

	updatedAuthor, err := server.store.UpdateAuthor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedAuthor)

}

type deleteAuthorRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}

func (server *Server) deleteAuthor(ctx *gin.Context) {
	var request deleteAuthorRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	deletedAuthor, err := server.store.DeleteAuthor(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, deletedAuthor)
}