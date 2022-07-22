package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestSetupDB(t *testing.T) {

	author1 := testCreateRandomAuthor(t)
	category1 := createRandomCategory(t)
	book1 := createRandomBook(t)

	require.NotEmpty(t, author1)
	require.NotEmpty(t, category1)
	require.NotEmpty(t, book1)

	
}