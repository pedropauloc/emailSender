package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// 3 A's to test
	// Arrange
	assertTest := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@gmail.com", "email2@gmail.com"}

	// Act
	campaign := NewCampaign(name, content, contacts)

	// Assert
	assertTest.Equal(campaign.Name, name)
	assertTest.Equal(campaign.Content, content)
	assertTest.Equal(len(campaign.Contacts), len(contacts))
}

// Para coisas dinamicas podemos utilizar uma função para testar
func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assertTest := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@gmail.com", "email2@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assertTest.NotNil(campaign.ID)

}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assertTest := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@gmail.com", "email2@gmail.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	// Mesmo se eu não setar a data, ele vai adicionar a data min e sempre vai ser != Nil
	// assertTest.NotNil(campaign.CreatedOn)

	assertTest.Greater(campaign.CreatedOn, now)

}
