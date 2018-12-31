package broker

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
)

func (bc BrokerController) Get(ctx context.Context, UserID int64) (*entity.Broker, error) {
	return bc.storage.Get(ctx, UserID)
}
