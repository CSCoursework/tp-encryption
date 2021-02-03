package caesar

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
		shift     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{plaintext: "hello world this is a ciphered message", shift: 7}, want: "OLSSV DVYSK AOPZ PZ H JPWOLYLK TLZZHNL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.args.plaintext, tt.args.shift); got != tt.want {
				t.Errorf("Encrypt(%v, %v) = %v, want %v", tt.args.plaintext, tt.args.shift, got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		ciphertext string
		shift      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{ciphertext: "OLSSV DVYSK AOPZ PZ H JPWOLYLK TLZZHNL", shift: 7}, want: "HELLO WORLD THIS IS A CIPHERED MESSAGE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.ciphertext, tt.args.shift); got != tt.want {
				t.Errorf("Decrypt(%v, %v) = %v, want %v", tt.args.ciphertext, tt.args.shift, got, tt.want)
			}
		})
	}
}
