package music

import (
	"fmt"
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}
	m0 := &Music{
		"1", "My Heart Will Go On", "Celion Dion", "Pop",
		"http://qbox.me/24501234", "MP3"}
	Add(m0)
	if Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}
	m := Find(Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}
	if Id != Id || Artist != Artist ||
		Name != Name || Genre != Genre ||
		Source != Source || Type != Type {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}
	m, err := Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}
	fmt.Println(Name)
}
