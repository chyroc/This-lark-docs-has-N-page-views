# This-lark-docs-has-N-page-views

lark docs: **[深度好文！这篇文档被阅读/点赞了 20/7 次](https://bytedance.feishu.cn/docs/doccn4uyl8Q8v3tviYpCCmH81Uh)**

## 最简化代码

```go
package main

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/doc"
	"github.com/chyroc/lark/larkext"
)

func example(ctx context.Context, appID, appSecret, docToken string) error {
	larkClient := lark.New(lark.WithAppCredential(appID, appSecret))
	docIns := larkext.NewDoc(larkClient, docToken)

	statistics, err := docIns.Statistics(ctx)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("这篇文档被阅读/点赞 %d/%d 次", statistics.Pv, statistics.LikeCount)
	return docIns.Update(ctx, doc.UpdateTitle(title))
}
```

## 安装

```shell
go install github.com/chyroc/This-lark-docs-has-N-page-views@latest
```

## 使用

```shell
This-lark-docs-has-N-page-views -h

Usage of This-lark-docs-has-N-page-views:
  -app_id string
    	lark app id
  -app_secret string
    	lark app secret
  -doc_token string
    	lark docs token(in docs url)
```

```shell
This-lark-docs-has-N-page-views -app_id "<lark app id>" -app_secret "<lark app secret>" -doc_token "<lark docs token>"
```

## 依赖

- 依赖 https://github.com/chyroc/lark sdk 实现
