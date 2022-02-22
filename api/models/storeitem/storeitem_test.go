package storeitem

import (
	"testing"
)

func TestAdd(t *testing.T) {
	_storeitem := New()
	_storeitem.Add(Item{"test", 0})
	if len(_storeitem.Items) != 1 {
		t.Errorf("There was no Item added!")
	}
}

func TestGetItems(t *testing.T) {
	_storeitem := New()
	_storeitem.Add(Item{"test", 0})
	_all_items := _storeitem.GetItems()
	if len(_all_items) != 1 {
		t.Errorf("There was no Item added!")
	}
}