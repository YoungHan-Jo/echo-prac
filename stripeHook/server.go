// server.go
//
// Use this sample code to handle webhook events in your integration.
//
// 1) Create a new Go module
//   go mod init example.com/stripe/webhooks/example
//
// 2) Paste this code into a new file (server.go)
//
// 3) Install dependencies
//   go get -u github.com/stripe/stripe-go
//
// 4) Run the server on http://localhost:4242
//   go run server.go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
)

// The library needs to be configured with your account's secret key.
// Ensure the key is kept out of any version control system you might be using.

func main() {
	stripe.Key = "sk_test_51NfeKZJ39MhrdOSIAQdurvWMSNNckcKJqUvYtbyRvh3AVzclc9IFw6eGdbyI8pFfDAfIfw2ivbjjLcSf5PDCxLlt001sSESrDj..."
	http.HandleFunc("/webhook", handleWebhook)
	addr := "localhost:4242"
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleWebhook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("----------webhook----------")
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// This is your Stripe CLI webhook secret for testing your endpoint locally.
	endpointSecret := "whsec_8a9709eea59336748ed1fd282a983ecb5b180355f9786a8bcf15fa20e079849c"
	// Pass the request body and Stripe-Signature header to ConstructEvent, along
	// with the webhook signing key.
	event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"),
		endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		fmt.Println("PaymentIntent was successful!")
		data := event.Data.Object
		fmt.Println("id: ", data["id"])
		fmt.Println("status : ", data["status"])

		// Then define and call a function to handle the event payment_intent.succeeded
	// ... handle other event types
	case "payment_intent.payment_failed":
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		fmt.Println("PaymentIntent was failed!")
		data := event.Data.Object
		fmt.Println("id: ", data["id"])
		fmt.Println("status : ", data["status"])
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}
