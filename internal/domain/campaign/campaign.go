package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	Id       string
	Name     string
	Content  string
	Contacts []Contact

	CreatedAt time.Time
}

func MakeCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("emails list is empty")
	}

	contacs := make([]Contact, len(emails))

	for index, email := range emails {
		contacs[index].Email = email
	}

	return &Campaign{
		Id:       xid.New().String(),
		Name:     name,
		Contacts: contacs,
		Content:  content,

		CreatedAt: time.Now(),
	}, nil
}
