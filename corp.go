package client

import (
    "fmt"
    "log"
)

func (client *Client) CorpAccountBalance(characterId string) (*AccountBalance, error) {
    resource := "/corp/AccountBalance.xml.aspx"
    output := AccountBalance{}

    err := ApiV2Request(client, resource, "", &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

func (client *Client) CorpAssetList(characterId string) (*AssetList, error) {
    resource := "/corp/AssetList.xml.aspx"
    params := fmt.Sprintf("?characterID=%v", characterId)
    output := AssetList{}

    err := ApiV2Request(client, resource, params, &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

func (client *Client) StarbaseDetail(itemId string) (*StarbaseDetail, error) {
    resource := "/corp/StarbaseDetail.xml.aspx"
    params := fmt.Sprintf("?itemID=%v", itemId)
    output := StarbaseDetail{}

    err := ApiV2Request(client, resource, params, &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

func (client *Client) StarbaseList() (*StarbaseList, error) {
    resource := "/corp/StarbaseList.xml.aspx"
    output := StarbaseList{}

    err := ApiV2Request(client, resource, "", &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

