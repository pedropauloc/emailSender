package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	// 3 A's to test
	// Arrange
	assertTest := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@gmail.com", "email2@gmail.com"}

	// Act
	campaign := NewCampaign(name, content, contacts)

	// Assert
	assertTest.Equal(campaign.ID, "1")
	assertTest.Equal(campaign.Name, name)
	assertTest.Equal(campaign.Content, content)
	assertTest.Equal(len(campaign.Contacts), len(contacts))
}
