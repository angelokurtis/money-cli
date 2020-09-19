package deliveries

import (
	"bytes"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/angelokurtis/money/internal/log"
	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	messages *gmail.UsersMessagesService
}

func NewService(messages *gmail.UsersMessagesService) *Service {
	return &Service{messages: messages}
}

func (s *Service) List(after time.Time) ([]*Request, error) {
	log.Debugf("getting all message from ifood.com.br from %v", after.Format("02/01/2006"))
	res, err := s.messages.
		List("me").
		Q("from:@ifood.com.br subject:(Avalie seu pedido) after:" + after.Format("01/02/2006")).
		Do()
	if err != nil {
		return nil, err
	}

	reqs := make([]*Request, 0, len(res.Messages))
	for _, message := range res.Messages {
		req, err := s.getRequest(message.Id)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, req)
	}
	return reqs, nil
}

func (s *Service) getRequest(id string) (req *Request, err error) {
	msg, err := s.messages.
		Get("me", id).
		Format("full").
		Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get email details")
	}
	body, err := base64.URLEncoding.DecodeString(msg.Payload.Body.Data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode email body")
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read HTML document")
	}
	req = &Request{}
	doc.
		Find("center > table:nth-child(2) > tbody > tr:nth-child(5) > td > table > tbody > tr > th > table > tbody > tr > td").
		Each(func(i int, s *goquery.Selection) {
			request := s.Find("p:nth-child(1)").Text()
			request = strings.Replace(request, "•", "", -1)
			request = strings.Replace(request, "Pedido:", "", -1)
			request = strings.TrimSpace(request)
			req.Id = request

			datetime := s.Find("p:nth-child(2)").Text()
			datetime = strings.Replace(datetime, "•", "", -1)
			datetime = strings.TrimSpace(datetime)
			t, err := time.Parse("02/01/2006  15:04", datetime)
			if err != nil {
				log.Debugf("failed to read time of delivery %v", request)
				return
			}
			req.Datetime = t
		})
	doc.
		Find("center > table:nth-child(2) > tbody > tr:nth-child(6) > td > table > tbody > tr > th > table > tbody > tr > td > p").
		Each(func(i int, s *goquery.Selection) {
			req.Place = s.Text()
		})
	doc.
		Find("center > table:nth-child(2) > tbody > tr:nth-child(11) > td > table > tbody > tr > th > table > tbody > tr > td:nth-child(2) > p").
		Each(func(i int, s *goquery.Selection) {
			value := s.Text()
			value = strings.Replace(value, "R$ ", "", -1)
			value = strings.Replace(value, ".", "", -1)
			value = strings.Replace(value, ",", ".", -1)
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Debugf("failed to read value of delivery %v. %+v", req.Id, err)
				return
			}
			req.Value = f
		})
	return req, nil
}

type Request struct {
	Id       string
	Datetime time.Time
	Place    string
	Value    float64
}
