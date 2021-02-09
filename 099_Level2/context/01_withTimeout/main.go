package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func doRequest(ctx context.Context, route string) (string, error) {

	deadline, ok := ctx.Deadline()
	if ok && time.Until(deadline) < 100*time.Millisecond {
		return "", fmt.Errorf("Deadline too short")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, route, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed Satus Code : %d", res.StatusCode)
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := doRequest(ctx, "https://golang.org/")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
