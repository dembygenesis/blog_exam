package mysql

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dembygenesis/blog_exam/pkg/models"
	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArticle_Create(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to init sqlmock: %v", err)
		return
	}

	fmt.Println("mock", mock)
	// ==============================
	// Establish fields and args

	// fields will be your inputs
	type fields struct {
		conn *sql.DB
	}

	// args will be your shit
	type args struct {
		article models.Article
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantID int
		wantErr   bool
	}{
		{
			name: "Test query error",
			fields: fields{
				conn: db,
			},
			args: args{article: models.Article{
				Title:   "Harry Potter",
				Content: "Once upon a time, Alohamora!",
				Author:  "J.K Rowling",
			}},
			wantID: 0,
			wantErr:   true,
		},
		{
			name: "Test unique",
			fields: fields{
				conn: db,
			},
			args: args{article: models.Article{
				Title:   "Exists",
				Content: "Once upon a time, Alohamora!",
				Author:  "J.K Rowling",
			}},
			wantID: 0,
			wantErr:   true,
		},
		{
			name: "Test failed entry",
			fields: fields{
				conn: db,
			},
			args: args{article: models.Article{
				Title:   "Exists",
				Content: "Once upon a time, Alohamora!",
				Author:  "J.K Rowling",
			}},
			wantID: 0,
			wantErr:   true,
		},
		{
			name: "Test Successful Creation",
			fields: fields{
				conn: db,
			},
			args: args{article: models.Article{
				Title:   "Exists",
				Content: "Once upon a time, Alohamora!",
				Author:  "J.K Rowling",
			}},
			wantID: 1,
			wantErr:   true,
		},
	}

	// Run tests
	for _, test := range tests {
		store := Article{conn: test.fields.conn}

		if test.name == "Test query error" {
			mock.ExpectQuery("SELECT (.+) FROM `article` WHERE .*").WillReturnError(errors.New("error trying to validate the article being added"))
			_, err = store.Create(test.args.article)
			assert.EqualError(t, err, "error trying to validate the article being added", "failed at: \"%v\"", test.name)
		}

		if test.name == "Test unique" {
			row := sqlmock.NewRows([]string{"id", "title", "content", "author"}).
				AddRow(1, "Exists", "Once upon a time, Alohamora!", "J.K Rowling")
			mock.ExpectQuery("SELECT (.+) FROM `article` WHERE .*").WillReturnRows(row)
			_, err = store.Create(test.args.article)
			assert.EqualErrorf(t, err, "error: record already exists", "failed at: \"%v\"", test.name)
		}

		if test.name == "Test failed entry" {
			// Pass empty records
			row := sqlmock.NewRows([]string{"id", "title", "content", "author"})
			mock.ExpectQuery("SELECT (.+) FROM `article` WHERE .*").WillReturnRows(row)

			// Fail insert
			mock.ExpectExec("^INSERT INTO `article`.*$").
				WithArgs(test.args.article.Title, test.args.article.Content, test.args.article.Author).
				WillReturnError(errors.New("error trying to insert a new entry"))

			_, err = store.Create(test.args.article)
			assert.EqualErrorf(t, err, "error trying to insert a new entry", "failed at: \"%v\"", test.name)
		}

		if test.name == "Test Successful Creation" {
			// Pass empty records
			row := sqlmock.NewRows([]string{"id", "title", "content", "author"})
			mock.ExpectQuery("SELECT (.+) FROM `article` WHERE .*").WillReturnRows(row)

			// Pass insert
			mock.ExpectExec("^INSERT INTO `article`.*$").
				WithArgs(test.args.article.Title, test.args.article.Content, test.args.article.Author).
				WillReturnResult(sqlmock.NewResult(1, 1))

			res, err := store.Create(test.args.article)
			if err != nil {
				assert.EqualErrorf(t, err, "error trying to insert a new entry", "failed at: \"%v\"", test.name)
			}
			assert.Equalf(t, test.wantID, res.Id, "failed at: \"%v\"", test.name)
		}
	}
}

