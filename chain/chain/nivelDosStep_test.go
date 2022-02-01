package chain

import "testing"

func Test_nivelDosStep_Next(t *testing.T) {
	type args struct {
		user string
		pass string
	}

	type fields struct {
		chain Chain
	}

	chain := NewNivelDosStepMock()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ValidatePass Error min mayus",
			fields:  fields{chain: chain},
			args:    args{user: "COORDINADOR", pass: "paaaaaaaassssss"},
			wantErr: true,
		},
		{
			name:    "ValidatePass Error min mayus",
			fields:  fields{chain: chain},
			args:    args{user: "coordinador", pass: "paaaaaaaassssss"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nds := &nivelDosStep{
				next: tt.fields.chain,
			}
			if err := nds.Next(tt.args.user, tt.args.pass); (err != nil) != tt.wantErr {
				t.Errorf("nivelDosStep.Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
