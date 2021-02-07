package repository

import (
	"errors"
	"testing"

	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestTodoItemPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewTodoItemPostgres(db)

	type args struct {
		listID int
		item   todo.TodoItem
	}

	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO todo_items").WithArgs(
					args.item.Title, args.item.Description,
				).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO lists_items").WithArgs(
					args.listID, id,
				).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			input: args{
				listID: 1,
				item: todo.TodoItem{
					Title:       "test title",
					Description: "test description",
				},
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO todo_items").WithArgs(
					args.item.Title, args.item.Description,
				).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			input: args{
				listID: 1,
				item: todo.TodoItem{
					Title:       "",
					Description: "test description",
				},
			},
			wantErr: true,
		},
		{
			name: "Failed 2nd Insert",
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO todo_items").WithArgs(
					args.item.Title, args.item.Description,
				).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO lists_items").WithArgs(
					args.listID, id,
				).WillReturnError(errors.New("insert error"))

				mock.ExpectRollback()
			},
			input: args{
				listID: 1,
				item: todo.TodoItem{
					Title:       "test title",
					Description: "test description",
				},
			},
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input, testCase.want)

			got, err := r.Create(testCase.input.listID, testCase.input.item)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
		})
	}
}
