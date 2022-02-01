package chain

import "testing"

func Test_nivelUnoStep_Next(t *testing.T) {
	type args struct {
		user string
		pass string
	}
	type fields struct {
		chain Chain
	}

	chain := NewNivelUnoStepErrMock()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ValidatePass success",
			fields:  fields{chain: chain},
			args:    args{user: "COORDINADOR", pass: "Password#2021111111"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nus := &nivelUnoStep{
				next: tt.fields.chain,
			}
			if err := nus.Next(tt.args.user, tt.args.pass); (err != nil) != tt.wantErr {
				t.Errorf("nivelUnoStep.Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
