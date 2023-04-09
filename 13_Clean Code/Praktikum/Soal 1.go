package main

type user struct {
	id       int
	username int
	password int
}

// Penulisan dengan gabungan dua kata harus menggunakan camelCase atau PascalCase untuk memudahkan pembaca
type userservice struct {
	// Penamaan property struct sulit untuk dipahami, t digunakan untuk apa?
	t []user
}

// Penulisan struct dan function harus menggunakan camelCase atau PascalCase untuk memudahkan pembaca
func (u userservice) getallusers() []user {
	// Penggunaan property t sulit dipahami maknanya
	return u.t
}

// Penulisan struct dan function harus menggunakan camelCase atau PascalCase untuk memudahkan pembaca
func (u userservice) getuserbyid(id int) user {
	// Penggunaan t dan r sulit untuk dipahami oleh pembaca penggunaannya
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
