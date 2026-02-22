package mathutils

// Add is a public function because it starts with a Capital letter!
func Add(a, b int) int {
	return a + b
}

// multiply is private! It can ONLY be used inside this `mathutils` folder.
func multiply(a, b int) int {
	return a * b
}
