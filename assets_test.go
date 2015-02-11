package gogame

import "testing"

func setup() {
	//init()
}

func TestAddFontAsset(t *testing.T) {

	setup()

	testCases := []struct {
		Request       FontAsset
		ErrorExpected bool
	}{
		{FontAsset{Id: "font1", FilePath: "./test_assets/fonts/droid-sans/DroidSans.ttf", Size: 16}, false}, // valid request
		{FontAsset{Id: "font2", FilePath: "./xxx", Size: 16}, true},                                            // unknown filepath
		{FontAsset{Id: "font3", FilePath: "./test_assets/fonts/droid-sans/DroidSans.ttf", Size: -1}, true},  // invalid font size
	}

	for _, test := range testCases {
		err := gameAssets.AddFontAsset(test.Request)
		if err == nil && test.ErrorExpected {
			t.Fatal("Error was expected and none received")
		}
		if err != nil && !test.ErrorExpected {
			t.Fatalf("Error was not expected and got: %s", err)
		}

	}
}
