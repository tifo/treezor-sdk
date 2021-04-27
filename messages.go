// This file provides functions for validating payloads from Treezor Webhooks.

package treezor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"hash"
	"mime"
	"net/http"

	"github.com/pkg/errors"
)

// genMAC generates the HMAC signature for a message provided the secret key
// and hashFunc.
func genMAC(message, key []byte, hashFunc func() hash.Hash) []byte {
	mac := hmac.New(hashFunc, key)
	_, _ = mac.Write(message)
	return mac.Sum(nil)
}

// checkMAC reports whether messageMAC is a valid HMAC tag for message.
func checkMAC(message, messageMAC, key []byte, hashFunc func() hash.Hash) bool {
	expectedMAC := genMAC(message, key, hashFunc)
	return hmac.Equal(messageMAC, expectedMAC)
}

// messageMAC returns the hex-decoded HMAC tag from the signature and its
// corresponding hash function.
func messageMAC(signature string) ([]byte, func() hash.Hash, error) {
	if signature == "" {
		return nil, nil, errors.New("missing signature")
	}
	buf, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return nil, nil, errors.Errorf("error decoding signature %q: %v", signature, err)
	}
	return buf, sha256.New, nil
}

// ValidatePayload validates an incoming Treezor Webhook event request
// and returns the (JSON) payload.
// The Content-Type header of the payload can only be "application/json".
// If the Content-Type is incorrect then an error is returned.
// secretKey is the Treezor Webhook secret message.
//
// Example usage:
//
//     func (s *TreezorEventMonitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//       payload, err := treezor.ValidatePayload(r, s.webhookSecretKey)
//       if err != nil { ... }
//       // Process payload...
//     }
//
func ValidatePayload(r *http.Request, secretKey []byte) (evt *Event, err error) {
	evt = new(Event)

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return nil, errors.Errorf("Unable to parse Webook request Content-Type %q", r.Header.Get("Content-Type"))
	}
	switch ct {
	case "application/json":
		err = json.NewDecoder(r.Body).Decode(evt)
		if err != nil {
			return nil, errors.New("Webhook request has invalid JSON payload")
		}
	// This case happend for kycliveness event, the content type is not json but it can be parsed as json.
	case "text/plain":
		err = json.NewDecoder(r.Body).Decode(evt)
		if err != nil {
			return nil, errors.New("Webhook request has invalid JSON payload")
		}
	default:
		return evt, errors.Errorf("Webhook request has unsupported Content-Type %q", ct)
	}
	if evt.RawPayload == nil {
		return evt, errors.New("Webhook request has missing payload")
	}
	if evt.PayloadSignature == nil {
		return evt, errors.New("Webhook request has missing signature")
	}
	if err := validateSignature(*evt.PayloadSignature, *evt.RawPayload, secretKey); err != nil {
		return evt, errors.WithStack(err)
	}
	return evt, nil
}

// validateSignature validates the signature for the given payload.
// signature is the Treezor hash signature delivered in the Webhook JSON.
// payload is the JSON payload sent by Treezor Webhooks.
// secretKey is the Treezor Webhook secret message.
func validateSignature(signature string, payload, secretKey []byte) error {
	messageMAC, hashFunc, err := messageMAC(signature)
	if err != nil {
		return errors.WithStack(err)
	}
	if !checkMAC(payload, messageMAC, secretKey, hashFunc) {
		return errors.New("payload signature check failed")
	}
	return nil
}
