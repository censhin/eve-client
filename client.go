package client

import (
    "encoding/xml"
    "fmt"
    "log"
    "net/http"

    "github.com/censhin/eve-poll/config"
)

type Client struct {
    BaseUrl string
    KeyId   string
    VCode   string
}

func NewClient() *Client {
    client := new(Client)
    conf := config.GetConfig()
    client.BaseUrl = conf.BaseUrl
    client.KeyId = conf.KeyId
    client.VCode = conf.VCode

    return client
}

func ApiV2Request(client *Client, resource string, params string, output interface{}) error {
    url := fmt.Sprintf("%v%v?keyID=%v&vCode=%v%v",
                       client.BaseUrl,
                       resource,
                       client.KeyId,
                       client.VCode,
                       params)

    resp, err := http.Get(url)
    if err != nil {
        log.Panic(err)
        return err
    }

    defer resp.Body.Close()

    err = xml.NewDecoder(resp.Body).Decode(&output)
    if err != nil {
        log.Panic(err)
        return err
    }

    return nil
}

