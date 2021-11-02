// Copyright 2021 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package telemetry

import (
	"github.com/ChainSafe/gossamer/lib/common"
)

//nolint
// notifyFinalizedTM holds `notify.finalized` telemetry message, which is
// supposed to be send when a new block gets finalized.
type notifyFinalizedTM struct {
	Best common.Hash `json:"best"`
	// Height is same as block.Header.Number
	Height string `json:"height"`
}

// NewNotifyFinalizedTM gets a new NotifyFinalizedTM struct.
func NewNotifyFinalizedTM(best common.Hash, height string) Message {
	return &notifyFinalizedTM{
		Best:   best,
		Height: height,
	}
}

func (notifyFinalizedTM) messageType() string {
	return notifyFinalizedMsg
}