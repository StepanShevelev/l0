package api

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func InitBackendApi() {
	http.HandleFunc("/API/get_data", apiGetData)

}
func apiGetData(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("not method GET!"))
		return
	}

	orderUid, okId := ParseUid(w, r)
	if !okId {
		w.Write([]byte(`{"error": "can't pars name"}`))
		return
	}

	GetHtmlByName(orderUid, w)
}

func GetHtmlByName(orderUid string, w http.ResponseWriter) {

	order, err := Caching.GetCache(orderUid)
	if err != nil {
		logrus.Info("cant find data in cache")
		w.Write([]byte("incorrect id"))
		return
	}
	logrus.Info("Html generated")

	var tmpl = `
<html>
<head>
<title>

</title>
</head>
<body>


<ul>
<li> OrderUID: {{.OrderUID}}</li>
<li>Track number: {{.TrackNumber}}</li>
<li>Entry: {{.Entry}}</li>
<li>Locale: {{.Locale}}</li>
<li>Internal signature: {{.InternalSignature}}</li>
<li>Customer id: {{.CustomerId}}</li>
<li>Delivery service: {{.DeliveryService}}</li>
<li>Shard key: {{.Shardkey}}</li>
<li>Sm id: {{.SmId}}</li>
<li>Date created: {{.DateCreated}}</li>
<li>Oof shard: {{.OofShard}}</li>
</ul>

<p>Delivery</p>
<ul>
<li>Name: {{.Delivery.Name}}</li>
<li>Phone: {{.Delivery.Phone}}</li>
<li>Zip: {{.Delivery.Zip}}</li>
<li>City: {{.Delivery.City}}</li>
<li>Address: {{.Delivery.Address}}</li>
<li>Region: {{.Delivery.Region}}</li>
<li>Email: {{.Delivery.Email}}</li>
</ul>

<p>Payment</p>
<ul>
<li>Transaction: {{.Payment.Transaction}}</li>
<li>Request id: {{.Payment.RequestId}}</li>
<li>Currency: {{.Payment.Currency}}</li>
<li>Provider: {{.Payment.Provider}}</li>
<li>Amount: {{.Payment.Amount}}</li>
<li>Payment dt: {{.Payment.PaymentDt}}</li>
<li>Bank: {{.Payment.Bank}}</li>
<li>Delivery cost: {{.Payment.DeliveryCost}}</li>
<li>Goods total: {{.Payment.GoodsTotal}}</li>
<li>Custom fee: {{.Payment.CustomFee}}</li>
</ul>

<p>Items</p>
<ul>
{{range .Items}}
<br>
<li>Chrt Id: {{.ChrtId}}</li>
<li>Track Number: {{.TrackNumber}}</li>
<li>Price: {{.Price}}</li>
<li>Rid: {{.Rid}}</li>
<li>Name: {{.Name}}</li>
<li>Sale: {{.Sale}}</li>
<li>Size: {{.Size}}</li>
<li>TotalPrice: {{.TotalPrice}}</li>
<li>Nm Id: {{.NmId}}</li>
<li>Brand: {{.Brand}}</li>
<li>Status: {{.Status}}</li>
{{end}}
</ul>


</body>
</html>
`

	a := order

	// Make and parse the HTML template
	t, err := template.New("111").Funcs(template.FuncMap{
		"btoa": func(b []byte) string { return string(b) },
	}).Parse(tmpl)
	if err != nil {
		logrus.Info("Error occurred while creating new template", err)
		return
	}
	err = t.Execute(w, a)
	if err != nil {
		logrus.Info("Error occurred while updating file data", err)
		return
	}

}

func ParseUid(w http.ResponseWriter, r *http.Request) (string, bool) {
	keys, ok := r.URL.Query()["uid"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "arguments params are missing"}`))
		return " ", false
	}
	orderUid := keys[0]

	return orderUid, true
}

func isMethodGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}
