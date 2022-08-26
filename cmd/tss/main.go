package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
)

// GenerateKeyPair generates a new key pair
func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privkey, &privkey.PublicKey, nil
}

func main() {
	// sign part
	privateKey, publicKey, err := generateKeyPair(1024)

	N := "0xb793f2f926170fad768f8b1a5769a2243b4cdcac4780194f59b39e1a2abc3bb8ea42db495d17bec7f7072a11ed4fa510e75a7886a5db6f71b7afca0090ca079889d18af0669829ed29a8e21d0c09bd19caaf2fe2cc8121bfc5687ac6698e3022f468a481426486cad263be1a119491e034a6e1ab78f19c066d4145a50f9ecff7"
	fmt.Println("len", len(N))

	modulus := publicKey.N.Bytes()
	m := fmt.Sprintf("0x%x", modulus)
	fmt.Println(len(m), m)

	e := strconv.FormatInt(int64(publicKey.E), 16)
	paddedE := fmt.Sprintf("0x%0256s", e)
	fmt.Println(len(paddedE), paddedE)

	if err != nil {
		fmt.Println("could not generate keypair: ", err.Error())
	}

	data := []byte("hello world")

	digest := sha256.Sum256(data)

	signature, signErr := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest[:])

	if signErr != nil {
		fmt.Println("Could not sign message:", signErr.Error())
	}

	// just to check that we can survive to and from b64
	b64sig := base64.StdEncoding.EncodeToString(signature)

	decodedSignature, _ := base64.StdEncoding.DecodeString(b64sig)

	fmt.Println("decodedSignature", b64sig)

	// verify part

	verifyErr := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest[:], decodedSignature)

	if verifyErr != nil {
		fmt.Println("Verification failed: ", verifyErr)
	}
}
