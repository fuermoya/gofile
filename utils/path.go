package utils

func JoinExt(p, ext string) string {
	for i := len(p) - 1; i >= 0 && p[i] != '/'; i-- {
		if p[i] == '.' {
			p = p[:i]
			break
		}
	}
	return p + ext
}
