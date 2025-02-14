package utils

import ()

func AuthenticateCode(rightCode, code string) bool {
	if rightCode == code {
		return true
	}
	return false
}
