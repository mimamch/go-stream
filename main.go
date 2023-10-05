package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/plain")

	send:
		for i := 1; i <= 500; i++ {
			// Membuat teks yang akan dikirim
			text := fmt.Sprintf("Data ke-%d pada %s", i, time.Now().Format("2006-01-02 15:04:05"))

			// Mengirim teks ke klien
			_, err := w.Write([]byte(text))

			if err != nil {
				fmt.Println("Gagal mengirim data:", err)
				return
			}

			// Memastikan data segera dikirim ke klien
			w.(http.Flusher).Flush()
			fmt.Println("sent: ", text)

			select {
			case <-r.Context().Done():
				fmt.Println("Connection closed")
				break send
			default:
			}

			// Menunggu sejenak sebelum mengirim data berikutnya (hanya untuk tujuan ilustrasi)
			time.Sleep(1 * time.Second)
		}
	})

	http.ListenAndServe(":5000", r)
}
