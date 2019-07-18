package hc

import (
	"fmt"
	"github.com/blocktree/openwallet/log"
	"testing"
)

type HCBlockObserver struct {
}

//BlockScanNotify 新区块扫描完成通知
func (o *HCBlockObserver) HCBlockScanNotify(block *Block, txs []*Transaction) error {
	log.Std.Notice("block height: %+v", block.Height)
	log.Std.Notice("block hash: %+v", block.Hash)
	for _, tx := range txs {
		log.Std.Notice("tx: %+v", tx)
	}
	return nil
}

func TestHCBlockScanner_ScanBlockTaskOrigin(t *testing.T) {

	bs := NewHCBlockScanner(tw)
	o := &HCBlockObserver{}
	bs.AddHCBlockObserver(o)

	bs.Scanning = true
	bs.ScanBlockTaskOrigin()

}

func TestSliceSplit(t *testing.T) {

	var (
		a     = []int{1, 2, 3}
		limit = 5
		b     = make([]int, 0)
		max   = len(a)
		step  = max / limit
	)

	for i := 0; i <= step; i++ {
		begin := i * limit
		end := (i + 1) * limit
		if end > max {
			end = max
		}

		b = a[begin:end]
		fmt.Printf("[%d]: %v \n", i, b)
	}

}
