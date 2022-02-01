package chain

import "testing"

func TestServiceImplementation_ValidatePassword(t *testing.T) {
	type args struct {
		user string
		pass string
	}

	//estructura para el manejo de dependencias
	type fields struct {
		chain Chain
	}

	chain := NewNivelUnoStepMock()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ValidatePass success",
			fields:  fields{chain: chain},
			args:    args{user: "COORDINADOR", pass: "Password#2021"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s := &ServiceImplementation{
			validateChain: tt.fields.chain,
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := s.ValidatePassword(tt.args.user, tt.args.pass); (err != nil) != tt.wantErr {
				t.Errorf("ServiceImplementation.ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
