package gogame

import "testing"

func TestNewGame(t *testing.T) {

	testCases := []struct {
		Title          string
		Width          int
		Height         int
		RenderCallback RenderFunction
		EventHandler   EventReceiverFunction
		ErrorExpected  bool
	}{
		{"", 100, 100, nil, nil, true},        // no title
		{"Title", 0, 100, nil, nil, true},     // invalid width
		{"Title", 100, 0, nil, nil, true},     // invalid height
		{"Title", 100, 1000, nil, nil, false}, // valid
	}

	for _, test := range testCases {
		_, err := NewGame(test.Title, test.Width, test.Height, test.RenderCallback, test.EventHandler)
		if err == nil && test.ErrorExpected {
			t.Fatal("Error was expected and none received")
		}

		if err != nil && !test.ErrorExpected {
			t.Fatalf("Error was not expected and got: %s", err)
		}

	}
}
