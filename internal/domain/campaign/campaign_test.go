package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Campaign Content"
	contacts = []string{"email_1@acme.com", "email_2@acme.com", "email_3@acme.com"}
)

func Test_NewCampaign_Create(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := MakeCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_CampaignId_IsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := MakeCampaign(name, content, contacts)

	assert.NotNil(campaign.Id)
}

func Test_CampaignCreatedAt_IsNotNill(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := MakeCampaign(name, content, contacts)

	assert.NotNil(campaign.CreatedAt)
	assert.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_MustBeValidatedName(t *testing.T) {
	assert := assert.New(t)

	_, error := MakeCampaign("", content, contacts)

	assert.NotNil(error.Error())
	assert.Equal("name is required", error.Error())
}

func Test_NewCampaign_MustBeValidatedContent(t *testing.T) {
	assert := assert.New(t)

	_, error := MakeCampaign(name, "", contacts)

	assert.NotNil(error.Error())
	assert.Equal("content is required", error.Error())
}

func Test_NewCampaign_MustBeValidatedContacts(t *testing.T) {
	assert := assert.New(t)

	_, error := MakeCampaign(name, content, []string{})

	assert.NotNil(error.Error())
	assert.Equal("emails list is empty", error.Error())
}

func Test_NewCampaign_MustBeValidatedEmails(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := MakeCampaign(name, content, contacts)

	for index, contact := range campaign.Contacts {
		assert.Equal(contact.Email, contacts[index])
	}
}
