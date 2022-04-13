package openapi2word_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zj-open-source/cmd/internal/openapi2word"
)

func TestOpenApi_GenerateDoc(t *testing.T) {
	gen := openapi2word.NewGenerateOpenAPIDoc("user-center", &url.URL{
		Scheme: "http",
		Path:   "srv-user-center/user-center",
	}, 3)
	gen.Load()
	err := gen.GenerateClientOpenAPIDoc("./user-center.docx")
	require.NoError(t, err)
}
