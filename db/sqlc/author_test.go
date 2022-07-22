package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/MichaelGenchev/restAPIBookstore/util"
	"github.com/stretchr/testify/require"
)


func testCreateRandomAuthor(t *testing.T) Author{
	arg := CreateAuthorParams{
		Name: util.RandomName(),
		Biography: util.RandomBio(),
	}

	author, err := testQueries.CreateAuthor(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.Name, author.Name)
	require.Equal(t, arg.Biography, author.Biography)

	require.NotZero(t, author.ID)

	return author

}

func TestCreateAuthor(t *testing.T) {
	testCreateRandomAuthor(t)

}

func TestGetAuthor(t *testing.T){

	author1 := testCreateRandomAuthor(t)

	author2, err := testQueries.GetAuthor(context.Background(), author1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author1.ID, author2.ID)
	require.Equal(t, author1.Name, author2.Name)
	require.Equal(t, author1.Biography, author2.Biography)


}

func TestUpdateAuthor(t *testing.T) {

	author1 := testCreateRandomAuthor(t)

	arg :=  UpdateAuthorParams{
		ID: author1.ID,
		Name: "updated name",
		Biography: "updated bio",
	}
	
	author2, err := testQueries.UpdateAuthor(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, author2)


	require.Equal(t, author1.ID, author2.ID)
	require.Equal(t, arg.Name, author2.Name)
	require.Equal(t, arg.Biography, author2.Biography)
}


func TestDeleteAuthor(t *testing.T) {

	author1 := testCreateRandomAuthor(t)

	author2, err := testQueries.DeleteAuthor(context.Background(), author1.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author1.Name, author2.Name)
	require.Equal(t, author1.Biography, author2.Biography)

	author3 , err := testQueries.GetAuthor(context.Background(), author1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, author3)

}


func TestListAuthors(t *testing.T) {

	for i := 0; i < 10; i++{
		testCreateRandomAuthor(t)
	}

	authors , err := testQueries.ListAuthors(context.Background())

	require.NoError(t, err)

	for _, author := range authors{

		require.NotEmpty(t, author )
	}

}