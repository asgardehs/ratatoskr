package pip

import (
	"fmt"

	"github.com/asgardehs/ratatoskr/embed_util"
	"github.com/asgardehs/ratatoskr/pip/internal/data"
)

func NewPipLib(name string) (*embed_util.EmbeddedFiles, error) {
	return embed_util.NewEmbeddedFiles(data.Data, fmt.Sprintf("pip-%s", name))
}
