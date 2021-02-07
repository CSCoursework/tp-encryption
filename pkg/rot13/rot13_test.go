package rot13

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{plaintext: "hello world this is a ciphered message"}, want: "URYYB JBEYQ GUVF VF N PVCURERQ ZRFFNTR"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.args.plaintext); got != tt.want {
				t.Errorf("Encrypt(%v) = %v, want %v", tt.args.plaintext, got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		ciphertext string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{ciphertext: "URYYB JBEYQ GUVF VF N PVCURERQ ZRFFNTR"}, want: "HELLO WORLD THIS IS A CIPHERED MESSAGE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.ciphertext); got != tt.want {
				t.Errorf("Decrypt(%v) = %v, want %v", tt.args.ciphertext, got, tt.want)
			}
		})
	}
}
