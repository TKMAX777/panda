# PandA API for Golang

## 概要
PandAのGo言語向けAPI。今のところAssignmentsとContentしか実装してない。

2021/04/21追記: お気に入りサイトの取得機能追加

## 目次
<!-- TOC -->

- [PandA API for Golang](#panda-api-for-golang)
    - [概要](#概要)
    - [目次](#目次)
    - [使い方](#使い方)

<!-- /TOC -->

## 使い方
```
go get -u github.com/TKMAX777/panda
```

```go
import "github.com/TKMAX777/panda"

func main() {
    Panda := panda.NewClient()
    err := Panda.Login(ECS_ID, PASSWORD)
    if err != nil {
        panic(err)
    }
    ...
}
```

