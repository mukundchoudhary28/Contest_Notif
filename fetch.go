package main

import ("encoding/json"
		"fmt"
		"io/ioutil"
		"log"
		"net/http"
		"time")
		

type Info struct{
	Info []Contest `json:"result"`
}

type Contest struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Phase string `json:"phase"`
	Frozen bool `json:"frozen"`
	Duration int64 `json:"durationSeconds"`
	Start int64 `json:"startTimeSeconds"`
	Relative int64 `json:"relativeTimeSeconds"`

}

type Output struct{
	Name string
	Start_time string
}

func main(){
	url := "https://codeforces.com/api/contest.list?gym=false"
	spaceClient := http.Client{
		Timeout: time.Second*5,
	}

	req,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil{
		log.Fatalln(err)
	}

	res, _ := spaceClient.Do(req)


	if res.Body != nil{
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		log.Fatalln(readErr)

	}


	contests := Info{}
	jsonErr := json.Unmarshal(body,&contests)
	if jsonErr != nil{
		log.Fatalln(jsonErr)
	}

	//fmt.Println(contests.Info.Start)

	currentTime := time.Now()
	now := currentTime.Format("02-Jan-2006")
	//fmt.Println(now)



	for i:=0;;i++{
		
		if contests.Info[i].Phase != "BEFORE"{
			break
		}

		var timestamp int64 = contests.Info[i].Start

		net_time := time.Unix(timestamp,0)
		date := net_time.Format("02-Jan-2006")
		final := net_time.Format("02-Jan-2006 03:04 pm")

		if now != date{
			continue
		}

		out := Output{
			Name: contests.Info[i].Name,	
			Start_time: final,
		}

		outJSON,_ := json.Marshal(out)
		fmt.Println(string(outJSON))
		fmt.Println()

	}
}
