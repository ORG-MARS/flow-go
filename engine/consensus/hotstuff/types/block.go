package types

import (
	"bytes"
	"crypto/sha256"
	"time"

	"github.com/dapperlabs/flow-go/model/flow"
)

type Block struct {
	// specified
	View        uint64
	QC          *QuorumCertificate
	PayloadHash []byte

	// configed
	Height  uint64
	ChainID string

	// autogenerated
	Timestamp time.Time
}

func (b Block) BlockMRH() flow.Identifier {
	data := bytes.Join(
		[][]byte{
			b.QC.BytesForSig(),
			make([]byte, b.View),
			b.PayloadHash,
			make([]byte, b.Height),
			[]byte(b.ChainID),
			//	TODO: put Timestamp here
		},
		[]byte{},
	)

	return sha256.Sum256(data)
}

func NewBlock(view uint64, qc *QuorumCertificate, payloadHash []byte, height uint64, chainID string) *Block {

	t := time.Now()

	return &Block{
		View:        view,
		QC:          qc,
		PayloadHash: payloadHash,
		Height:      height,
		ChainID:     chainID,
		Timestamp:   t,
	}
}
