package client

import (
    "fmt"
    "log"
)

func (client *Client) CharNotifications (characterId string) (*NotificationList, error) {
    resource := "/char/Notifications.xml.aspx"
    params := fmt.Sprintf("&characterID=%v", characterId)
    output := NotificationList{}

    err := ApiV2Request(client, resource, params, &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

func (client *Client) CharNotificationTexts (characterId string, notificationIds string) (*NotificationTextList, error) {
    // notificationIds should be a comma separated list of IDs returned from CharNotifications()
    // passed in a string
    resource := "/char/NotificationTexts.xml.aspx"
    params := fmt.Sprintf("&characterID=%v&IDs=%v", characterId, notificationIds)
    output := NotificationTextList{}

    err := ApiV2Request(client, resource, params, &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

