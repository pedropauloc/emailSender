package campaign

import (
	"fmt"
	"testing"
)

// é sempre bom testar o ponto de integração
// como por exemplo quando eu cria uma nova campanha
// testamos o comportamento

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@gmail.com", "email2@gmail.com"}

	campaign := NewCampaign(name, content, contacts)
	fmt.Println(campaign)
	
	if campaign.ID != "1" {
		t.Errorf("expected 1")
	} else if campaign.Name != name {
		t.Errorf("expected correct name")
	} else if campaign.Content != content {
		t.Errorf("expected correct content")
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("expected correct contacts")
	}
}
