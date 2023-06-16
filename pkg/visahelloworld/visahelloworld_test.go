package visahelloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Log("Hello World")

	err := HelloWorldRequest()
	if err != nil {
		t.Error("Error in HelloWorldRequest: ", err)
	}
}
