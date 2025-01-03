package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/A-ndrey/go-yadisk"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] <cmd>\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "\ncmd:\n%s\n", strings.Join(methods(), "\t\n"))
		fmt.Fprintf(flag.CommandLine.Output(), "\noptions:\n")
		flag.PrintDefaults()
	}
	host := flag.String("h", yadisk.DefaultHost, "Host for the all requests to API")
	params := flag.String("p", "{}", "Request's parameters in json format")
	downloadURL := flag.String("d", "", "Download URL")
	uploadURL := flag.String("u", "", "Upload URL")
	filePath := flag.String("f", "", "File path for downloading/uploading")
	flag.Parse()

	cmd := flag.Arg(0)

	if *downloadURL != "" && *uploadURL != "" {
		fmt.Fprintln(os.Stderr, "only one of the flags [-d, -u] may be specified")
		os.Exit(1)
	}

	token := os.Getenv("YA_DISK_TOKEN")

	client := yadisk.NewClient(token, *host, yadisk.WithHttpClient(&http.Client{Timeout: time.Minute}))

	var err error
	if *uploadURL != "" {
		err = upload(client, *uploadURL, *filePath)
	} else if *downloadURL != "" {
		err = download(client, *downloadURL, *filePath)
	} else {
		err = call(client, cmd, *params)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func methods() []string {
	var res []string
	clientType := reflect.TypeOf(&yadisk.Client{})
	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		if !method.IsExported() {
			continue
		}
		if method.Type.NumIn() < 3 {
			continue
		}

		paramsType := method.Type.In(2)
		if !strings.HasSuffix(paramsType.Name(), "Params") {
			continue
		}
		res = append(res, method.Name)
	}

	return res
}

func upload(client *yadisk.Client, destURL string, filePath string) error {
	r := os.Stdin
	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		r = file
	}
	return client.Upload(context.Background(), destURL, r)
}

func download(client *yadisk.Client, destURL string, filePath string) error {
	r := os.Stdout
	if filePath != "" {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		r = file
	}
	return client.Download(context.Background(), destURL, r)
}

func call(client *yadisk.Client, cmd string, params string) error {
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

		if method.Type.NumIn() < 3 {
			continue
		}

		paramsType := method.Type.In(2)
		if !strings.HasSuffix(paramsType.Name(), "Params") {
			continue
		}
		paramsInst := reflect.New(paramsType).Interface()
		if err := json.Unmarshal([]byte(params), paramsInst); err != nil {
			return fmt.Errorf("can't parse params for type %s: %w", paramsType.Name(), err)
		}

		res := clientValue.Method(i).Call([]reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(paramsInst).Elem()})
		if !res[1].IsNil() {
			return res[1].Interface().(error)
		}
		if res[0].IsNil() {
			return nil
		}

		jdec := json.NewEncoder(os.Stdout)
		jdec.SetEscapeHTML(false)
		jdec.SetIndent("", "\t")
		if err := jdec.Encode(res[0].Interface()); err != nil {
			return fmt.Errorf("can't marshal response: %w", err)
		}
	}

	return nil
}
