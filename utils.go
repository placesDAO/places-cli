package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func Ok(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (token *Token) print() {
	t := table.NewWriter()

	tTemp := table.Table{}
	tTemp.Render()

	t.SetStyle(table.StyleBold)
	t.SetTitle("Token #" + token.Id.String())
	t.Style().Title.Align = text.AlignCenter
	t.AppendRow([]interface{}{"Name", token.Place.Name})
	t.AppendRow([]interface{}{"Owner", token.Owner.String()})
	t.AppendRow([]interface{}{"StreetAddress", token.Place.StreetAddress})
	t.AppendRow([]interface{}{"Sublocality", token.Place.Sublocality})
	t.AppendRow([]interface{}{"Locality", token.Place.Locality})
	t.AppendRow([]interface{}{"SubadministrativeArea", token.Place.SubadministrativeArea})
	t.AppendRow([]interface{}{"AdministrativeArea", token.Place.AdministrativeArea})
	t.AppendRow([]interface{}{"PostalCode", token.Place.PostalCode})
	t.AppendRow([]interface{}{"Country", token.Place.Country})
	t.AppendRow([]interface{}{"CountryCode", token.Place.CountryCode})

	t.AppendRow([]interface{}{"Latitude", token.Place.Location.Latitude})
	t.AppendRow([]interface{}{"Longitude", token.Place.Location.Longitude})
	t.AppendRow([]interface{}{"Altitude", token.Place.Location.Altitude})

	t.AppendRow([]interface{}{"LatitudeInt", token.Place.Location.LatitudeInt})
	t.AppendRow([]interface{}{"LongitudeInt", token.Place.Location.LongitudeInt})
	t.AppendRow([]interface{}{"AltitudeInt", token.Place.Location.AltitudeInt})

	for _, attribute := range token.Place.Attributes {
		t.AppendRow([]interface{}{"Trait", attribute})
	}

	fmt.Println(t.Render())
}
