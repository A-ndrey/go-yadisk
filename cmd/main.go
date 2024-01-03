package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/A-ndrey/go-yadisk"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <cmd>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\ncmd:\n%s\n", strings.Join(Methods(), "\t\n"))
		fmt.Fprintf(os.Stderr, "\noptions:\n")
		flag.PrintDefaults()
	}
	host := flag.String("h", "cloud-api.yandex.net", "Host for the all requests to API")
	params := flag.String("p", "{}", "Params for request in json format")
	flag.Parse()

	cmd := flag.Arg(0)

	token := os.Getenv("YA_DISK_TOKEN")

	client := yadisk.NewClient(token, *host, yadisk.WithHttpClient(&http.Client{Timeout: time.Minute}))
	Call(client, cmd, *params)
}

func Methods() []string {
	var res []string
	clientType := reflect.TypeOf(&yadisk.Client{})
	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		if !method.IsExported() {
			continue
		}
		res = append(res, method.Name)
	}

	return res
}

func Call(client *yadisk.Client, cmd string, params string) {
	clientValue := reflect.ValueOf(client)
	clientType := reflect.TypeOf(client)
	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		if !method.IsExported() {
			continue
		}
		if method.Name != cmd {
			continue
		}

		paramsType := method.Type.In(2)
		paramsInst := reflect.New(paramsType).Interface()
		if err := json.Unmarshal([]byte(params), paramsInst); err != nil {
			fmt.Printf("can't parse params for type %s, %s", paramsType.Name(), err.Error())
			return
		}

		res := clientValue.Method(i).Call([]reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(paramsInst).Elem()})
		if !res[1].IsNil() {
			fmt.Println(res[1].Interface().(error).Error())
			return
		}
		if res[0].IsNil() {
			return
		}

		data, err := json.MarshalIndent(res[0].Interface(), "", "\t")
		if err != nil {
			fmt.Printf("can't marshal response, %s", err.Error())
			return
		}

		fmt.Println(string(data))
	}
}
