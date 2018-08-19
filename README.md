# Jobcan Server

ジョブカンの勝手APIサーバーです。

## Read Status

現在の勤怠ステータスを取得します。

### Request

```
POST /status
```

#### Body

Property   |Type    |Description
-----------|--------|-------------
`client_id`|`string`|Client ID of Jobcan Account
`email`    |`string`|Email Address of Jobcan Account
`password` |`string`|Password of Jobcan Account

### Response

```
{
  status: "having_breakfast"|"resting"|"working"
}
```

## Punch

打刻をします。

### Request

```
POST /punch
```

#### Body

Property   |Type    |Description
-----------|--------|-------------
`client_id`|`string`|Client ID of Jobcan Account
`email`    |`string`|Email Address of Jobcan Account
`password` |`string`|Password of Jobcan Account

### Response

```
{
  status: "having_breakfast"|"resting"|"working"
}
```
