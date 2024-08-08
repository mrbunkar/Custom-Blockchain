package core

import "io"

// Transaction needs to be signed
type Transaction struct {
	Data []byte
}

func (t *Transaction) NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}
