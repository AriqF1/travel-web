package booking

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ValidateSeatNumber memeriksa apakah format nomor kursi valid dan berada dalam 
// jangkauan kuota yang tersedia.
// Contoh format yang valid: "A1", "A12" (case-insensitive).
func ValidateSeatNumber(seat string, seatCount int) error {

    // normalisasi teks untuk mencegah error akibat spasi atau huruf kecil
    seat = strings.ToUpper(strings.TrimSpace(seat))

    // Semua nomor kursi harus diawali dengan blok/zona "A"
    if !strings.HasPrefix(seat, "A") {
        return errors.New("invalid seat format")
    }

    // Ekstrak angka setelah huruf "A" untuk divalidasi sebagai integer
    number, err := strconv.Atoi(
        strings.TrimPrefix(seat, "A"),
    )
    if err != nil {
        return errors.New("invalid seat format")
    }

    // Memastikan nomor kursi tidak bernilai 0/negatif, dan tidak melebihi total kursi yang tersedia
    if number < 1 || number > seatCount {
        return fmt.Errorf(
            "seat must be between A1 and A%d",
            seatCount,
        )
    }

    return nil
}