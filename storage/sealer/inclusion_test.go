package sealer

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/go-commp-utils/v2"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/require"
)

func TestGetSubpieceWithProof(t *testing.T) {
	bz := make([]byte, 508)
	for i := 0; i < len(bz); i++ {
		bz[i] = byte(i % 256)
	}

	pieceCid, err := commp.GeneratePieceCIDFromFile(0, bytes.NewReader(bz), abi.UnpaddedPieceSize(len(bz)))
	require.NoError(t, err)

	subPieces, proof, err := GetSubpiecesWithProof(bz, 0, 3)
	require.NoError(t, err)

	require.NoError(t, VerifySubpiecesProof(pieceCid, subPieces, proof))

	// Flip a bit, proof should now fail
	subPieces[0][0] ^= 1

	// TODO: this doesn't currently pass, because of a quirk in the merkletree library that just uses
	// cached hashes. We need to either use a different library or understand how this is supposed to work.
	require.Error(t, VerifySubpiecesProof(pieceCid, subPieces, proof))
}
