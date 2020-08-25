package grpc

import (
	"context"
	"testing"

	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/converter"
	pb "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/stretchr/testify/assert"
)

func TestServerRender(t *testing.T) {
	lc, wc := converter.NewConverters()
	s := NewServer(lc, wc)
	src := `foo https://google.com/ bar
`
	reply, err := s.Render(context.Background(), &pb.RenderRequest{Src: src})
	assert.NoError(t, err)
	assert.Equal(t, `foo <a href="https://google.com/">https://google.com/</a> bar
`, reply.Html)
}
