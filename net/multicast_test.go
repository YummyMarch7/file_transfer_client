package net

import "testing"

func TestMulticast_sendmsg(t *testing.T) {
	payload := []byte("hello")
	sendMessage(payload)
}
