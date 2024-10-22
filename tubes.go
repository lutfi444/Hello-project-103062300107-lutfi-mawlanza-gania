package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Struktur pertanyaan
type Question struct {
	ID        int
	Content   string
	Choices   [4]string
	Answer    string
	Correct   int
	Incorrect int
}

// Struktur peserta
type Participant struct {
	ID    int
	Name  string
	Score int
}

// Struktur Sistem Kuis
type QuizSystem struct {
	Questions    []Question
	Participants []Participant
	NextQID      int
	NextPID      int
}

// Sistem kuis
func InitializeSystem() *QuizSystem {
	qs := &QuizSystem{
		Questions:    make([]Question, 0),
		Participants: make([]Participant, 0),
		NextQID:      1,
		NextPID:      1,
	}
	qs.LoadSampleQuestions()
	return qs
}

// Masukan 10 soal pertama
func (qs *QuizSystem) LoadSampleQuestions() {
	qs.AddQuestion("Apa ibukota Indonesia?", [4]string{"Jakarta", "Bandung", "Surabaya", "Yogyakarta"}, "Jakarta")
	qs.AddQuestion("Siapa presiden pertama Indonesia?", [4]string{"Sukarno", "Suharto", "Habibie", "Megawati"}, "Sukarno")
	qs.AddQuestion("Berapa jumlah provinsi di Indonesia?", [4]string{"34", "33", "32", "35"}, "34")
	qs.AddQuestion("Gunung tertinggi di Indonesia adalah?", [4]string{"Semeru", "Rinjani", "Jayawijaya", "Slamet"}, "Jayawijaya")
	qs.AddQuestion("Sungai terpanjang di Indonesia adalah?", [4]string{"Mahakam", "Kapuas", "Barito", "Brantas"}, "Kapuas")
	qs.AddQuestion("Apa bahasa resmi di Indonesia?", [4]string{"Jawa", "Sunda", "Melayu", "Indonesia"}, "Indonesia")
	qs.AddQuestion("Dimana Candi Borobudur berada?", [4]string{"Jawa Timur", "Jawa Tengah", "Yogyakarta", "Bali"}, "Jawa Tengah")
	qs.AddQuestion("Danau terbesar di Indonesia adalah?", [4]string{"Danau Toba", "Danau Maninjau", "Danau Singkarak", "Danau Sentani"}, "Danau Toba")
	qs.AddQuestion("Pulau terbesar di Indonesia adalah?", [4]string{"Sumatra", "Kalimantan", "Sulawesi", "Papua"}, "Kalimantan")
	qs.AddQuestion("Pahlawan nasional yang dikenal dengan semboyan 'Merdeka atau Mati' adalah?", [4]string{"Diponegoro", "Sudirman", "Bung Tomo", "Pattimura"}, "Pattimura")
}

// Penambahan pertanyaan dan jawaban
func (qs *QuizSystem) AddQuestion(content string, choices [4]string, answer string) {
	q := Question{
		ID:      qs.NextQID,
		Content: content,
		Choices: choices,
		Answer:  answer,
	}
	qs.Questions = append(qs.Questions, q)
	qs.NextQID++
}

// Mengedit pertanyaan
func (qs *QuizSystem) EditQuestion(id int, content string, choices [4]string, answer string) bool {
	for i := range qs.Questions {
		if qs.Questions[i].ID == id {
			qs.Questions[i].Content = content
			qs.Questions[i].Choices = choices
			qs.Questions[i].Answer = answer
			return true
		}
	}
	return false
}

// Menghapus pertanyaan
func (qs *QuizSystem) DeleteQuestion(id int) bool {
	for i := range qs.Questions {
		if qs.Questions[i].ID == id {
			qs.Questions = append(qs.Questions[:i], qs.Questions[i+1:]...)
			return true
		}
	}
	return false
}

// Mendaftarkan peserta baru
func (qs *QuizSystem) RegisterParticipant(name string) {
	p := Participant{
		ID:   qs.NextPID,
		Name: name,
	}
	qs.Participants = append(qs.Participants, p)
	qs.NextPID++
}

// Peserta mengikuti kuis
func (qs *QuizSystem) TakeQuiz(participantID int, numQuestions int) {
	participant := qs.GetParticipant(participantID)
	if participant == nil {
		fmt.Println("Peserta tidak ditemukan!")
		return
	}

	questions := qs.GetRandomQuestions(numQuestions)
	score := 0

	for _, q := range questions {
		fmt.Println(q.Content)
		for i, choice := range q.Choices {
			fmt.Printf("%d. %s\n", i+1, choice)
		}

		var answer int
		fmt.Print("Jawaban: ")
		fmt.Scan(&answer)

		if q.Choices[answer-1] == q.Answer {
			score++
			q.Correct++
		} else {
			q.Incorrect++
		}
	}

	participant.Score = score
	fmt.Printf("Skor Akhir: %d/%d\n", score, numQuestions)
}

// ID peserta
func (qs *QuizSystem) GetParticipant(id int) *Participant {
	for i := range qs.Participants {
		if qs.Participants[i].ID == id {
			return &qs.Participants[i]
		}
	}
	return nil
}

// Mengacak pertanyaan
func (qs *QuizSystem) GetRandomQuestions(num int) []Question {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(qs.Questions), func(i, j int) {
		qs.Questions[i], qs.Questions[j] = qs.Questions[j], qs.Questions[i]
	})
	if num > len(qs.Questions) {
		num = len(qs.Questions)
	}
	return qs.Questions[:num]
}

// Insertion sort for participants
func (qs *QuizSystem) InsertionSortParticipants(ascending bool) {
	for i := 1; i < len(qs.Participants); i++ {
		key := qs.Participants[i]
		j := i - 1

		if ascending {
			for j >= 0 && qs.Participants[j].Score > key.Score {
				qs.Participants[j+1] = qs.Participants[j]
				j = j - 1
			}
		} else {
			for j >= 0 && qs.Participants[j].Score < key.Score {
				qs.Participants[j+1] = qs.Participants[j]
				j = j - 1
			}
		}

		qs.Participants[j+1] = key
	}
}

// Selection sort for participants
func (qs *QuizSystem) SelectionSortParticipants(ascending bool) {
	for i := 0; i < len(qs.Participants)-1; i++ {
		idx := i
		for j := i + 1; j < len(qs.Participants); j++ {
			if ascending {
				if qs.Participants[j].Score < qs.Participants[idx].Score {
					idx = j
				}
			} else {
				if qs.Participants[j].Score > qs.Participants[idx].Score {
					idx = j
				}
			}
		}
		qs.Participants[i], qs.Participants[idx] = qs.Participants[idx], qs.Participants[i]
	}
}

// / SortParticipantsByScore sorts participants by score using the specified sorting algorithm
func (qs *QuizSystem) SortParticipantsByScore(algorithm string, ascending bool) {
	switch algorithm {
	case "insertion":
		qs.InsertionSortParticipants(ascending)
	case "selection":
		qs.SelectionSortParticipants(ascending)
	default:
		if ascending {
			sort.Slice(qs.Participants, func(i, j int) bool {
				return qs.Participants[i].Score < qs.Participants[j].Score
			})
		} else {
			sort.Slice(qs.Participants, func(i, j int) bool {
				return qs.Participants[i].Score > qs.Participants[j].Score
			})
		}
	}
}

// main function
func main() {
	qs := InitializeSystem()

	// Menu loop
	for {
		var choice int
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Soal")
		fmt.Println("2. Edit Soal")
		fmt.Println("3. Hapus Soal")
		fmt.Println("4. Daftar Peserta")
		fmt.Println("5. Ikuti Kuis")
		fmt.Println("6. Tampilkan Peserta Berdasarkan Skor")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var content, answer string
			var choices [4]string
			fmt.Print("Masukkan soal: ")
			fmt.Scan(&content)
			for i := 0; i < 4; i++ {
				fmt.Printf("Masukkan pilihan %d: ", i+1)
				fmt.Scan(&choices[i])
			}
			fmt.Print("Masukkan jawaban: ")
			fmt.Scan(&answer)
			qs.AddQuestion(content, choices, answer)
		case 2:
			var id int
			var content, answer string
			var choices [4]string
			fmt.Print("Masukkan ID soal yang akan diedit: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan soal baru: ")
			fmt.Scan(&content)
			for i := 0; i < 4; i++ {
				fmt.Printf("Masukkan pilihan %d: ", i+1)
				fmt.Scan(&choices[i])
			}
			fmt.Print("Masukkan jawaban baru: ")
			fmt.Scan(&answer)
			if !qs.EditQuestion(id, content, choices, answer) {
				fmt.Println("Soal tidak ditemukan!")
			}
		case 3:
			var id int
			fmt.Print("Masukkan ID soal yang akan dihapus: ")
			fmt.Scan(&id)
			if !qs.DeleteQuestion(id) {
				fmt.Println("Soal tidak ditemukan!")
			}
		case 4:
			var name string
			fmt.Print("Masukkan nama peserta: ")
			fmt.Scan(&name)
			qs.RegisterParticipant(name)
		case 5:
			var id, numQuestions int
			fmt.Print("Masukkan ID peserta: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan jumlah soal: ")
			fmt.Scan(&numQuestions)
			qs.TakeQuiz(id, numQuestions)
		case 6:
			var order int
			var algorithm string
			fmt.Print("Urutkan secara (1) Ascending atau (2) Descending: ")
			fmt.Scan(&order)
			fmt.Print("Pilih algoritma sorting (default, insertion, selection): ")
			fmt.Scan(&algorithm)
			qs.SortParticipantsByScore(algorithm, order == 1)
			for _, p := range qs.Participants {
				fmt.Printf("ID: %d, Nama: %s, Skor: %d\n", p.ID, p.Name, p.Score)
			}
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
