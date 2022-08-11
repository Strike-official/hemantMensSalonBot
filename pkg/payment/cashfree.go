package payment

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Strike-official/hemantMensSalonBot/internal/model"
	"github.com/google/uuid"
)

const (
	linkCurrency string = "INR"
)

type Payload struct {
	CustomerDetails CustomerDetails `json:"customer_details"`
	LinkNotify      LinkNotify      `json:"link_notify"`
	LinkID          string          `json:"link_id"`
	LinkAmount      int             `json:"link_amount"`
	LinkCurrency    string          `json:"link_currency"`
	LinkPurpose     string          `json:"link_purpose"`
	LinkExpiryTime  time.Time       `json:"link_expiry_time"`
	LinkMeta        LinkMeta        `json:"link_meta"`
}
type CustomerDetails struct {
	CustomerPhone string `json:"customer_phone"`
	CustomerName  string `json:"customer_name"`
}
type LinkNotify struct {
	SendSms bool `json:"send_sms"`
}
type LinkMeta struct {
	UPIIntent bool `json:"upi_intent"`
}

type LinkResponse struct {
	LinkURL string
	LinkID  string
}

type PayeeData struct {
	PayeePrefix      string
	PayeeDisplayName string
	PaymentAmount    int
}

func RequestPaymentLink(request model.Request_Structure, payeeData PayeeData) LinkResponse {

	var result LinkResponse
	id := uuid.New()
	paymentLinkId := payeeData.PayeePrefix + "_" + id.String()
	paymentLinkId = paymentLinkId[:len(payeeData.PayeePrefix)+8]
	log.Println("paymentLinkId: ", paymentLinkId)
	isoTimePlus10String := time.Now().Add(time.Minute * 10).Format(time.RFC3339)
	isoTimePlus10, err := time.Parse(time.RFC3339, isoTimePlus10String)
	if err != nil {
		log.Println("[ERROR] Error in converting iso string to time, err: ", err)
		result := LinkResponse{
			LinkURL: "PAYMENT FAILED",
			LinkID:  paymentLinkId,
		}
		return result
	}

	data := Payload{
		CustomerDetails: CustomerDetails{
			CustomerPhone: request.Bybrisk_session_variables.Phone,
			CustomerName:  request.Bybrisk_session_variables.Username,
		},
		LinkNotify: LinkNotify{
			SendSms: true,
		},
		LinkID:         paymentLinkId,
		LinkAmount:     payeeData.PaymentAmount,
		LinkCurrency:   linkCurrency,
		LinkPurpose:    payeeData.PayeeDisplayName,
		LinkExpiryTime: isoTimePlus10,
		LinkMeta: LinkMeta{
			UPIIntent: true,
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("[ERROR] Error in marshling payload data, err: ", err)
		result = LinkResponse{
			LinkURL: "PAYMENT FAILED",
			LinkID:  data.LinkID,
		}
		return result
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", model.Conf.PaymentLinkURL, body)
	if err != nil {
		log.Println("[ERROR] Error in forming request to cashfree, err: ", err)
		result = LinkResponse{
			LinkURL: "PAYMENT FAILED",
			LinkID:  data.LinkID,
		}
		return result
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Version", model.Conf.XApiVersion)
	req.Header.Set("X-Client-Id", model.Conf.XClientId)
	req.Header.Set("X-Client-Secret", model.Conf.XClientSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[ERROR] Error in sending payment link request to cashfree, err: ", err)
		result = LinkResponse{
			LinkURL: "PAYMENT FAILED",
			LinkID:  data.LinkID,
		}
		return result
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	log.Println("[Response from cashfree] res: ", res)
	if res["link_url"] != nil {
		result = LinkResponse{
			LinkURL: res["link_url"].(string),
			LinkID:  data.LinkID,
		}
	} else {
		result = LinkResponse{
			LinkURL: "PAYMENT FAILED",
			LinkID:  data.LinkID,
		}
	}

	defer resp.Body.Close()

	return result
}

func GetPaymentStatus(linkID string) string {

	req, err := http.NewRequest("GET", model.Conf.PaymentLinkURL+"/"+linkID, nil)
	if err != nil {
		log.Println("[ERROR] Error in forming request to cashfree, err: ", err)
		return "ERROR"
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Api-Version", model.Conf.XApiVersion)
	req.Header.Set("X-Client-Id", model.Conf.XClientId)
	req.Header.Set("X-Client-Secret", model.Conf.XClientSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[ERROR] Error in getting payment status, err: ", err)
		return "ERROR"
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	defer resp.Body.Close()

	return res["link_status"].(string)
}
