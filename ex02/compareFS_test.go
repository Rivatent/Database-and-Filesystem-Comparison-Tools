package main

import (
	"bytes"
	"testing"
)

func TestCompareFS(t *testing.T) {
	t.Run("test compare snapshot for adding new element", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldSnapshot := "test_add_1.txt"
		newSnapshot := "test_add_2.txt"
		CompareSnapshots(buffer, oldSnapshot, newSnapshot)
		got := buffer.String()
		want := "ADDED /etc/systemd/system/very_important/stash_location.jpg\n"

		if got != want {
			t.Errorf("got %s\nwant %s\n", got, want)
		}
	})

	t.Run("test compare snapshot for removing elements", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldSnapshot := "test_remove_1.txt"
		newSnapshot := "test_remove_2.txt"
		CompareSnapshots(buffer, oldSnapshot, newSnapshot)
		got := buffer.String()
		want := "REMOVED /var/log/orders.log\n"

		if got != want {
			t.Errorf("got %s\nwant %s\n", got, want)
		}
	})
}
