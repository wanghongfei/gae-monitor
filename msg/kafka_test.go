package msg

import "testing"

func TestSendMessage(t *testing.T) {
	SendMessage("dev-gae-charge", "demo")
}
