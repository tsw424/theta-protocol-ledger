package sync

import (
	"github.com/thetatoken/ukulele/common"
	"github.com/thetatoken/ukulele/dispatcher"
)

type RequestManager struct{}

func (rm *RequestManager) EnqueueBlocks(endHash common.Bytes) {}

func (rm *RequestManager) handleInvResponse(invResp *dispatcher.InventoryResponse) {}