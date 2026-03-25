package repo

import (
	"context"
	"rest-notes-api/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const DSN = "postgres://user:password@localhost:5432/notes_db?sslmode=disable"

func TestPostgresRepo_Integration(t *testing.T) {
	repo, err := NewPostgresRepo(DSN)
	if err != nil {
		t.Skip("Skipping integration test: database not available")
	}

	defer repo.Close()

	ctx := context.Background()

	note := domain.Note{
		Title:   "Integration Test",
		Content: "Works with DB",
	}
	id, err := repo.Create(ctx, note)
	// Create
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	//FindByID

	savedNote, err := repo.FindByID(ctx, note.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Integration Test", savedNote.Title)
	assert.Equal(t, "Works with DB", savedNote.Content)
	assert.WithinDuration(t, time.Now(), savedNote.CreateAt, 2*time.Second)

	//Update

	savedNote.Title = "Update title"
	err = repo.Update(ctx, savedNote)
	assert.NoError(t, err)

	updateNote, err := repo.FindByID(ctx, savedNote.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Update Test", updateNote.Title)

	//Delete

	err = repo.Delete(ctx, id)
	assert.NoError(t, err)

	_, err = repo.FindByID(ctx, id)
	assert.Error(t, err)
}
