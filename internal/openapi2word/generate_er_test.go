package openapi2word_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zj-open-source/cmd/internal/openapi2word"
)

func TestNewEr(t *testing.T) {
	er, err := openapi2word.NewEr([]string{
		"http://srv-analyze.intelliep.d.rktl.work/analyze/er",
	})
	require.NoError(t, err)

	err = er.GenerateDoc()
	require.NoError(t, err)

	er.Document().SaveToFile("er.docx")
}
