package reprocess

import (
	"context"
	"errors"
	"karhub/internal/entity"
	"testing"
)


type mockRepo struct {
	updateFn func(beer entity.BeerStyle) error
	insertFn func(beer entity.BeerStyle) error
	deleteFn func(style string) error
}

func (m *mockRepo) Update(ctx context.Context, b entity.BeerStyle) (entity.BeerStyle, error) {
	return b, m.updateFn(b)
}
func (m *mockRepo) Insert(ctx context.Context, b entity.BeerStyle) (entity.BeerStyle, error) {
	return b, m.insertFn(b)
}
func (m *mockRepo) Delete(ctx context.Context, s string) error {
	return m.deleteFn(s)
}


func TestReprocess_Reprocess(t *testing.T) {
	ctx := context.Background()
	beer := entity.BeerStyle{Style: "Stout"}

	tests := []struct {
		name      string
		input     entity.Reprocess
		mockSetup func(m *mockRepo, called *bool)
		wantErr   bool
	}{
		{
			name:  "should call Update when the type is update",
			input: entity.Reprocess{QueryType: "update", BeerStyle: beer},
			mockSetup: func(m *mockRepo, called *bool) {
				m.updateFn = func(b entity.BeerStyle) error {
					*called = true
					return nil
				}
			},
			wantErr: false,
		},
		{
			name:  "should return error if the Insert fails",
			input: entity.Reprocess{QueryType: "insert", BeerStyle: beer},
			mockSetup: func(m *mockRepo, called *bool) {
				m.insertFn = func(b entity.BeerStyle) error {
					*called = true
					return errors.New("insert failed")
				}
			},
			wantErr: true,
		},
		{
			name:  "should call Delete when the type is delete",
			input: entity.Reprocess{QueryType: "delete", BeerStyle: beer},
			mockSetup: func(m *mockRepo, called *bool) {
				m.deleteFn = func(s string) error {
					*called = true
					if s != "Stout" {
						return errors.New("wrong style")
					}
					return nil
				}
			},
			wantErr: false,
		},
		{
			name:  "should do nothing if the type is unknown",
			input: entity.Reprocess{QueryType: "unknown"},
			mockSetup: func(m *mockRepo, called *bool) {
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			repo := &mockRepo{}
			tt.mockSetup(repo, &called)

			uc := NewReprocess(repo)
			err := uc.Reprocess(ctx, tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Reprocess() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.input.QueryType != "unknown" && !called && !tt.wantErr {
				t.Errorf("the expected repository function was not called")
			}
		})
	}
}
