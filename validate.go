package sendcloud_email_go

import (
	"errors"
	"strings"
)

const MAX_RECEIVERS = 100
const MAX_MAILLIST = 5
const MAX_CONTENT_SIZE = 10 * 1024 * 1024

func (e *TemplateMail)validateSendEmailTemplate() error {
	if err := e.To.validateReceiver(); err != nil {
		return err
	}
	if err := e.body.validateMailBody(); err != nil {
		return err
	}
	if e.content.TemplateInvokeName == "" {
		return errors.New("templateInvokeName cannot be empty")
	}
	return nil
}

func (e *CommonMail)validateSendCommonEmail() error {
	if err := e.To.validateReceiver(); err != nil {
		return err
	}
	if err := e.body.validateMailBody(); err != nil {
		return err
	}
	switch {
	case len(e.content.Html) == 0 && len(e.content.Plain) == 0:
		return errors.New("html or plain cannot be empty")
	case len(e.content.Html) > 0 && len(e.content.Html) > MAX_CONTENT_SIZE:
		return errors.New("html content is too long")
	case len(e.content.Plain) > 0 && len(e.content.Plain) > MAX_CONTENT_SIZE:
		return errors.New("plain Content is too long")
	}
	return nil
}

func (e *Receiver)validateReceiver() error {
	if len(e.To) == 0 {
		return errors.New("to cannot be empty")
	}
	if e.UseAddressList {
		to := strings.Split(e.To, ";")
		if len(to) > MAX_MAILLIST {
			return errors.New("address list exceeds limit")
		}
	}else {
		to := strings.Split(e.To, ";")
		cc := strings.Split(e.CC, ";")
		bcc := strings.Split(e.BCC, ";")
		receivers := len(to)
		receivers += len(cc)
		receivers += len(bcc)
		// Check if the total number of receivers exceeds the maximum allowed
		if receivers > MAX_RECEIVERS {
			return errors.New("the total number of receivers exceeds the maximum allowed")
		}
	}
	return nil
}

func (e *MailBody)validateMailBody() error {
	switch {
	case len(e.From) == 0:
		return errors.New("from cannot be empty")
	case len(e.Subject) == 0:
		return errors.New("subject cannot be empty")
	}
	return nil
}