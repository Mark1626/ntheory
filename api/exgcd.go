package handler

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

func ExGcd(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}

	aRaw := r.Form.Get("a")
	a, err := strconv.ParseInt(aRaw, 10, 64)
	if err != nil {
		log.Printf("Error parsing limit: %s", err)
		return
	}

	mRaw := r.Form.Get("m")
	m, err := strconv.ParseInt(mRaw, 10, 64)
	if err != nil {
		log.Printf("Error parsing limit: %s", err)
		return
	}

	k := new(big.Int).ModInverse(big.NewInt(a), big.NewInt(m))

	fmt.Fprintf(w, k.String())
}
