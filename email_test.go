package sendcloud_email_go

import (
	"context"
	"testing"
	"time"
)

func TestSendCommonEmail(t *testing.T) {
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &SendEmailArgs{
		EmailCommonFields: EmailCommonFields{
			To: "a@ifaxin.com",
            From: "SendCloud@SendCloud.com",
            Subject: "Email from SendCloud SDK",
            FromName: "SendCloud",
		},
		Html:  "<p>This is an HTML email.</p>",
	}
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}