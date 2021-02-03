package keyedcaesar

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       string
		shift     int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{args: args{plaintext: "hello world this is a ciphered message", shift: 7, key: "thisakey"}, want: "LFPPU SUXPD TLMZ MZ Y CMVLFXFD QFZZYJF", wantErr: false},
		{name: "check duplicated letters", args: args{plaintext: "hello world this is a ciphered message", shift: 7, key: "thisaakey"}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.plaintext, tt.args.key, tt.args.shift)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt(%v, %v, %v) error = %v, wantErr %v", tt.args.plaintext, tt.args.key, tt.args.shift, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt(%v, %v, %v) = %v, want %v", tt.args.plaintext, tt.args.key, tt.args.shift, got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		ciphertext string
		key        string
		shift      int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{args: args{ciphertext: "LFPPU SUXPD TLMZ MZ Y CMVLFXFD QFZZYJF", shift: 7, key: "thisakey"}, want: "HELLO WORLD THIS IS A CIPHERED MESSAGE", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.ciphertext, tt.args.key, tt.args.shift)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt(%v, %v, %v) error = %v, wantErr %v", tt.args.ciphertext, tt.args.key, tt.args.shift, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt(%v, %v, %v) = %v, want %v", tt.args.ciphertext, tt.args.key, tt.args.shift, got, tt.want)
			}
		})
	}
}
