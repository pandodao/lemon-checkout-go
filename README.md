# lemon-checkout-go

A Simple Go SDK package for https://lemonsqueezy.com

## Example: Create checkout

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

### Example: Handle webhook

```go
import (
	"github.com/pandodao/lemon-checkout-go"
)

type (
	LemonWebhookPayload struct {
		lemon.WebhookPayload
		Meta struct {
			EventName  string `json:"event_name"`
			CustomData struct {
				MyTraceID    string `json:"my_trace_id"`
			} `json:"custom_data"`
		} `json:"meta"`
	}
)

func HandleLemonWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body := &LemonWebhookPayload{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&body); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		if body.Meta.EventName != "order_created" {
			w.Write([]byte("pass"))
			return
		}

		if body.Data.Attributes.Status != "paid" {
			w.Write([]byte("pass"))
			return
		}

		myTraceID := body.Meta.CustomData.MyTraceID

		// query `myTraceID` in your database
		// ...

		render.JSON(w, "ok")
	}
}

```

## Options

* WithDebug - The relevant request and response logs will be printed out.
* WithLogger - Custom logger for debug log.

