package sendcloud_email_go

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

const (
	sendCommonPath   = "/apiv2/mail/send"
	sendTemplatePath = "/apiv2/mail/sendtemplate"
)

func (client *SendCloud) SendCommonEmail(ctx context.Context, args *CommonMail) (*SendEmailResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}
	if err := args.validateSendCommonEmail(); err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}
	var req *http.Request
	var err error
	sendCommonUrl := client.apiBase + sendCommonPath
	if args.body.Attachments == nil {
		params:= client.PrepareSendCommonEmailParams(args)
		formDataEncoded := params.Encode()
		req, err = http.NewRequest("POST", sendCommonUrl, bytes.NewBufferString(formDataEncoded))
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		multipartWriter,payload, err := client.MultipartSendCommonMail(args)
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req, err = http.NewRequest("POST", sendCommonUrl, payload)
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	}
	responseData := new(SendEmailResult)
	err = client.Request(ctx, req, &responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (client *SendCloud) SendEmailTemplate(ctx context.Context, args *TemplateMail) (*SendEmailResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}
	if err := args.validateSendEmailTemplate(); err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}
	var req *http.Request
	var err error
	sendTemplateUrl := client.apiBase+sendTemplatePath
	if args.body.Attachments == nil {
		params:= client.PrepareSendEmailTemplateParams(args)
		formDataEncoded := params.Encode()
		req, err = http.NewRequest("POST", sendTemplateUrl, bytes.NewBufferString(formDataEncoded))
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		multipartWriter,payload, err := client.MultipartSendEmailTemplate(args)
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req, err = http.NewRequest("POST", sendTemplateUrl, payload)
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	}
	responseData := new(SendEmailResult)
	err = client.Request(ctx, req, &responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
