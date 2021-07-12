package controller

import (
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n.", expected, actual)
	}
}

func TestCtl_ListBooks(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}

func TestCtl_CreateBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}

func TestCtl_GetBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}

func TestCtl_UpdateBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}

func TestCtl_DeleteBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// setup
		// run
		// check
	})
	t.Run("fail", func(t *testing.T) {
		// setup
		// check
		// run
	})
}
