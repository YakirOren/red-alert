package redAlert

import (
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"polling-to-ws/broadcast"
	"time"
)

type RedAlertClient struct {
	hub           *broadcast.Hub
	previousAlert []byte
}

func NewRedAlertClient(hub *broadcast.Hub) RedAlertClient {
	return RedAlertClient{hub: hub}
}

type Response struct {
	Id    string   `json:"id"`
	Cat   string   `json:"cat"`
	Title string   `json:"title"`
	Data  []string `json:"data"`
	Desc  string   `json:"desc"`
}

func (c *RedAlertClient) Run() {
	for {
		time.Sleep(1 * time.Second)
		res, err := FetchAlerts()
		if err != nil {
			log.Warn(err)
			continue
		}

		data, err := c.ParseResponse(res.Body)
		res.Body.Close()
		if err != nil {
			log.Debug(err)
			continue
		}

		c.hub.Broadcast <- data
		log.Info("Sent")

	}

}

func (c *RedAlertClient) ParseResponse(data io.Reader) ([]byte, error) {
	var redAlertResponse Response

	if err := json.NewDecoder(data).Decode(&redAlertResponse); err != nil {
		return nil, errors.New("invalid json")
	}

	marshaledData, err := json.Marshal(redAlertResponse)
	if err != nil {
		return nil, err
	}

	if bytes.Equal(marshaledData, c.previousAlert) {
		log.Debugf("Data equal %v", marshaledData)
		return nil, err
	}

	c.previousAlert = marshaledData
	return marshaledData, nil
}

func FetchAlerts() (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, "https://www.oref.org.il/WarningMessages/alert/alerts.json", nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Referer", "https://www.oref.org.il/11088-13708-he/Pakar.aspx")

	res, err := http.DefaultClient.Do(req)
	return res, err
}
