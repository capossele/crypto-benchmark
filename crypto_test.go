package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/oasisprotocol/ed25519"
	"golang.org/x/crypto/blake2b"
)

type zeroReader struct{}

func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

func BenchmarkECDSAP256KeyGeneration(b *testing.B) {
	p256 := elliptic.P256()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ecdsa.GenerateKey(p256, rand.Reader)
	}
}

func BenchmarkECDSAP256Sign(b *testing.B) {
	p256 := elliptic.P256()
	priv, _ := ecdsa.GenerateKey(p256, rand.Reader)
	msg := "IOTA"
	b.ResetTimer()
	hash := sha256.Sum256([]byte(msg))
	for i := 0; i < b.N; i++ {
		if _, _, err := ecdsa.Sign(rand.Reader, priv, hash[:]); err != nil {
			b.Fatal("Signing failed")
		}
	}
}

func BenchmarkECDSAP256Verify(b *testing.B) {
	b.ResetTimer()
	p256 := elliptic.P256()
	msg := "IOTA"
	hash := sha256.Sum256([]byte(msg))
	priv, _ := ecdsa.GenerateKey(p256, rand.Reader)
	r, s, _ := ecdsa.Sign(rand.Reader, priv, hash[:])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !ecdsa.Verify(&priv.PublicKey, hash[:], r, s) {
			b.Fatal("Verification failed")
		}
	}
}

func BenchmarkEd25519KeyGeneration(b *testing.B) {
	var zero zeroReader
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := ed25519.GenerateKey(zero); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEd25519Sign(b *testing.B) {
	var zero zeroReader
	_, priv, err := ed25519.GenerateKey(zero)
	if err != nil {
		b.Fatal(err)
	}
	msg := []byte("IOTA")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ed25519.Sign(priv, msg)
	}
}

func BenchmarkEd25519Verify(b *testing.B) {
	var zero zeroReader
	pub, priv, err := ed25519.GenerateKey(zero)
	if err != nil {
		b.Fatal(err)
	}
	msg := []byte("IOTA")
	signature := ed25519.Sign(priv, msg)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !ed25519.Verify(pub, msg, signature) {
			b.Fatal("Verification failed")
		}
	}
}

func BenchmarkSHA256(b *testing.B) {
	msg := "IOTA"
	b.ResetTimer()
	hash := sha256.Sum256([]byte(msg))
	for i := 0; i < b.N; i++ {
		hash = sha256.Sum256(hash[:])
	}
}

func BenchmarkBlake2(b *testing.B) {
	msg := "IOTA"
	b.ResetTimer()
	hash := blake2b.Sum256([]byte(msg))
	for i := 0; i < b.N; i++ {
		hash = blake2b.Sum256(hash[:])
	}

}
