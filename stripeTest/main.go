package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/card"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/paymentmethod"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
	"github.com/stripe/stripe-go/v74/subscription"
)

func main() {
	stripe.Key = "sk_test_51NfeKZJ39MhrdOSIAQdurvWMSNNckcKJqUvYtbyRvh3AVzclc9IFw6eGdbyI8pFfDAfIfw2ivbjjLcSf5PDCxLlt001sSESrDj"

	err := createCard()
	if err != nil {
		fmt.Println(err)
	}
}

func createCard() error {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_OSwy42WPhIJtir"),
		// Token:    stripe.String("tok_jcb"),
		Token: stripe.String("tok_1NhTmEJ39MhrdOSIDukixKt4"),
	}
	c, err := card.New(params)

	fmt.Println(c.ID)
	return err
}

func createSession() error {
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_1NfzmjJ39MhrdOSI1lK8jKJ3"),
				Quantity: stripe.Int64(3),
			},
		},
		Mode:       stripe.String("payment"),                       // payment, setup, subscription
		SuccessURL: stripe.String("http://localhost:3333/success"), // 決済成功した後、戻ってくるUrl
	}
	s, err := session.New(params)

	fmt.Println("session id : ", s.ID)
	fmt.Println("cancel url : ", s.CancelURL)
	fmt.Println("pi : ", s.PaymentIntent)
	fmt.Println("url : ", s.URL)
	fmt.Println("success url: ", s.SuccessURL)

	return err
}

func confirmPaymentIntent() error {
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_1NhSjZJ39MhrdOSIR62u98H1"),
	}
	pi, err := paymentintent.Confirm(
		"pi_3NhSCKJ39MhrdOSI1OOwc8sr",
		params,
	)

	fmt.Println("pi status: ", pi.Status)

	return err
}

func createPaymentIntent() error {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(40000),
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		Customer: stripe.String("cus_OSbCLY7xLaBtTK"), // test customer1 ( cus_OSbCLY7xLaBtTK )
	}
	pi, err := paymentintent.New(params)

	fmt.Println(pi.ID)

	return err
}

func createPaymentMethod() error {
	params := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String("tok_1NhSioJ39MhrdOSI3uohUpVp"),
		},
		Type: stripe.String("card"),
	}
	pm, err := paymentmethod.New(params)

	fmt.Println(pm.ID)

	return err
}

func getPaymentIntent() error {
	pi, err := paymentintent.Get(
		"pi_3NhO4GJ39MhrdOSI1BxQ5YxD",
		nil,
	)

	fmt.Println(pi.Status)

	return err
}

func getSession() error {

	s, err := session.Get(
		"cs_test_a1Ybob68FTfhBeS7gbxb64HyNuiSVhkQGLecYOFRQBmzOiuNeKugvZtD4G",
		nil,
	)
	fmt.Println("session id : ", s.ID)
	fmt.Println("cancel url : ", s.CancelURL)
	fmt.Println("status: ", s.Status)
	fmt.Println("pi_id : ", s.PaymentIntent)
	// fmt.Println("success url: ", s.SuccessURL)
	// fmt.Println(s.LineItems.Data[0].Price.ID)

	return err
}

func cancelSubscription() error {
	s, err := subscription.Cancel(
		"sub_1Ng0NYJ39MhrdOSIi7myrYSA",
		nil,
	)

	fmt.Println(s.ID)

	return err
}

func createSubscription() error {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_OSw75Flxl2ZDU2"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String("price_1Ng02DJ39MhrdOSI8OnbmthF"),
			},
		},
	}
	s, err := subscription.New(params)

	fmt.Println("subscription: ", s.ID)

	return err
}

func createProductAndPrice() error {
	product_params := &stripe.ProductParams{
		Name:        stripe.String("Subscription product sample 1"),
		Description: stripe.String("2万円/月 subscription"),
	}
	starter_product, _ := product.New(product_params)

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(20000),
	}
	starter_price, err := price.New(price_params)

	fmt.Println("Success! product id: " + starter_product.ID)
	fmt.Println("Success! price id: " + starter_price.ID)

	return err
}

func createCustomer() error {
	params := &stripe.CustomerParams{
		Description: stripe.String("test customer5"),
	}
	c, err := customer.New(params)

	fmt.Println(c.ID)

	return err
}
