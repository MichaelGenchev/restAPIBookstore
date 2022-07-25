package api

import (
	"database/sql"
	"net/http"

	db "github.com/MichaelGenchev/restAPIBookstore/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" bindning:"required"`
}

func (server *Server) createCategory(ctx *gin.Context){
	var request CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createdCategory, err := server.store.CreateCategory(ctx, request.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, createdCategory)

}

type listCategoriesResponse struct {
	Categories []listCategoryModel `json:"categories"`
}

type listCategoryModel struct {
	ID    int64		`json:"id"`
	Name  string 	`json:"name"`
	Books []db.Book `json:"books"`
}

func (server *Server) listCategories(ctx *gin.Context){

	categories, err := server.store.ListCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	var categoriesResponse listCategoriesResponse

	for _, category := range categories {

		var listCategoryModel listCategoryModel

		listCategoryModel.ID = category.ID
		listCategoryModel.Name = category.Name

		booksForThisCategory, err := server.store.GetBookByCategory(ctx, category.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return 
		}

		listCategoryModel.Books = booksForThisCategory


		categoriesResponse.Categories = append(categoriesResponse.Categories, listCategoryModel)
	}

	ctx.JSON(http.StatusOK, categoriesResponse)
}


type getCategoryRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}

func(server *Server) getCategory(ctx *gin.Context){
	var request getAccountRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 		
	}

	ctx.JSON(http.StatusOK, category)
}

type updateCategoryRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}

type updateCategoryChangesBody struct {
	Name string `json:"name"`
}

func(server *Server) updateCategory(ctx *gin.Context){

	var request updateCategoryRequest

	var changes updateCategoryChangesBody
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&changes); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCategoryParams {
		ID: request.ID,
		Name: changes.Name,
	}

	updatedCategory, err := server.store.UpdateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)

}

type deleteCategoryRequest struct {
	ID int64 `uri:"id" bindning:"required,min=1"`
}
func(server *Server) deleteCategory(ctx *gin.Context){

	var request deleteCategoryRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	booksForThisCategory , err := server.store.GetBookByCategory(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	for _, book := range booksForThisCategory {
		_, err := server.store.DeleteBook(ctx, book.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return 
		}
	}

	deletedCategory, err := server.store.DeleteCategory(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, deletedCategory)
}


