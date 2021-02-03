package vigenere

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{plaintext: "hello world this is a ciphered message", key: "cryptii"}, want: "JVJAH EWTCB IAQA KJ Y RBXPGICS FMAURET"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.args.plaintext, tt.args.key); got != tt.want {
				t.Errorf("Encrypt(%v, %v) = %v, want %v", tt.args.plaintext, tt.args.key, got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		ciphertext string
		key        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{ciphertext: "JVJAH EWTCB IAQA KJ Y RBXPGICS FMAURET", key: "cryptii"}, want: "HELLO WORLD THIS IS A CIPHERED MESSAGE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.ciphertext, tt.args.key); got != tt.want {
				t.Errorf("Decrypt(%v, %v) = %v, want %v", tt.args.ciphertext, tt.args.key, got, tt.want)
			}
		})
	}
}
