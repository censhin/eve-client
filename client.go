package client

import (
    "encoding/xml"
    "fmt"
    "log"
    "net/http"
)

type Client struct {
    BaseUrl string
    KeyId   string
    VCode   string
}

func NewClient(baseUrl string, keyId string, vCode string) *Client {
    client := new(Client)
    client.BaseUrl = baseUrl
    client.KeyId = keyId
    client.VCode = vCode

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

