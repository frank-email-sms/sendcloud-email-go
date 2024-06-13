package sendcloud_email_go

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestSendCommonEmail(t *testing.T) {
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &SendEmailArgs{
		EmailCommonFields: EmailCommonFields{
			To: "a@ifaxin.com;b@ifaxin.com",
            From: "SendCloud@SendCloud.com",
            Subject: "Email from SendCloud SDK",
            FromName: "SendCloud",
		},
		Html:  "<p>This is an HTML email.</p>",
	}
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendCommonEmailWithAttachment(t *testing.T) {
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	attachment1, err := os.Open("path/to/attachment1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer attachment1.Close()
	args :=  &SendEmailArgs{
		EmailCommonFields: EmailCommonFields{
			To: "a@ifaxin.com;b@ifaxin.com",
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
			Attachments: []*os.File{
				attachment1,
            },
		},
		Html:  "<p>This is an HTML email.</p>",
	}
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Message)
}