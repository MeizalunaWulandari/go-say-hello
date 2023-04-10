package go_say_hello

func SayHello(name string)string{
	return "Hello " + name
}

/**
 * MEMBUAT MODULE
 * Untuk membuat module golang cukup jalankan go mod init namaModule
 * golang akan secara otomatis membuat file bernama go.mod
 * file tersebut hanya berisikan nama module dan versi golang
 * 
 * RILIS MODULE
 * Untuk rilis versi module cukup dengan memuat tag pada git repositorynya
 * git tag v1.0.0
 * untuk penamaan tag pada git usaha diawali dengan v (v1.0.0)
 * 
 * UPGRADE MODULE
 * Untuk upgrade module cukup dengan membuat tag baru pada git
 * disarankan jika melakukan upgrade sekecil apapun versinya juga diubah 
 * hal ini untuk menghindari module tidak terdownload pada project yang menggunakannya 
 * 
 * MAJOR UPGRADE
 * Major update biasanya terjadi dikarenakan ada perubahan pada isi kode program sehingga membuatnya tidak backward compatible
 * Sebaiknya hal ini sebisa mungkin untuk dihindari
 * Namun jika tidak bisa dihindari, stategi terbaik adalah merubah nama module
 */