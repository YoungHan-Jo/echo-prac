package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/card"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/customer"
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

// func createInvoice() error {
// 	params := &stripe.InvoiceParams{
// 		Customer: stripe.String("cus_OSbCLY7xLaBtTK"),
// 	}
// 	in, err := invoice.New(params)

// 	fmt.Println(in.ID)

// 	return err
// }

// func createPaymentIntent() error {
// 	params := &stripe.PaymentIntentParams{
// 		Amount:   stripe.Int64(40000),
// 		Currency: stripe.String(string(stripe.CurrencyJPY)),
// 		Customer: stripe.String("cus_OSvPe6Edz92bsw"), // test customer1 ( cus_OSbCLY7xLaBtTK )
// 	}
// 	pi, err := paymentintent.New(params)

// 	fmt.Println(pi.ID)

// 	return err
// }

// func confirmPaymentIntent() error {
// 	params := &stripe.PaymentIntentConfirmParams{
// 		PaymentMethod: stripe.String("pm_card_visa"),
// 		// PaymentMethod: stripe.String("pm_1NhiwyJ39MhrdOSIOpEM1Zjr"),
// 		// PaymentMethod: stripe.String("pm_card_visa_chargeDeclined"),
// 	}
// 	pi, err := paymentintent.Confirm(
// 		"pi_3NhivsJ39MhrdOSI00YtoeJc",
// 		params,
// 	)

// 	fmt.Println("pi status: ", pi.Status)

// 	return err
// }

// func createPaymentMethod() error {
// 	params := &stripe.PaymentMethodParams{
// 		Card: &stripe.PaymentMethodCardParams{
// 			// Token: stripe.String("tok_1NhVUUJ39MhrdOSI3b3W7NIG"),
// 			Token: stripe.String("tok_visa_chargeDeclined"),
// 		},
// 		Type: stripe.String("card"),
// 	}
// 	pm, err := paymentmethod.New(params)

// 	fmt.Println(pm.ID)

// 	return err
// }

func createSession() error {
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price: stripe.String("price_1NfzmjJ39MhrdOSI1lK8jKJ3"),
				// Price:    stripe.String("price_wrong_id"),
				Quantity: stripe.Int64(3),
			},
		},
		Customer:   stripe.String("cus_OSvPe6Edz92bsw"),            // customer
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

// func createProductAndPrice() error {
// 	product_params := &stripe.ProductParams{
// 		Name:        stripe.String("Subscription product sample 1"),
// 		Description: stripe.String("2万円/月 subscription"),
// 	}
// 	starter_product, _ := product.New(product_params)

// 	price_params := &stripe.PriceParams{
// 		Currency: stripe.String(string(stripe.CurrencyJPY)),
// 		Product:  stripe.String(starter_product.ID),
// 		Recurring: &stripe.PriceRecurringParams{
// 			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
// 		},
// 		UnitAmount: stripe.Int64(20000),
// 	}
// 	starter_price, err := price.New(price_params)

// 	fmt.Println("Success! product id: " + starter_product.ID)
// 	fmt.Println("Success! price id: " + starter_price.ID)

// 	return err
// }

// func getPaymentIntent() error {
// 	pi, err := paymentintent.Get(
// 		"pi_3NhO4GJ39MhrdOSI1BxQ5YxD",
// 		nil,
// 	)

// 	fmt.Println(pi.Status)

// 	return err
// }

func createCustomer() error {
	params := &stripe.CustomerParams{
		Description: stripe.String("subscription failure test customer"),
	}
	c, err := customer.New(params)

	fmt.Println(c.ID)

	return err
}

func createCard() error {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_OUnfSvQazEfN9k"),
		Token:    stripe.String("tok_1Ninx6J39MhrdOSIdLidY2oj"),
		// Token: stripe.String("tok_wrong_token"),
	}
	c, err := card.New(params)

	fmt.Println(c.ID)
	return err
}

func createTestProductAndPrice() error {
	product_params := &stripe.ProductParams{
		Name:        stripe.String("product sample 1"),
		Description: stripe.String("1万円"),
	}
	starter_product, _ := product.New(product_params)

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		Product:  stripe.String(starter_product.ID),
		// Recurring: &stripe.PriceRecurringParams{
		// 	Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		// },
		UnitAmount: stripe.Int64(10000),
	}
	starter_price, err := price.New(price_params)

	fmt.Println("Success! product id: " + starter_product.ID)
	fmt.Println("Success! price id: " + starter_price.ID)

	return err
}

func createSubscription() error {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_OUnfSvQazEfN9k"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String("price_1Ng02DJ39MhrdOSI8OnbmthF"),
				// Price: stripe.String("price_1NhUdqJ39MhrdOSIQqNwAlDO"), // not subscription
				// Price: stripe.String("price_wrong_id"),
			},
		},
	}
	s, err := subscription.New(params)

	fmt.Println("subscription: ", s.ID)

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
