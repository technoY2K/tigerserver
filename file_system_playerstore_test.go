package tigerserver

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanData := createTempFile(t, `[
		{"Name": "storm", "Wins": 10},
		{"Name": "rogue", "Wins": 30
		}]`)

	defer cleanData()

	store := NewFileSystemPlayerStore(database)

	t.Run("/league from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := []Player{
			{"storm", 10},
			{"rogue", 30},
		}

		assertLeague(t, got, want)
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get a player score from a league", func(t *testing.T) {
		got := store.GetPlayerScore("Rogue")
		want := 30

		if got != want {
			t.Errorf("got %d but wanted %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Storm")
		got := store.GetPlayerScore("Storm")
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Cyclops")
		got := store.GetPlayerScore("Cyclops")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but wanted %d", got, want)
	}
}
