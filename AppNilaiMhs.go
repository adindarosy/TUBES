package main

import "fmt"

import "os"

import "os/exec"

/* DATA KELOMPOK 
   JUDUL   : APLIKASI NILAI MAHASISWA
   ANGGOTA : - MUHAMMAD ALMERIZTAMA FAUZAN (1301190251)
			 - ADINDA PUTRI ROSYADI (1301190307)
*/


const N = 100

type ujian struct {
	quiz int
	uts  int
	uas  int
}

type matakuliah struct {
	namaMK string
	nilai  ujian
}

type arrMK [8]matakuliah

type mahasiswa struct {
	id            string
	nama          string
	jenis_kelamin string
	matkul        arrMK
	sks			  int
	ip			  float64
	ipk			  float64
}

type arrmahasiswa [N]mahasiswa

func clearcmd() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func menu(datamahasiswa *arrmahasiswa, jumlahmahasiswa *int) {
	var n int

	clearcmd()
	fmt.Println("---------MENU UTAMA--------")
	fmt.Println(" 1. tambah mahasiswa		\n 2. edit mahasiswa		\n 3. hapus mahasiswa		\n 4. List mahasiswa		\n 5. Input Nilai		\n 6. Lihat nilai		\n 7. Hapus Nilai")
	fmt.Println("---------------------------")
	fmt.Print("Ketikkan no menu yang akan dipilih : ")
	fmt.Scan(&n)
	for n > 7 || n < 0 {
		fmt.Print("Ketikkan no menu yang akan dipilih : ")
		fmt.Scan(&n)
	}
	if n == 1 {
		 tambahmahasiswa(&*datamahasiswa, &*jumlahmahasiswa)
	} else if n == 2 {
		 editmahasiswa(&*datamahasiswa, *jumlahmahasiswa)
	} else if n == 3 {
		 hapusmahasiswa(&*datamahasiswa, &*jumlahmahasiswa)
	} else if n == 4 {
		 listmahasiswa(*datamahasiswa, *jumlahmahasiswa)
	} else if n == 5 {
		 inputNilai(&*datamahasiswa, *jumlahmahasiswa)
	} else if n == 6 {
		 perhitungannilaimahasiswa(*datamahasiswa, *jumlahmahasiswa)
	} else if n == 7 {
		 hapusNilai(&*datamahasiswa, &*jumlahmahasiswa)
	}

}

func SelSort(datamahasiswa *arrmahasiswa, jumlahmahasiswa int){ //Pengurutan berdasarkan SKS mahasiswa *DESCENDING
	var i, idx_min, pass int
	var temp mahasiswa

	pass = 0
	for pass < jumlahmahasiswa - 1 {
		idx_min = pass
		i = pass + 1
		for i < jumlahmahasiswa {
			if datamahasiswa[idx_min].sks < datamahasiswa[i].sks {
				idx_min = i
			}
			i++
		}
		temp = datamahasiswa[idx_min]
		datamahasiswa[idx_min] = datamahasiswa[pass]
		datamahasiswa[pass] = temp
		pass++
	}
}

func iSort(datamahasiswa *arrmahasiswa, jumlahmahasiswa int){ //pengurutan berdasarkan IPK mahasiswa *DESCENDING
	var pass, i int
	var temp mahasiswa

	pass = 0
	for pass < jumlahmahasiswa - 1 {
		i = pass + 1
		temp = datamahasiswa[i]
		if i > 0 && temp.ipk > datamahasiswa[i-1].ipk {
			datamahasiswa[i] = datamahasiswa[i-1]
			i--
		}
		datamahasiswa[i] = temp
		pass++
	}
}

func seqsearch(datamahasiswa arrmahasiswa, jumlahmahasiswa int, key string) int { //pencarian dengan NIM MAHASISWA
	var i, search int
	search = -1
	for i = 0; i < jumlahmahasiswa; i++ {
		if datamahasiswa[i].id == key {
			search = i
		}
	}

	if search != -1 {
		return search
	} else {
		return -1
	}

}

func search(datamahasiswa arrmahasiswa, jumlahmahasiswa int, key string) int { //pencarian dengan NIM MAHASISWA
	var search int
	search = -1
	for i := 0; i < jumlahmahasiswa; i++ {
		for j := 0 ; j < 8 ; j++ {
			if datamahasiswa[i].matkul[j].namaMK == key {
				search = j
			}
		}
	}	

	if search != -1 {
		return search
	} else {
		return -1
	}

}
func tambahmahasiswa(datamahasiswa *arrmahasiswa, jumlahmahasiswa *int) {
	var nama, jenis_kelamin, nim string

	clearcmd()

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&nama)
	for nama == "" {
		fmt.Scan(&nama)
	}
	fmt.Print("Masukkan jenis kelamin L/P : ")
	fmt.Scan(&jenis_kelamin)
	for jenis_kelamin != "L" && jenis_kelamin != "P" {
		fmt.Scan(&jenis_kelamin)
	}
	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)
	for nama == "" {
		fmt.Scan(&nim)
	}
	datamahasiswa[*jumlahmahasiswa].nama = nama
	datamahasiswa[*jumlahmahasiswa].jenis_kelamin = jenis_kelamin
	datamahasiswa[*jumlahmahasiswa].id = nim
	*jumlahmahasiswa++

}

func editmahasiswa(datamahasiswa *arrmahasiswa, jumlahmahasiswa int) {
	var hasil, N int
	var data arrmahasiswa
	var namabaru, kelamin_baru, id string

	clearcmd()

	fmt.Print("Cari NIM mahasiswa : ")
	fmt.Scan(&id)

	data = *datamahasiswa
	N = jumlahmahasiswa
	hasil = seqsearch(data, N, id)
	if hasil != -1 {
		fmt.Print("NIM DITEMUKAN", "\nMasukkan nama baru : ")
		fmt.Scan(&namabaru)
		fmt.Print("Konfirmasi jenis kelamin : ")
		fmt.Scan(&kelamin_baru)
		datamahasiswa[hasil].nama = namabaru
		datamahasiswa[hasil].jenis_kelamin = kelamin_baru
	} else {
		fmt.Println("NIM tidak ditemukan")
	}
}

func hapusmahasiswa(datamahasiswa *arrmahasiswa, jumlahmahasiswa *int) {
	var id string
	var data arrmahasiswa
	var N, hasil int
	clearcmd()
	fmt.Print("Masukkan NIM yang ingin dihapus :")
	fmt.Scan(&id)

	data = *datamahasiswa
	N = *jumlahmahasiswa

	hasil = seqsearch(data, N, id)
	if hasil != -1 {
		datamahasiswa[hasil].nama = ""
		datamahasiswa[hasil].jenis_kelamin = ""
	} else {
		fmt.Println("NIM tidak ditemukan")
	}
}

func listmahasiswa(datamahasiswa arrmahasiswa, jumlahmahasiswa int){ //Berdasarkan NIM atau MATA KULIAH yang diambil
	var data arrmahasiswa
	var N, sign int
	var back string

	clearcmd()

	data = datamahasiswa
	N = jumlahmahasiswa

	fmt.Print("Berdasarkan : \n1. IPK \n2. SKS ", "\nPilihan : ")
	fmt.Scan(&sign)
	if sign == 1 {
		iSort(&data, N)
		fmt.Println("---DAFTAR MAHASISWA BERDASARKAN IPK TERTINGGI---", "\n", "\nFormat : NAMA - IPK")
		for i := 0 ; i < N ; i++ {
			fmt.Println(data[i].nama, data[i].ipk)
		}
	}else if sign ==2 {
		SelSort(&data, N)
		fmt.Println("---DAFTAR MAHASISWA BERDASARKAN SKS TERBANYAK---", "\n", "\nFormat : NAMA - SKS")
		for i := 0 ; i < N ; i++ {
			fmt.Println(data[i].nama, data[i].sks)
		}
	}

	fmt.Println("Ketik back untuk kembali ke menu")
	fmt.Scan(&back)
	for back != "back" {
		fmt.Scan(&back)
	}
}

func inputNilai(datamahasiswa *arrmahasiswa, jumlahmahasiswa int) { // Mengisi sks, mata kuliah, dan nilai mahasiswa
	var x, N, hasil, nilai, sks, rata, jumsks int
	var mkuliah, id, back string
	var indeks, ip, ipk, jumip float64
	var data arrmahasiswa

	clearcmd()

	fmt.Print("Masukkan id yang ingin diinputkan nilai : ")
	fmt.Scan(&id)

	data = *datamahasiswa
	N = jumlahmahasiswa

	hasil = seqsearch(data, N, id)
	fmt.Print("Jumlah matkul yang akan diinput nilai : ")
	fmt.Scan(&x)
	fmt.Println("1.Kalkulus \n2.Dasar Algortitma Pemrograman \n3.Logika Matematika \n4.Literasi TIK \n5.Pengantar Teknik Informatika \n6.B.Inggris \n7.Pembentukan Karakter")
	for i := 0 ; i < x ; i++ {
		fmt.Print("Pilih Mata Kuliah : ")
		fmt.Scan(&mkuliah)
		for mkuliah == "" {
			fmt.Print("Pilih Mata Kuliah : ")
			fmt.Scan(&mkuliah)
		}
		datamahasiswa[hasil].matkul[i].namaMK = mkuliah
		
		fmt.Print("Masukkan Jumlah SKS : ")
		fmt.Scan(&sks)
		for sks > 8 || sks < 0 {
			fmt.Scan(&sks)
		}

		fmt.Print("Masukkan nilai ", mkuliah ,"\nQuiz (25%) : ")
		fmt.Scan(&nilai)	
		for nilai > 100 || nilai < 0 {
			fmt.Scan(&nilai)
		}
		datamahasiswa[hasil].matkul[i].nilai.quiz = nilai

		fmt.Print("UTS (35%) : ")
		fmt.Scan(&nilai)
		for nilai > 100 || nilai < 0 {
			fmt.Scan(&nilai)
		}
		datamahasiswa[hasil].matkul[i].nilai.uts = nilai

		fmt.Print("UAS (40%) : ")
		fmt.Scan(&nilai)
		for nilai > 100 || nilai < 0 {
			fmt.Scan(&nilai)
		}
		datamahasiswa[hasil].matkul[i].nilai.uas = nilai
		rata = ((datamahasiswa[hasil].matkul[i].nilai.quiz * 25/100) + (datamahasiswa[hasil].matkul[i].nilai.uts * 35/100) + (datamahasiswa[hasil].matkul[i].nilai.uas * 40/100))
		if rata > 80 {
			indeks = 4
			fmt.Println("Grade : A")
		} else if rata <= 80 && rata > 70 {
			indeks = 3.5
			fmt.Println("Grade : AB")
		} else if rata <= 70 && rata > 60 {
			indeks = 3
			fmt.Println("Grade : B")
		} else if rata <= 60 && rata > 50 {
			indeks = 2
			fmt.Println("Grade : C")
		} else if rata <= 50 && rata > 45 {
			indeks = 1
			fmt.Println("Grade : D")
		} else if rata <= 45 {
			indeks = 0
			fmt.Println("Grade : E")
		}
		ip = indeks*float64(sks)
		jumsks = jumsks + sks
		jumip = jumip + ip
	}
	datamahasiswa[hasil].sks = jumsks
	datamahasiswa[hasil].ip = jumip
	ipk = jumip / float64(jumsks)
	datamahasiswa[hasil].ipk=ipk
	fmt.Println("IPK MAHASISWA : ", ipk)
	fmt.Println("Ketik back untuk kembali ke menu")
	fmt.Scan(&back)
	for back != "back" {
		fmt.Scan(&back)
	}
}

func hapusNilai(datamahasiswa *arrmahasiswa, jumlahmahasiswa *int) {
	var id string
	var data arrmahasiswa
	var N, hasil int
	clearcmd()
	fmt.Print("Masukkan NIM yang ingin akan dihapus Nilainya :")
	fmt.Scan(&id)

	data = *datamahasiswa
	N = *jumlahmahasiswa

	hasil = seqsearch(data, N, id)
	if hasil != -1 {
		for i:=0 ; i < 8 ; i++ {
		datamahasiswa[hasil].matkul[i].namaMK = ""
		datamahasiswa[hasil].matkul[i].nilai.quiz = 0
		datamahasiswa[hasil].matkul[i].nilai.uts = 0
		datamahasiswa[hasil].matkul[i].nilai.uas = 0
		}
	} else {
		fmt.Println("NIM tidak ditemukan")
	}
}

func perhitungannilaimahasiswa(datamahasiswa arrmahasiswa, jumlahmahasiswa int){ //Berdasarkan NIM atau MATA KULIAH yang diambil
	var data arrmahasiswa
	var N, hasil, sign int
	var id, matkul, back string

	clearcmd()

	fmt.Println("Berdasarkan : \n1. NIM \n2.Mata Kuliah ", "\nPilihan : ")
	fmt.Scan(&sign)
	if sign == 1 {
		fmt.Print("Masukkan NIM : ")
		fmt.Scan(&id)

		data = datamahasiswa
		N = jumlahmahasiswa

		hasil = seqsearch(data, N, id)

		fmt.Println("Nama : ", datamahasiswa[hasil].nama, "\nKelamin : ", datamahasiswa[hasil].jenis_kelamin, "\nNIM : ", datamahasiswa[hasil].id)
		fmt.Println("\n","----TRANSKRIP NILAI----")
			for i := 0 ; i < 8 ; i++ {
				fmt.Println("\n","Nilai ", datamahasiswa[hasil].matkul[i].namaMK)
				fmt.Print("Quiz : ", datamahasiswa[hasil].matkul[i].nilai.quiz, "\nUTS : ", datamahasiswa[hasil].matkul[i].nilai.uts, "\nUAS : ", datamahasiswa[hasil].matkul[i].nilai.uas)
			}
		fmt.Println("\nSKS : ", datamahasiswa[hasil].sks)
	
	} else if sign == 2 {
		fmt.Print("Masukkan Mata Kuliah : ")
		fmt.Scan(&matkul)

		data = datamahasiswa
		N = jumlahmahasiswa

		hasil = search(data, N, matkul)

		fmt.Println("\n","----TRANSKRIP NILAI", matkul, "---")
			for i := 0 ; i < jumlahmahasiswa ; i++ {
				fmt.Println("\nNama : ", datamahasiswa[i].nama, "\nKelamin : ", datamahasiswa[i].jenis_kelamin, "\nNIM : ", datamahasiswa[i].id)
				fmt.Print("Quiz : ", datamahasiswa[i].matkul[hasil].nilai.quiz, "\nUTS : ", datamahasiswa[i].matkul[hasil].nilai.uts, "\nUAS : ", datamahasiswa[i].matkul[hasil].nilai.uas)
			}
	} 
	fmt.Println("\nKetik back untuk kembali ke menu")
	fmt.Scan(&back)
	for back != "back" {
		fmt.Scan(&back)
	}
}



func main(){
	var jumlahmahasiswa int
	var datamahasiswa arrmahasiswa
	var stop bool

	for !stop {
		menu(&datamahasiswa, &jumlahmahasiswa)
	}

}
