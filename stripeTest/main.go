package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
)

func main() {
	stripe.Key = "sk_test_51NfeKZJ39MhrdOSIAQdurvWMSNNckcKJqUvYtbyRvh3AVzclc9IFw6eGdbyI8pFfDAfIfw2ivbjjLcSf5PDCxLlt001sSESrDj"

	// result, _ := price.Get("price_1NfeXsJ39MhrdOSI4ukNQIwa", nil)
	// price.Search()
	// fmt.Println(result)
}

func createCustomer() {
	params := &stripe.CustomerParams{
		Description: stripe.String("test customer1"),
	}
	c, _ := customer.New(params)

	fmt.Println(c)
}

func createProductAndPrice() {
	product_params := &stripe.ProductParams{
		Name:        stripe.String("Subscription product test1"),
		Description: stripe.String("１万円/月 subscription"),
	}
	starter_product, _ := product.New(product_params)

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(10000),
	}
	starter_price, _ := price.New(price_params)

	fmt.Println("Success! Here is your starter subscription product id: " + starter_product.ID)
	fmt.Println("Success! Here is your starter subscription price id: " + starter_price.ID)
}
