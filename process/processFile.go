package process

import (
	"encoding/json"
	"fmt"
	"net/http"
	"processXML/decompress"
	"processXML/models"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/Jeffail/gabs"
	"github.com/thedevsaddam/gojsonq"
)

func ProcessXML(rw http.ResponseWriter, r *http.Request) {

	if result, err := decompress.DecompressXML() ; err != nil {
		models.SendNotFound(rw)
	} else {
		xml := strings.NewReader(result)
		jsonResult, err := xj.Convert(xml)
		if err != nil {
			fmt.Println(err)
		}
		data := jsonResult.String()		

		jsonData := []byte(data)
		jsonParsed, err := gabs.ParseJSON(jsonData)
		if err != nil {
			panic(err)
		}

		detail := "{\"Detalle_Partidos\":" + jsonParsed.Path("Consolidado.Boletin.Detalle_Circunscripcion.lin.Detalle_Partidos_Totales.lin").String() +", \"Detalle_Candidatos\":"+ jsonParsed.Path("Consolidado.Boletin.Detalle_Circunscripcion.lin.Detalle_Candidato.lin").String()+"}"

		jq := gojsonq.New().JSONString(detail)
		res := jq.From("Detalle_Candidatos").Where("Candidato.-V", "=", "001").OrWhere("Candidato.-V", "=", "002").Get()

		detalleCandidatos, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		finalDetail := "{\"Detalle_Candidatos\":"+ string(detalleCandidatos)+",\"Detalle_Partidos\":" + jsonParsed.Path("Consolidado.Boletin.Detalle_Circunscripcion.lin.Detalle_Partidos_Totales.lin").String() +"}"

		models.SendData(rw, json.RawMessage(finalDetail))

	}

}