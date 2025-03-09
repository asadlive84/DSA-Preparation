func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func gcdOfStrings(str1 string, str2 string) string {

    if str1+str2 != str2 +str1{
        return ""
    }

	str1Len := len(str1)
	str2Len := len(str2)

	gcd := gcd(str1Len, str2Len)

	return string(str1[:gcd])
}