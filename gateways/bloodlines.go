package gateways

import (
	"fmt"
	"net/http"

	"github.com/pborman/uuid"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/models"
)

/*Bloodlines wraps all the methods of the bloodlines API */
type Bloodlines interface {
	GetAllContent(offset int, limit int) ([]*models.Content, error)
	NewContent(newContent *models.Content) (*models.Content, error)
	GetContentByID(id uuid.UUID) (*models.Content, error)
	UpdateContent(update *models.Content) (*models.Content, error)
	DeleteContent(id uuid.UUID) error
	GetAllReceipts(offset int, limit int) ([]*models.Receipt, error)
	SendReceipt(receipt *models.Receipt) (*models.Receipt, error)
	GetReceiptByID(id uuid.UUID) (*models.Receipt, error)
	GetAllTriggers(offset int, limit int) ([]*models.Trigger, error)
	NewTrigger(t *models.Trigger) (*models.Trigger, error)
	GetTriggerByKey(key string) (*models.Trigger, error)
	UpdateTrigger(update *models.Trigger) (*models.Trigger, error)
	DeleteTrigger(key string) error
	ActivateTrigger(key string, receipt *models.Receipt) (*models.SendRequest, error)
}

type bloodlines struct {
	*BaseService
	url    string
	client *http.Client
}

/*NewBloodlines creates and returns a Bloodlines struct pointed at the service denoted in config*/
func NewBloodlines(config config.Bloodlines) Bloodlines {
	return &bloodlines{
		BaseService: NewBaseService(),
		url:         fmt.Sprintf("https://%s:%s/api/", config.Host, config.Port),
	}
}

func (b *bloodlines) GetAllContent(offset int, limit int) ([]*models.Content, error) {
	url := fmt.Sprintf("%scontent?offset=%d&limit=%d", b.url, offset, limit)

	content := make([]*models.Content, 0)
	err := b.ServiceSend(http.MethodGet, url, nil, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (b *bloodlines) NewContent(newContent *models.Content) (*models.Content, error) {
	url := fmt.Sprintf("%scontent", b.url)

	var content *models.Content
	err := b.ServiceSend(http.MethodPost, url, newContent, content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (b *bloodlines) GetContentByID(id uuid.UUID) (*models.Content, error) {
	url := fmt.Sprintf("%scontent/%s", b.url, id.String())

	var content *models.Content
	err := b.ServiceSend(http.MethodGet, url, nil, content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (b *bloodlines) UpdateContent(update *models.Content) (*models.Content, error) {
	url := fmt.Sprintf("%scontent/%s", b.url, update.ID.String())

	var content *models.Content
	err := b.ServiceSend(http.MethodPut, url, update, content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (b *bloodlines) DeleteContent(id uuid.UUID) error {
	url := fmt.Sprintf("%scontent/%s", b.url, id.String())

	err := b.ServiceSend(http.MethodDelete, url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (b *bloodlines) GetAllReceipts(offset int, limit int) ([]*models.Receipt, error) {
	url := fmt.Sprintf("%sreceipt?offset=%d&limit=%d", b.url, offset, limit)

	var receipts []*models.Receipt
	err := b.ServiceSend(http.MethodGet, url, nil, &receipts)
	if err != nil {
		return nil, err
	}

	return receipts, nil
}

func (b *bloodlines) SendReceipt(receipt *models.Receipt) (*models.Receipt, error) {
	url := fmt.Sprintf("%sreceipt/send", b.url)

	var sent *models.Receipt
	err := b.ServiceSend(http.MethodPost, url, receipt, sent)
	if err != nil {
		return nil, err
	}

	return sent, nil
}

func (b *bloodlines) GetReceiptByID(id uuid.UUID) (*models.Receipt, error) {
	url := fmt.Sprintf("%sreceipt/%s", b.url, id.String())

	var receipt *models.Receipt
	err := b.ServiceSend(http.MethodGet, url, nil, receipt)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (b *bloodlines) GetAllTriggers(offset int, limit int) ([]*models.Trigger, error) {
	url := fmt.Sprintf("%strigger?offset=%d&limit=%d", b.url, offset, limit)

	var triggers []*models.Trigger
	err := b.ServiceSend(http.MethodGet, url, nil, &triggers)
	if err != nil {
		return nil, err
	}

	return triggers, nil
}

func (b *bloodlines) NewTrigger(t *models.Trigger) (*models.Trigger, error) {
	url := fmt.Sprintf("%strigger", b.url)

	var trigger *models.Trigger
	err := b.ServiceSend(http.MethodPost, url, t, trigger)
	if err != nil {
		return nil, err
	}

	return trigger, nil
}

func (b *bloodlines) GetTriggerByKey(key string) (*models.Trigger, error) {
	url := fmt.Sprintf("%strigger/%s", b.url, key)

	var trigger *models.Trigger
	err := b.ServiceSend(http.MethodGet, url, nil, trigger)
	if err != nil {
		return nil, err
	}

	return trigger, nil
}

func (b *bloodlines) UpdateTrigger(update *models.Trigger) (*models.Trigger, error) {
	url := fmt.Sprintf("%strigger/%s", b.url, update.Key)

	var trigger *models.Trigger
	err := b.ServiceSend(http.MethodPut, url, update, trigger)
	if err != nil {
		return nil, err
	}

	return trigger, nil
}

func (b *bloodlines) DeleteTrigger(key string) error {
	url := fmt.Sprintf("%strigger/%s", b.url, key)

	err := b.ServiceSend(http.MethodDelete, url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (b *bloodlines) ActivateTrigger(key string, receipt *models.Receipt) (*models.SendRequest, error) {
	url := fmt.Sprintf("%strigger/%s/activate", b.url, key)

	var request *models.SendRequest
	err := b.ServiceSend(http.MethodPost, url, receipt, request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
