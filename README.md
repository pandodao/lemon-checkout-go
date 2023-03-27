# lemon-checkout-go

A Simple Go SDK package for https://lemonsqueezy.com

## Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/pandodao/lemon-checkout-go"
)

const key = "YOUR_API_KEY"

type Custom struct {
	// put your custom data here
	MyTraceID string `json:"my_trace_id"`
}

func main() {
	custom := Custom{
		MyTraceID: "123456",
	}

	ctx := context.Background()

	client := lemon.New(key, lemon.WithDebug(true))

	resp, err := client.CreateCheckoutSimple(ctx, YOUR_STORE_ID, []int64{YOUR_VARIANT_ID}, "YOUR_REDIRECT_URL", lemon.CheckoutData{
		Custom: custom,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("payment URL: %+v\n", resp.Attributes.URL)
}
```

## Options

* WithDebug - The relevant request and response logs will be printed out.
* WithLogger - Custom logger for debug log.

