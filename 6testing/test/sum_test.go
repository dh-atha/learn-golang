package test

import (
	"learn1/6testing/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

// jika import "github.com/stretchr/testify/assert" ada garis merahnya, maka jalankan command go get github.com/stretchr/testify/assert
// untuk menjalankan test, bisa langsung klik icon run disamping kiri line function Test, contoh line 13

func TestSumNative(t *testing.T) {
	if function.Sum(2, 3) != 5 {
		t.Error("harus 5")
	}

	// sengaja error
	if function.Sum(2, 3) != 6 {
		t.Error("harus 6")
	}
}

func TestSumUseAssert(t *testing.T) {
	assert.Equal(t, function.Sum(2, 3), 5)
	// sengaja error
	assert.Equal(t, function.Sum(2, 3), 5)
}

var (
	koneksi string
	status  bool
)

func setup() func() {
	koneksi = "konek"
	status = true

	return func() {
		koneksi = ""
		status = false
	}
}

func TestSetup(t *testing.T) {
	// setup berfungsi untuk mengisi variable dan teardown untuk mengosongkan variable
	// statement defer artinya fungsi akan dijalankan terakhir sebelum fungsi selesai
	// bahkan ketika program exit secara paksa a.k.a panic
	// defer bisa lebih dari satu kali, bertumpuk
	teardown := setup()
	defer teardown()

	assert.Equal(t, "konek", koneksi)
	assert.Equal(t, true, status)
}
