package expenses

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		want := Expense{
			Id:     "4",
			Title:  "test",
			Amount: 500,
			Note:   "test",
			Tags:   []string{"test", "test", "test"},
		}

		got := Insert(want.Title, want.Note, want.Amount, want.Tags)

		if got.Title != want.Title || got.Amount != want.Amount || got.Note != want.Note {
			t.Error("want:", want, " got:", got)
		}
	})
}
