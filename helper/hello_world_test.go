package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTableHelloWold(b *testing.B) {
	bencmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "Bima",
			Request: "Bima",
		},
		{
			Name:    "komang",
			Request: "Komang",
		},
	}
	for _, bench := range bencmarks {
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.Request)
			}
		})
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Bima", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bima")
		}
	})
	b.Run("Komang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Komang")
		}
	})
}

func BenchmarkHelloWorldBima(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bima")
	}
}

func BenchmarkHelloWorldKomang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Komang")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Bima",
			request:  "Bima",
			expected: "Hello Bima",
		},
		{
			name:     "Komang",
			request:  "Komang",
			expected: "Hello Komang",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Bima", func(t *testing.T) {
		result := HelloWorld("Bima")
		//memanggil t.Fail() jika error
		require.Equal(t, "Hello Bima", result, "result must be Hello Bima")

	})

	t.Run("Komang", func(t *testing.T) {
		result := HelloWorld("Komang")
		//memanggil t.Fail() jika error
		require.Equal(t, "Hello Komang", result, "result must be Hello Komang")

	})
}

func TestMain(m *testing.M) {
	//before

	fmt.Println("Before unit test")

	m.Run()

	//after

	fmt.Println("after unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cant run in windows")
	}

	result := HelloWorld("Bima")
	//memanggil t.Fail() jika error
	assert.Equal(t, "Hello Bima", result, "result must be Hello Bima")

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bima")
	//memanggil t.Fail() jika error
	assert.Equal(t, "Hello Bima", result, "result must be Hello Bima")

	fmt.Println("Done tes")

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bima")
	//memanggil t.FailNow() jika error
	require.Equal(t, "Hello Bima", result, "result must be Hello Bima")

	fmt.Println("Done tes")

}

func TestHelloWorldBima(t *testing.T) {

	result := HelloWorld("Bima")

	if result != "Hello Bima" {
		// menjalankan kode program berikutnya walaupun fail
		// t.Fail()

		// menjalankan kode program berikutnya walaupun fail dengan message
		t.Error("result must be Hello Bima")
	}

	fmt.Println("Done tes")
}

func TestHelloWorldKomang(t *testing.T) {

	result := HelloWorld("Komang")

	if result != "Hello Komang" {
		// berhenti disini jika fail
		// t.FailNow()

		// berhenti disini jika fail dengan message
		t.Fatal("result must be Hello Komang")
	}
	fmt.Println("Done tes")

}
