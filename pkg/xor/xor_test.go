package xor

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{args: args{plaintext: "hello world this is a string", key: 63}, want: "575a5353501f48504d535b1f4b57564c1f564c1f5e1f4c4b4d565158", wantErr: false},
		{name: "check key too large", args: args{plaintext: "hello world this is a ciphered message", key: 257}, want: "", wantErr: true},
		{name: "check key too small", args: args{plaintext: "hello world this is a ciphered message", key: -2}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.plaintext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt(%v, %v) error = %v, wantErr %v", tt.args.plaintext, tt.args.key, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt(%v, %v) = %v, want %v", tt.args.plaintext, tt.args.key, got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		ciphertext string
		key        int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{args: args{ciphertext: "575a5353501f48504d535b1f4b57564c1f564c1f5e1f4c4b4d565158", key: 63}, want: "hello world this is a string", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.ciphertext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt(%v, %v) error = %v, wantErr %v", tt.args.ciphertext, tt.args.key, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt(%v, %v) = %v, want %v", tt.args.ciphertext, tt.args.key, got, tt.want)
			}
		})
	}
}
