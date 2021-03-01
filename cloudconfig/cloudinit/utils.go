package cloudinit

import (
	"encoding/base64"
	"fmt"

	"github.com/juju/utils/v2"
)

func removeStringFromSlice(slice []string, val string) []string {
	for i, str := range slice {
		if str == val {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func addFileCmds(filename string, data []byte, mode uint, binary bool) []string {
	// Note: recent versions of cloud-init have the "write_files"
	// module, which can write arbitrary files. We currently support
	// 12.04 LTS, which uses an older version of cloud-init without
	// this module.
	// TODO (aznashwan): eagerly await 2017 and to do the right thing here
	p := utils.ShQuote(filename)

	cmds := []string{fmt.Sprintf("install -D -m %o /dev/null %s", mode, p)}
	// Don't use the shell's echo builtin here; the interpretation
	// of escape sequences differs between shells, namely bash and
	// dash. Instead, we use printf (or we could use /bin/echo).
	if binary {
		encoded := base64.StdEncoding.EncodeToString(data)
		cmds = append(cmds, fmt.Sprintf(`printf %%s %s | base64 -d > %s`, encoded, p))
	} else {
		cmds = append(cmds, fmt.Sprintf(`printf '%%s\n' %s > %s`, utils.ShQuote(string(data)), p))
	}

	return cmds
}
