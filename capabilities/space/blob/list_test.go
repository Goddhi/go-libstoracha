package blob_test

import (
	"testing"
	"time"

	"github.com/storacha/go-libstoracha/capabilities/internal/testutil"
	"github.com/storacha/go-libstoracha/capabilities/space/blob"
	"github.com/stretchr/testify/require"
)

func TestRoundTripListCaveats(t *testing.T) {
	cursor := "test-cursor"
	size := uint64(10)
	nb := blob.ListCaveats{
		Cursor: &cursor,
		Size:   &size,
	}

	node, err := nb.ToIPLD()
	require.NoError(t, err)

	rnb, err := blob.ListCaveatsReader.Read(node)
	require.NoError(t, err)
	require.Equal(t, nb, rnb)
}

func TestRoundTripListOk(t *testing.T) {
	cursor := "test-cursor"
	results := []blob.ListBlobItem{
		{
			Blob: blob.Blob{
				Digest: testutil.RandomMultihash(t),
				Size:   uint64(1024),
			},
			InsertedAt: time.Now().Truncate(time.Second),
		},
		{
			Blob: blob.Blob{
				Digest: testutil.RandomMultihash(t),
				Size:   uint64(2048),
			},
			InsertedAt: time.Now().Truncate(time.Second),
		},
	}
	ok := blob.ListOk{
		Cursor:  &cursor,
		Size:    uint64(len(results)),
		Results: results,
	}

	node, err := ok.ToIPLD()
	require.NoError(t, err)

	rok, err := blob.ListOkReader.Read(node)
	require.NoError(t, err)
	require.Equal(t, ok, rok)
}
