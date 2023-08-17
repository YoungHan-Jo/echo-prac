package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/card"
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

func createCard() error {
	params := &stripe.CardParams{
		Customer: stripe.String("cus_OSwy42WPhIJtir"),
		// Token:    stripe.String("tok_jcb"),
		Token: stripe.String("tok_1Ng2a4J39MhrdOSILWxRZrsl"),
	}
	c, err := card.New(params)

	fmt.Println(c)
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

// func createToken() error {
// 	params := &stripe.TokenParams{
// 		Card: &stripe.CardParams{
// 			Number:   stripe.String("5555555555554444"),
// 			ExpMonth: stripe.String("12"),
// 			ExpYear:  stripe.String("2024"),
// 			CVC:      stripe.String("123"),
// 			Name:     stripe.String("Sample B"),
// 		},
// 	}
// 	t, err := token.New(params)

// 	fmt.Println("t :", t)
// 	fmt.Println("id:", t.ID)
// 	fmt.Println("card:", t.Card)
// 	fmt.Println("livemode:", t.Livemode)

// 	return err
// }
