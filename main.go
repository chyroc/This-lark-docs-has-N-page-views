package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/doc"
	"github.com/chyroc/lark/larkext"
)

func main() {
	appID := ""
	appSecret := ""
	docToken := ""

	flag.StringVar(&appID, "app_id", "", "lark app id")
	flag.StringVar(&appSecret, "app_secret", "", "lark app secret")
	flag.StringVar(&docToken, "doc_token", "", "lark docs token(in docs url)")
	flag.Parse()

	// first: init lark sdk client
	cli := NewLarkDocTitleUpdater(lark.New(lark.WithAppCredential(appID, appSecret)))

	// second: loop update
	for {
		_ = cli.Update(docToken)
		time.Sleep(time.Second)
	}
}

type LarkDocTitleUpdater struct {
	cli       *lark.Lark
	lock      sync.Mutex
	pvCount   map[string]int64
	likeCount map[string]int64
}

func NewLarkDocTitleUpdater(cli *lark.Lark) *LarkDocTitleUpdater {
	return &LarkDocTitleUpdater{
		cli:       cli,
		pvCount:   map[string]int64{},
		likeCount: map[string]int64{},
	}
}

func (r *LarkDocTitleUpdater) Update(docToken string) (err error) {
	defer func() {
		if err != nil {
			fmt.Printf("update fail: %s\n", err)
		}
	}()
	ctx := context.Background()
	docIns := larkext.NewDoc(r.cli, docToken)

	statistics, err := docIns.Statistics(ctx)
	if err != nil {
		return err
	}

	r.lock.Lock()
	defer r.lock.Unlock()
	if r.pvCount[docToken] == statistics.Pv && r.likeCount[docToken] == statistics.LikeCount {
		fmt.Printf("no change, skip\n")
		return nil
	}

	r.pvCount[docToken] = statistics.Pv
	r.likeCount[docToken] = statistics.LikeCount
	title := fmt.Sprintf("这篇文档被阅读/点赞 %d/%d 次", statistics.Pv, statistics.LikeCount)

	if err = docIns.Update(ctx, doc.UpdateTitle(title)); err != nil {
		return err
	}
	fmt.Printf("update pv=%d, like=%d\n", statistics.Pv, statistics.LikeCount)
	return nil
}
