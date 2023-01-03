package services

import "github.com/sfreiberg/gotwilio"

type Message struct {
	Content string
	To      string
	Type    string
}

func (m *Message) SendSMS(credentials *gotwilio.Twilio) (*gotwilio.SmsResponse, *gotwilio.Exception, error) {
	resp, excep, err := credentials.SendSMS("+17438002068", m.To, m.Content, "", "")

	if err != nil {
		return nil, excep, err
	}
	return resp, excep, nil
}

func (m *Message) SendWhatsapp(credentials *gotwilio.Twilio) (*gotwilio.SmsResponse, *gotwilio.Exception, error) {
	resp, excep, err := credentials.SendWhatsApp("+17438002068", m.To, m.Content, "", "")

	if err != nil {
		return nil, excep, err
	}
	return resp, excep, nil
}

func SendingMessage(parms Message, twilio *gotwilio.Twilio) (error, *gotwilio.SmsResponse) {
	resp, _, err := parms.SendSMS(twilio)
	if err != nil {
		return err, nil
	}
	return nil, resp
}
