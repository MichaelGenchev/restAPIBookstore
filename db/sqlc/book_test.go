package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomBook(t *testing.T) Book {

	arg := CreateBookParams{
		Title:  "kniga",
		Author: 1,
		Category: 1,
		Price: 13.5,
	}

	book, err := testQueries.CreateBook(context.Background(), arg)


	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, arg.Title, book.Title)
	require.Equal(t, arg.Author, book.Author)
	require.Equal(t, arg.Category, book.Category)
	require.Equal(t, arg.Price, book.Price)
	
	require.NotZero(t, book.ID)

	return book
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}


func TestGetBook(t *testing.T) {

	book1 := createRandomBook(t)

	book2, err := testQueries.GetBook(context.Background(), book1.ID)

	
	require.NoError(t, err)
	require.NotEmpty(t, book2)


	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Price, book2.Price)
	require.Equal(t, book1.Category, book2.Category)

}

func TestDeleteBook(t *testing.T){

	book1 := createRandomBook(t)

	book2, err := testQueries.DeleteBook(context.Background(), book1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, book2)


	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Price, book2.Price)
	require.Equal(t, book1.Category, book2.Category)

	book3, err := testQueries.GetBook(context.Background(), book1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book3)

}


func TestUpdateBook(t *testing.T) {

	book1 := createRandomBook(t)

	arg := UpdateBookParams{
		ID: book1.ID,
		Price: 24.99,
	}

	book2, err := testQueries.UpdateBook(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, book2)


	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, arg.Price, book2.Price)
	require.Equal(t, book1.Category, book2.Category)


}


func TestListBooks(t *testing.T) {

	for i := 0; i < 10; i++{
		createRandomBook(t)
	}

	authors , err := testQueries.ListBooks(context.Background())

	require.NoError(t, err)

	for _, author := range authors{

		require.NotEmpty(t, author )
	}


}