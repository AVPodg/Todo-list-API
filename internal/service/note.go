package service

import (
	"context"
	"fmt"
	"rest-notes-api/internal/domain"
)

// NoteRepository - это контракт (интерфейс).
// Сервису всё равно, кто его реализует: реальная база Postgres или Mock-объект в тестах.
type NoteRepository interface {
	Create(ctx context.Context, note domain.Note) (string, error)
	GetAll(ctx context.Context) ([]domain.Note, error)
	FindByID(ctx context.Context, id string) (domain.Note, error)
	Update(ctx context.Context, note domain.Note) error
	Delete(ctx context.Context, id string) error
}

type NoteService struct {
	repo NoteRepository
}

// NewNoteService принимает интерфейс. Теперь сюда можно передать и PostgresRepo, и MockNoteRepo.
func NewNoteService(repo NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) CreateNote(ctx context.Context, title, content string) (string, error) {
	if title == "" {
		return "", fmt.Errorf("title is required")
	}

	note := domain.Note{
		Title:   title,
		Content: content,
	}

	return s.repo.Create(ctx, note)
}

func (s *NoteService) GetAllNotes(ctx context.Context) ([]domain.Note, error) {
	return s.repo.GetAll(ctx)
}

func (s *NoteService) GetNoteByID(ctx context.Context, id string) (domain.Note, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *NoteService) UpdateNote(ctx context.Context, id, title, content string) error {
	if title == "" {
		return fmt.Errorf("title is required")
	}

	// Создаем объект с новыми данными, но старым ID
	note := domain.Note{
		ID:      id,
		Title:   title,
		Content: content,
	}

	return s.repo.Update(ctx, note)
}

func (s *NoteService) DeleteNote(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
