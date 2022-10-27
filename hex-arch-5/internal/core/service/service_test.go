package patientservice_test

import (
	"reflect"
	"testing"

	"github.com/devpablocristo/golang/hex-arch-5/internal/core/domain"
	"github.com/devpablocristo/golang/hex-arch-5/internal/core/service/ports"
)

func TestPatientService_GetPatient(t *testing.T) {
	type fields struct {
		patientRepository ports.PatientRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Patient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PatientService{
				patientRepository: tt.fields.patientRepository,
			}
			got, err := ps.GetPatient(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatientService.GetPatient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatientService.GetPatient() = %v, want %v", got, tt.want)
			}
		})
	}
}
