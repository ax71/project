package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // Import package strconv untuk menggunakan Atoi
	"strings"
)

type student struct {
	name string // Nama mahasiswa
	nim  int    // Nomor Induk Mahasiswa (NIM)
}

func (s *student) readInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Nama: ")
	name, _ := reader.ReadString('\n')
	s.name = strings.TrimSpace(name)

	fmt.Print("Masukan NIM: ")
	// Ganti fmt.Scan menjadi fmt.Scanln untuk menghindari masalah buffer
	_, err := fmt.Scanln(&s.nim)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

type pertanyaan struct {
	soal    string      // Pertanyaan
	jawaban interface{} // Jawaban yang dipilih oleh pengguna
}

func (p *pertanyaan) bacaInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Apa nama lain dari Oksigen?\n")
	fmt.Print("A. O2\nB. Api\nC. Angin\nD. Udara\nE. Hujan\nJawaban Anda: ")

	soal, _ := reader.ReadString('\n')
	p.soal = strings.TrimSpace(soal)

	jawaban, _ := reader.ReadString('\n')
	jawaban = strings.TrimSpace(jawaban)
	if angka, err := strconv.Atoi(jawaban); err == nil {
		p.jawaban = angka
	} else {
		p.jawaban = strings.ToUpper(jawaban)
	}

	switch p.jawaban.(type) {
	case int, string: // Validasi untuk int atau string
		// Tidak melakukan apa-apa, jawaban valid
	default:
		fmt.Println("Pilihan tidak valid. Menggunakan jawaban default A.")
		p.jawaban = "A"
	}
}

func (p *pertanyaan) bacaInputKedua() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Apa warna langit pada siang hari?\n")
	fmt.Print("A. Hitam\nB. Merah\nC. Hijau\nD. Kuning\nE. Biru\nJawaban Anda: ")

	soal, _ := reader.ReadString('\n')
	p.soal = strings.TrimSpace(soal)

	jawaban, _ := reader.ReadString('\n')
	p.jawaban = strings.ToUpper(strings.TrimSpace(jawaban))

	switch p.jawaban {
	case "A", "B", "C", "D", "E":
		// Tidak melakukan apa-apa, jawaban valid
	default:
		fmt.Println("Pilihan tidak valid. Menggunakan jawaban default E.")
		p.jawaban = "E"
	}
}

func (p *pertanyaan) bacaInputKetiga() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Apa warna langit pada malam hari?\n")
	fmt.Print("A. Hitam\nB. Merah\nC. Hijau\nD. Kuning\nE. Biru\nJawaban Anda: ")

	soal, _ := reader.ReadString('\n')
	p.soal = strings.TrimSpace(soal)

	jawaban, _ := reader.ReadString('\n')
	p.jawaban = strings.ToUpper(strings.TrimSpace(jawaban))

	switch p.jawaban {
	case "A", "B", "C", "D", "E":
		// Tidak melakukan apa-apa, jawaban valid
	default:
		fmt.Println("Pilihan tidak valid. Menggunakan jawaban default A.")
		p.jawaban = "A"
	}
}

func (p *pertanyaan) bacaInputKeempat() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(" 1+ 1 ?\n")
	fmt.Print("A. 2\nB. 3\nC. 11\nD. 5\nE. 9\nJawaban Anda: ")

	soal, _ := reader.ReadString('\n')
	p.soal = strings.TrimSpace(soal)

	jawaban, _ := reader.ReadString('\n')
	jawaban = strings.TrimSpace(jawaban)
	if angka, err := strconv.Atoi(jawaban); err == nil {
		p.jawaban = angka
	} else {
		p.jawaban = strings.ToUpper(jawaban)
	}

	switch p.jawaban.(type) {
	case int, string: // Validasi untuk int atau string
		// Tidak melakukan apa-apa, jawaban valid
	default:
		fmt.Println("Pilihan tidak valid. Menggunakan jawaban default A.")
		p.jawaban = "A"
	}
}

func main() {
	var s1 student
	s1.readInput()

	fmt.Println("\nData Mahasiswa:")
	fmt.Println("Nama:", s1.name)
	fmt.Println("NIM:", s1.nim)

	var q1 pertanyaan
	q1.bacaInput()

	fmt.Println("\nPertanyaan 1:")
	fmt.Println("Pertanyaan:", q1.soal)
	fmt.Println("Jawaban yang benar: A. O2")
	fmt.Printf("Jawaban Anda: %v\n", q1.jawaban)

	var q2 pertanyaan
	q2.bacaInputKedua()
	fmt.Println("\nPertanyaan 2:")
	fmt.Println("Pertanyaan:", q2.soal)
	fmt.Println("Jawaban yang benar: E. Biru")
	fmt.Printf("Jawaban Anda: %v\n", q2.jawaban)

	var q3 pertanyaan
	q3.bacaInputKetiga()
	fmt.Println("\nPertanyaan 3:")
	fmt.Println("Pertanyaan:", q3.soal)
	fmt.Println("Jawaban yang benar: A. Hitam")
	fmt.Printf("Jawaban Anda: %v\n", q3.jawaban)

	var q4 pertanyaan
	q4.bacaInputKeempat()
	fmt.Println("\nPertanyaan 4:")
	fmt.Println("Pertanyaan:", q4.soal)
	fmt.Println("Jawaban yang benar: A. 2")
	fmt.Printf("Jawaban Anda: %v\n", q4.jawaban)

	// Hitung jawaban benar
	jawabanBenar := 0
	if q1.jawaban == "A" {
		jawabanBenar++
	}
	if q2.jawaban == "E" {
		jawabanBenar++
	}
	if q3.jawaban == "A" {
		jawabanBenar++
	}
	if q4.jawaban == 2 {
		jawabanBenar++
	}

	// Menampilkan Data pengisi pertanyaan dan nilai jawaban
	fmt.Printf("\nData Pengisi Pertanyaan dan nilai:\n")
	fmt.Printf("Nama Mahasiswa: %s\n", s1.name)
	fmt.Printf("NIM Mahasiswa: %d\n", s1.nim)
	fmt.Printf("Jumlah Jawaban Benar: %d\n", jawabanBenar)
	fmt.Printf("Jumlah Jawaban Salah: %d\n", 4-jawabanBenar)
}
