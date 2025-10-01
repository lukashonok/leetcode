package main

import (
	"strings"
)

func simplifyPath(path string) string {
	newPath := strings.Replace(path, "//", "/", -1)
	for len(newPath) != len(path) {
		path = newPath
		newPath = strings.Replace(path, "//", "/", -1)
	}
	path = newPath

	parts := strings.Split(path, "/")

	stack := []string{}
	for _, part := range parts {
		switch part {
		case "":
		case ".":
			continue
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}

// func main() {
// 	fmt.Println(simplifyPath("//home/..////.//.../ew///qw/"))
// }
