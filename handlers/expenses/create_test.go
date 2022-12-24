package expenses_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
)

type Expense struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

var url = "http://localhost" + os.Getenv("PORT")

func TestCreate(t *testing.T) {
	t.Run("create should return statusCreated", func(t *testing.T) {

		body := bytes.NewBufferString(`{
			"title": "buy a new phone",
			"amount": 39000,
			"note": "buy a new phone",
			"tags": ["gadget", "shopping"]
		}`)

		var expenses Expense

		res := request(http.MethodPost, url+"/expenses", body)
		res.Decode(&expenses)

		if res.StatusCode != http.StatusCreated {
			t.Errorf("StatusCode should be: %d, but %d", http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("create should return statusBadRequest", func(t *testing.T) {

		body := bytes.NewBufferString(`{
			"title": "buy a new phone",
			"amount": 0,
			"note": "buy a new phone",
			"tags": ["gadget", "shopping"]
		}`)

		var expenses Expense

		res := request(http.MethodPost, url+"/expenses", body)
		res.Decode(&expenses)

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("StatusCode should be: %d, but %d", http.StatusBadRequest, res.StatusCode)
		}
	})
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v) // อ่านของจาก body แล้วนำไปยัดใส่ struct ที่เราส่งมา
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}
