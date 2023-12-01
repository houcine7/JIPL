package utils

// check if the a given character is letter
func IsLetter(char rune) bool{
	if (char >= 'a' && char<='z') || (char >='A' && char <='Z'){
		return true
	} 
	return false
}

// checks if a given character is digit or not
func IsDigit(char rune) bool{
	if char >='0' && char<='9'{
		return true
	}
	return false
}
