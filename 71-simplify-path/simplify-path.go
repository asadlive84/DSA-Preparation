func simplifyPath(path string) string {
	stack := []string{}
	pathList := strings.Split(path, "/")

	for _, data := range pathList {
		//fmt.Println("data==> ", data)
		if data == "" || data == "." {
			continue
		} else if data == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				continue
			}
		} else {
			stack = append(stack, data)
		}
	}
	//fmt.Println("stack ", stack)
	return "/" + strings.Join(stack, "/")
}