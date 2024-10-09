package v2

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PacketReceiptKey returns the store key of under which a packet
// receipt is stored
func PacketReceiptKey(sourceID string, sequence uint64) []byte {
	return []byte(fmt.Sprintf("receipts/channels/%s/sequences/%s", sourceID, sdk.Uint64ToBigEndian(sequence)))
}

// PacketAcknowledgementKey returns the store key of under which a packet acknowledgement is stored.
func PacketAcknowledgementKey(sourceID string, sequence uint64) []byte {
	return []byte(fmt.Sprintf("acks/channels/%s/sequences/%s", sourceID, sdk.Uint64ToBigEndian(sequence)))
}

// PacketCommitmentKey returns the store key of under which a packet commitment is stored.
func PacketCommitmentKey(sourceID string, sequence uint64) []byte {
	return []byte(fmt.Sprintf("commitments/channels/%s/sequences/%s", sourceID, sdk.Uint64ToBigEndian(sequence)))
}

// NextSequenceSendKey returns the store key for the next sequence send of a given sourceID.
func NextSequenceSendKey(sourceID string) []byte {
	return []byte(fmt.Sprintf("nextSequenceSend/%s", sourceID))
}