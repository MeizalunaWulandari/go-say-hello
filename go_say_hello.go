package go_say_hello

func SayHello()string{
	return "Hello"
}

/**
 * Untuk membuat module golang cukup jalankan go mod init namaModule
 * golang akan secara otomatis membuat file bernama go.mod
 * file tersebut hanya berisikan nama module dan versi golang
 * 
 * Untuk rilis versi module cukup dengan memuat tag pada git repositorynya
 * git tag v1.0.0
 * untuk penamaan tag pada git usaha diawali dengan v (v1.0.0)
 * 
 */