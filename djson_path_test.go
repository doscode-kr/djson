package djson

import (
	"log"
	"testing"
)

func TestPutPath(t *testing.T) {
	jsonDoc := `[
		{
			"name":"Ricardo Longa",
			"idade":28,
			"skills":[
				"Golang","Android"
			]
		},
		{
			"name":"Hery Victor",
			"idade":32,
			"skills":[
				"Golang",
				"Java"
			]
		}
	]`

	aJson := New().Parse(jsonDoc)

	err := aJson.UpdatePath(`[1]["name"]`, Object{
		"first":  "kim",
		"family": "kim",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(aJson.String())

	err = aJson.UpdatePath(`[1]["name"]["first"]`, "seo")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(aJson.String())

	err = aJson.PushBackPath(`[1]["skills"]`, "kotlin")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(aJson.String())

	err = aJson.RemovePath(`[1]["name"]["family"]`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(aJson.String())

	err = aJson.RemovePath(`[1]["name"]`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(aJson.String())

	err = aJson.RemovePath(`[1]`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(aJson.String())
}

func TestGetAsArrayObjectPath(t *testing.T) {
	jsonDoc := `{
		"hospital":{
		  "hospital_name":"록스병원",
		  "doctor_name":"김의사",
		  "department":"신경과"
		},
		"medicines": [ {
		  "name": "타이레놀",
		  "dose_event" : [
			{
			  "date" : "2021-02-02",
			  "time" : ["#B+30","#L+60"]
			}
		  ]
		}
		] 
	  }`

	aJson := New().Parse(jsonDoc)

	aJson.UpdatePath(`[medicines][2]`, "010-1234-5665")

	log.Println(aJson.ToString())

	dJson, ok := aJson.Array("medicines")
	if !ok {
		log.Fatal("GetAsArray() failed")
	}

	log.Println(dJson.ToString())

	pJson, ok := dJson.ArrayPath(`[0]["dose_event"]`)
	if !ok {
		log.Fatal("GetAsArrayPath() failed")
	}

	log.Println(pJson.ToString())
}

func TestGetKeysPath(t *testing.T) {
	jsonDoc := `{
		"hospital":{
		  "hospital_name":"록스병원",
		  "doctor_name":"김의사",
		  "department":"신경과"
		},
		"medicines": [ {
		  "name": "타이레놀",
		  "dose_event" : [
			{
			  "date" : "2021-02-02",
			  "time" : ["#B+30","#L+60"]
			}
		  ]
		}
		] 
	  }`

	aJson := New().Parse(jsonDoc)

	log.Println(aJson.GetKeys("hospital"))
	log.Println(aJson.GetKeys("medicines"))
	log.Println(aJson.KeysPath(`[hospital]`))
	log.Println(aJson.KeysPath(`[medicines][0][dose_event][0]`))

}

func TestUpdatePath2(t *testing.T) {
	jsonDoc := `[{
		"id":"111",
		"name" :"222"
	}]`

	aJson := New().Parse(jsonDoc)
	bJson := aJson.Clone()

	aJson.UpdatePath(`[0]["xxx"]`, "xxxx")

	log.Println(aJson.ToString())
	log.Println(bJson.ToString())

}
