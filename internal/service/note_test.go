package service

import (
	"context"
	"rest-notes-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNoteRepo struct {
	mock.Mock
}

func (m *MockNoteRepo) Create(ctx context.Context, note domain.Note) (string, error) {
	args := m.Called(ctx, note)

	return args.String(0), args.Error(1)
}

func (m *MockNoteRepo) GetAll(ctx context.Context) ([]domain.Note, error) { return nil, nil }
func (m *MockNoteRepo) FindByID(ctx context.Context, id string) (domain.Note, error) {
	return domain.Note{}, nil
}
func (m *MockNoteRepo) Update(ctx context.Context, note domain.Note) error { return nil }
func (m *MockNoteRepo) Delete(ctx context.Context, id string) error        { return nil }

func TestNoteService_CreateNote(t *testing.T) {
	tests := []struct {
		name          string
		title         string
		content       string
		mockBehavior  func(m *MockNoteRepo, note domain.Note)
		expectedID    string
		expectedError bool
	}{
		{
			name:    "Success",
			title:   "Test Title",
			content: "Test Content",
			mockBehavior: func(m *MockNoteRepo, note domain.Note) {
				m.On("Create", mock.Anything, mock.MatchedBy(func(n domain.Note) bool {
					return n.Title == "Test Title" && n.Content == "Test Content"
				})).Return("new-uuid", nil)
			},
			expectedID:    "new-uuid",
			expectedError: false,
		},

		{
			name:          "Empty Title",
			title:         "",
			content:       "Content",
			mockBehavior:  func(m *MockNoteRepo, note domain.Note) {},
			expectedID:    "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockNoteRepo)
			tt.mockBehavior(mockRepo, domain.Note{Title: tt.title, Content: tt.content})

			service := NewNoteService(mockRepo)

			id, err := service.CreateNote(context.Background(), tt.title, tt.content)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, id)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
