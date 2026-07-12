package main

import "fmt"

type BankError struct {
	KodeStatus   string
	PesanNasabah string
}

func main() {

	if e := transferDana(100000, 100000); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("success")
	}

}

func (b *BankError) Error() string {
	return fmt.Sprintf("kodeStatus = %v\npesanNasabah %v\n", b.KodeStatus, b.PesanNasabah)
}

func transferDana(saldo, jumlahTarik int) error {

	if jumlahTarik > saldo {
		return &BankError{KodeStatus: "INSUFFICIENT_FUNDS", PesanNasabah: "Maaf, saldo tabungan Anda tidak mencukupi untuk melakukan transaksi ini, Jun!"}
	}

	return nil
}
