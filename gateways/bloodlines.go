package gateways

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/models"
)

type Bloodlines interface {
}

type bloodlines struct {
	host   string
	port   string
	url    string
	client *http.Client
}

func NewBloodlines(config config.Bloodlines) Bloodlines {
	return &bloodlines{
		host:   config.Host,
		port:   config.Port,
		url:    fmt.Sprintf("https://%s:%s/api/", config.Host, config.Port),
		client: &http.Client{},
	}
}

func (b *bloodlines) GetAllContent(offset int, limit int) ([]*models.Content, error) {
	url := fmt.Sprintf("%scontent?offset=%d&limit=%d", b.url, offset, limit)
	raw, err := ServiceGet(b.client, url)
	if err != nil {
		return nil, err
	}

	var content []*models.Content
	err = json.Unmarshal(raw, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (b *bloodlines) GetContentByID(id uuid.UUID) (*models.Content, error) {
	url := fmt.Sprintf("%scontent/%s", b.url, id.String())
	raw, err := ServiceGet(b.client, url)
	if err != nil {
		return nil, err
	}

	var content *models.Content
	err = json.Unmarshal(raw, content)
	if err != nil {
		return nil, err
	}

	return content, nil
}
