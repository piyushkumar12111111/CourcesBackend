package handlers

import (
    "encoding/json"
    "net/http"
     
	"github.com/piyushkumar/authentication/authentication/mail"
)

// MailRequest holds the data for a mail request
type MailRequest struct {
    To      []string `json:"to"`
    Subject string   `json:"subject"`
    Body    string   `json:"body"`
}

// SendMailHandler handles requests to send mail
func SendMailHandler(w http.ResponseWriter, r *http.Request) {
    var req MailRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := mail.SendMail(req.To, req.Subject, req.Body); err != nil {
        http.Error(w, "Failed to send mail", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
