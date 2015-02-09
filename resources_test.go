package gogame

import "testing"

func setup() {
	//init()
}

func TestAddFontResource(t *testing.T) {

	setup()

	testCases := []struct {
		Request       FontResource
		ErrorExpected bool
	}{
		{FontResource{Id: "font1", FilePath: "./test_resources/fonts/droid-sans/DroidSans.ttf", Size: 16}, false}, // valid request
		{FontResource{Id: "font2", FilePath: "./xxx", Size: 16}, true},                                            // unknown filepath
		{FontResource{Id: "font3", FilePath: "./test_resources/fonts/droid-sans/DroidSans.ttf", Size: -1}, true},  // invalid font size
	}

	for _, test := range testCases {
		err := gameAssets.AddFontResource(test.Request)
		if err == nil && test.ErrorExpected {
			t.Fatal("Error was expected and none received")
		}
		if err != nil && !test.ErrorExpected {
			t.Fatalf("Error was not expected and got: %s", err)
		}

	}
}
