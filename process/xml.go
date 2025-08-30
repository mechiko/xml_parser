package process

import "encoding/xml"

// Documents was generated 2025-08-30 17:22:39 by https://xml-to-go.github.io/ in Ukraine.
type Documents struct {
	XMLName  xml.Name `xml:"Documents" json:"documents,omitempty"`
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Xsd      string   `xml:"xsd,attr" json:"xsd,omitempty"`
	Xsi      string   `xml:"xsi,attr" json:"xsi,omitempty"`
	Xmlxsi   string   `xml:"xmlxsi,attr" json:"xmlxsi,omitempty"`
	Version  string   `xml:"Version,attr" json:"version,omitempty"`
	Document struct {
		Text             string `xml:",chardata" json:"text,omitempty"`
		WayBillV3Transit struct {
			Text   string `xml:",chardata" json:"text,omitempty"`
			Header struct {
				Text         string `xml:",chardata" json:"text,omitempty"`
				Version      string `xml:"Version"`
				NUMBER       string `xml:"NUMBER"`
				Date         string `xml:"Date"`
				ShippingDate string `xml:"ShippingDate"`
				Note         string `xml:"Note"`
				Shipper      struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					ShipperId string `xml:"ShipperId"`
					Name      string `xml:"Name"`
				} `xml:"Shipper" json:"shipper,omitempty"`
				Consignee struct {
					Text    string `xml:",chardata" json:"text,omitempty"`
					FSRARID string `xml:"FSRAR_ID"`
					Name    string `xml:"Name"`
				} `xml:"Consignee" json:"consignee,omitempty"`
			} `xml:"Header" json:"header,omitempty"`
			Content struct {
				Text     string `xml:",chardata" json:"text,omitempty"`
				Position struct {
					Text         string `xml:",chardata" json:"text,omitempty"`
					Identity     string `xml:"Identity"`
					FullName     string `xml:"FullName"`
					EAN          string `xml:"EAN"`
					Alccode      string `xml:"Alccode"`
					BoxCapacity  string `xml:"BoxCapacity"`
					BottlingDate string `xml:"BottlingDate"`
					Quantity     string `xml:"Quantity"`
					MarkInfo     struct {
						Text   string `xml:",chardata" json:"text,omitempty"`
						Boxpos []struct {
							Text      string `xml:",chardata" json:"text,omitempty"`
							Boxnumber string `xml:"boxnumber"`
							Amclist   struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Amc  []struct {
									Text         string `xml:",chardata" json:"text,omitempty"`
									SerialNumber string `xml:"SerialNumber,attr" json:"serialnumber,omitempty"`
								} `xml:"amc" json:"amc,omitempty"`
							} `xml:"amclist" json:"amclist,omitempty"`
						} `xml:"boxpos" json:"boxpos,omitempty"`
					} `xml:"MarkInfo" json:"markinfo,omitempty"`
					BoxInfo struct {
						Text    string `xml:",chardata" json:"text,omitempty"`
						Boxtree struct {
							Text string `xml:",chardata" json:"text,omitempty"`
							Bl   []struct {
								Text   string `xml:",chardata" json:"text,omitempty"`
								Boxnum string `xml:"boxnum"`
								Bl     struct {
									Text   string   `xml:",chardata" json:"text,omitempty"`
									Boxnum []string `xml:"boxnum"`
								} `xml:"bl" json:"bl,omitempty"`
							} `xml:"bl" json:"bl,omitempty"`
						} `xml:"boxtree" json:"boxtree,omitempty"`
					} `xml:"boxInfo" json:"boxinfo,omitempty"`
				} `xml:"Position" json:"position,omitempty"`
			} `xml:"Content" json:"content,omitempty"`
		} `xml:"WayBill_v3_transit" json:"waybill_v3_transit,omitempty"`
	} `xml:"Document" json:"document,omitempty"`
}
