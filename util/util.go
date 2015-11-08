package util

func NormalizeDir(path string) string {
	if path[len(path)-1] == '/' {
		return path
	} else {
		return path + "/"
	}
}
