package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type SyncGetDetailProduitV3Response struct {
	Text           string `xml:",chardata"`
	Saq            string `xml:"saq,attr"`
	Xsi            string `xml:"xsi,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
	DataArea       struct {
		Text                       string `xml:",chardata"`
		GetDetailProduitV3Response struct {
			Text   string `xml:",chardata"`
			Return struct {
				Text        string `xml:",chardata"`
				Description struct {
					Text string `xml:",chardata"`
				} `xml:"description"`
				ID struct {
					Text string `xml:",chardata"`
				} `xml:"id"`
				PartNumber struct {
					Text string `xml:",chardata"`
				} `xml:"partNumber"`
				Format struct {
					Text string `xml:",chardata"`
				} `xml:"format"`
				IdentiteProduit struct {
					Text string `xml:",chardata"`
				} `xml:"identiteProduit"`
				Millesime struct {
					Text string `xml:",chardata"`
				} `xml:"millesime"`
				PastilleGout struct {
					Text string `xml:",chardata"`
				} `xml:"pastilleGout"`
				Pays struct {
					Text string `xml:",chardata"`
				} `xml:"pays"`
				Prix struct {
					Text string `xml:",chardata"`
				} `xml:"prix"`
				UrlPastille struct {
					Text string `xml:",chardata"`
				} `xml:"urlPastille"`
				QuantiteParEmballage struct {
					Text string `xml:",chardata"`
				} `xml:"quantiteParEmballage"`
				UrlProduit struct {
					Text string `xml:",chardata"`
				} `xml:"urlProduit"`
				NbPromotionsEnVigueur struct {
					Text string `xml:",chardata"`
				} `xml:"nbPromotionsEnVigueur"`
				IdPromotion struct {
					Text string `xml:",chardata"`
				} `xml:"idPromotion"`
				IndDispoEnLigne struct {
					Text string `xml:",chardata"`
				} `xml:"indDispoEnLigne"`
				QteDispoEnLigne struct {
					Text string `xml:",chardata"`
				} `xml:"qteDispoEnLigne"`
				IndDispoEnSuccursale struct {
					Text string `xml:",chardata"`
				} `xml:"indDispoEnSuccursale"`
				RemarqueLivraison struct {
					Text string `xml:",chardata"`
				} `xml:"remarqueLivraison"`
				ArgLong struct {
					Text string `xml:",chardata"`
				} `xml:"argLong"`
				ListeAttributs []struct {
					Text         string `xml:",chardata"`
					TypeAttribut struct {
						Text string `xml:",chardata"`
					} `xml:"typeAttribut"`
					Value struct {
						Text string `xml:",chardata"`
					} `xml:"value"`
				} `xml:"listeAttributs"`
			} `xml:"return"`
		} `xml:"getDetailProduitV3Response"`
	} `xml:"DataArea"`
}

type SyncGetSuccursalesResponse struct {
	Text           string `xml:",chardata"`
	Saq            string `xml:"saq,attr"`
	Xsi            string `xml:"xsi,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
	DataArea       struct {
		Text                   string `xml:",chardata"`
		GetSuccursalesResponse struct {
			Text   string `xml:",chardata"`
			Return []struct {
				Text    string `xml:",chardata"`
				Adresse struct {
					Text string `xml:",chardata"`
				} `xml:"adresse"`
				Banniere struct {
					Text string `xml:",chardata"`
				} `xml:"banniere"`
				Latitude struct {
					Text string `xml:",chardata"`
				} `xml:"latitude"`
				Longitude struct {
					Text string `xml:",chardata"`
				} `xml:"longitude"`
				NbProduit struct {
					Text string `xml:",chardata"`
				} `xml:"nbProduit"`
				Region struct {
					Text string `xml:",chardata"`
				} `xml:"region"`
				SuccursaleId struct {
					Text string `xml:",chardata"`
				} `xml:"succursaleId"`
				Telephone struct {
					Text string `xml:",chardata"`
				} `xml:"telephone"`
				Ville struct {
					Text string `xml:",chardata"`
				} `xml:"ville"`
				Ouvert struct {
					Text string `xml:",chardata"`
				} `xml:"ouvert"`
				HeureOuverture struct {
					Text string `xml:",chardata"`
				} `xml:"heureOuverture"`
				HeureFermeture struct {
					Text string `xml:",chardata"`
				} `xml:"heureFermeture"`
			} `xml:"return"`
		} `xml:"getSuccursalesResponse"`
	} `xml:"DataArea"`
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Soapenc string   `xml:"soapenc,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Header  struct {
		Text string `xml:",chardata"`
	} `xml:"Header"`
	Body struct {
		Text                           string                         `xml:",chardata"`
		SyncGetSuccursalesResponse     SyncGetSuccursalesResponse     `xml:"SyncGetSuccursalesResponse"`
		SyncGetDetailProduitV3Response SyncGetDetailProduitV3Response `xml:"SyncGetDetailProduitV3Response"`
	} `xml:"Body"`
}

type ListResponse struct {
	Metadata struct {
		NbPages           float64 `json:"nbPages"`
		NbResults         int     `json:"nbResults"`
		NbResultsPerPages int     `json:"nbResultsPerPages"`
		StartElement      int     `json:"startElement"`
	} `json:"Metadata"`
	Produits []struct {
		Description           string   `json:"description"`
		FamilleAccordList     []string `json:"familleAccordList,omitempty"`
		Format                string   `json:"format"`
		ID                    int      `json:"id"`
		IDPromotion           string   `json:"idPromotion"`
		IdentiteProduit       string   `json:"identiteProduit"`
		Millesime             string   `json:"millesime"`
		NbPromotionsEnVigueur string   `json:"nbPromotionsEnVigueur"`
		PartNumber            int      `json:"partNumber"`
		PastilleGout          string   `json:"pastilleGout,omitempty"`
		Pays                  string   `json:"pays"`
		Prix                  float64  `json:"prix"`
		QuantiteParEmballage  float64  `json:"quantiteParEmballage"`
		TypeSpiritueux        string   `json:"typeSpiritueux,omitempty"`
		URLPastille           string   `json:"urlPastille,omitempty"`
		URLTypeSpiritueux     string   `json:"urlTypeSpiritueux,omitempty"`
	} `json:"Produits"`
	SearchFacets []map[string]interface{} `json:"SearchFacets`
}

func sendLocationRequest(id int) (envelope Envelope) {
	// Location request (POST https://www.saq.com:9400/webapp/wcs/services/ServiceMobile)

	body := strings.NewReader(fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:mob="http://mobile.saq.com/">
		<soapenv:Body>
			<mob:SyncGetSuccursales>
				<mob:getSuccursales>
					<arg0>en</arg0>
					<arg1>%v</arg1>
				</mob:getSuccursales>
			</mob:SyncGetSuccursales>
		</soapenv:Body>
	</soapenv:Envelope>
	`, id))

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://www.saq.com:9400/webapp/wcs/services/ServiceMobile", body)

	// Headers
	req.Header.Add("SOAPAction", "https://www.saq.com:9400/webapp/wcs/services/ServiceMobile")
	req.Header.Add("Content-Type", "text/plain; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(respBody, &envelope)
	return
}

func sendListRequest(args *url.Values) (listResponse ListResponse) {

	// Create client
	client := &http.Client{}

	// Create request
	// fmt.Println("https://www.saq.com/wcs/resources/store/20002/saqsearch/searchfacets?" + args.Encode())
	req, err := http.NewRequest("GET", "https://www.saq.com/wcs/resources/store/20002/saqsearch/searchfacets?"+args.Encode(), nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(respBody, &listResponse)
	return
}

func sendItemRequest(id int) (envelope Envelope) {
	// Item request (POST https://www.saq.com:9400/webapp/wcs/services/ServiceMobile)

	body := strings.NewReader(fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:mob="http://mobile.saq.com/">
		<soapenv:Body>
			<mob:SyncGetDetailProduitV3>
				<mob:getDetailProduitV3>
					<lang>en</lang>		
					<produitId>%v</produitId>
				</mob:getDetailProduitV3>
			</mob:SyncGetDetailProduitV3>
		</soapenv:Body>
	</soapenv:Envelope>
	`, id))

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://www.saq.com:9400/webapp/wcs/services/ServiceMobile", body)

	// Headers
	req.Header.Add("SOAPAction", "https://www.saq.com:9400/webapp/wcs/services/ServiceMobile")
	req.Header.Add("Content-Type", "text/plain; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(respBody, &envelope)
	return
}

func main() {
	var ids []int

	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			i, _ := strconv.Atoi(v)
			if i != 0 {
				ids = append(ids, i)
			}
		}
	}

	// args := url.Values{}

	// sortMethods := []string{
	// 	"PERTINENCE",
	// 	"PRIX_CROISSANT",
	// 	"PRIX_DECROISSANT",
	// 	"NOM_A_Z",
	// 	"NOM_Z_A",
	// }

	// args.Set("langId", "en")
	// args.Set("nbResultPerPage", "20")
	// args.Set("startElement", "0")
	// args.Set("sort", sortMethods[0])
	// args.Set("showFacets", "true")
	// args.Set("showProducts", "true")
	// args.Set("int_available", "1")
	// args.Set("int_avail_store", "1")
	// args.Set("Categorie_marketing_level1", "Spirit")
	// args.Set("Categorie_marketing_level2", "Scotch and whisky")
	// args.Set("Categorie_marketing_level3", "Scotch single malt")
	// listResponse := sendListRequest(&args)

	// for _, v := range listResponse.Produits {
	// 	fmt.Println(v.ID, v.Description)
	// }

	type Entry struct {
		Banniere  string
		Adresse   string
		Latitude  string
		Longitude string
		Produits  map[string]string
	}

	success := make(map[string]Entry)

	for _, i := range ids {
		item := sendItemRequest(i).Body.SyncGetDetailProduitV3Response.DataArea.GetDetailProduitV3Response.Return

		envelope := sendLocationRequest(i)
		for _, p := range envelope.Body.SyncGetSuccursalesResponse.DataArea.GetSuccursalesResponse.Return {
			if strings.Contains(p.Banniere.Text, "Restauration") {
				continue
			}

			if _, ok := success[p.SuccursaleId.Text]; !ok {
				success[p.SuccursaleId.Text] = Entry{
					p.Banniere.Text,
					p.Adresse.Text + ", " + p.Region.Text + " (" + p.Ville.Text + ")",
					p.Latitude.Text,
					p.Longitude.Text,
					make(map[string]string),
				}
			}

			success[p.SuccursaleId.Text].Produits[item.Description.Text+" ("+item.ID.Text+")"] = p.NbProduit.Text
		}
	}

	for _, v := range success {
		if len(v.Produits) == len(ids) {
			spew.Dump(v)
		}
	}
}
