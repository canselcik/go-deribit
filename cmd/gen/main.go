package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type data struct {
	FuncName, Args, Format, Type, AccessLevel, Channel, Package string
}

var subscriptions = map[string]struct {
	funcName, notificationType string
	isPrivate                  bool
}{
	"announcements": {"Announcements", "AnnouncementNotification", true},
	"book.{instrument_name}.{group}.{depth}.{interval}": {"BookGroup", "BookNotification", false},
	"book.{instrument_name}.{interval}":                 {"BookInterval", "BookNotificationRaw", false},
	"deribit_price_index.{index_name}":                  {"DeribitPriceIndex", "DeribitPriceIndexNotification", false},
	"deribit_price_ranking.{index_name}":                {"DeribitPriceRanking", "DeribitPriceRankingNotification", false},
	"estimated_expiration_price.{index_name}":           {"EstimatedExpirationPrice", "EstimatedExpirationPriceNotification", false},
	"markprice.options.{index_name}":                    {"MarkPriceOptions", "MarkpriceOptionsNotification", false},
	"perpetual.{instrument_name}.{interval}":            {"Perpetual", "PerpetualNotification", false},
	"quote.{instrument_name}":                           {"Quote", "QuoteNotification", false},
	"ticker.{instrument_name}.{interval}":               {"Ticker", "TickerNotification", false},
	"trades.{instrument_name}.{interval}":               {"Trades", "PublicTrade", false},
	"user.orders.{instrument_name}.{interval}":          {"UserOrdersInstrumentName", "Order", true},
	"user.orders.{kind}.{currency}.{interval}":          {"UserOrdersKind", "Order", true},
	"user.portfolio.{currency}":                         {"UserPortfolio", "UserPortfolioNotification", true},
	"user.trades.{instrument_name}.{interval}":          {"UserTradesInstrument", "UserTrade", true},
	"user.trades.{kind}.{currency}.{interval}":          {"UserTradesKind", "UserTrade", true},
}

func main() {
	var d data
	var t *template.Template

	fmt.Println("// Code generated by make generate-methods DO NOT EDIT.\n\npackage deribit")

	for c, params := range subscriptions {
		d.Args = ""
		d.Channel = c
		d.FuncName = "Subscribe" + params.funcName
		d.Type = params.notificationType
		d.AccessLevel = "Public"
		d.Package = "public"
		if params.isPrivate {
			d.AccessLevel = "Private"
			d.Package = "private"
		}
		re := regexp.MustCompile(`\{(.*?)\}`)
		match := re.FindAllStringSubmatch(c, -1)
		if len(match) > 0 {
			args := make([]string, len(match))
			for i, m := range match {
				args[i] = m[1]
			}
			d.Args = strings.Join(args, ", ")
			d.Format = re.ReplaceAllString(c, "%s")
		} else {
			d.Format = c
		}
		t = template.Must(template.New("subMethod").Parse(subscriptionTemplate))
		err := t.Execute(os.Stdout, d)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var subscriptionTemplate = `
// {{.FuncName}} subscribes to the {{.Channel}} channel
func (e *Exchange) {{.FuncName}}({{.Args}}{{if .Args}} string{{end}}) (chan models.{{.Type}}, error) {
	chans := []string{ {{if .Args}}fmt.Sprintf("{{.Format}}", {{.Args}} ) {{else}}"{{.Format}}"{{end}} }
	
	c := make(chan *RPCNotification)
	out := make(chan models.{{.Type}})
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.calls.addSubscription(sub)

    client := e.Client().{{.AccessLevel}}
	if _, err := client.Get{{.AccessLevel}}Subscribe(&{{.Package}}.Get{{.AccessLevel}}SubscribeParams{Channels: chans}); err != nil {
		e.calls.deleteSubscription(chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{': 
					var ret models.{{.Type}}
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- ret
				case '[': 
					var rets []models.{{.Type}}
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}
`
