package lottery

import (
	"testing"
	"urubu-do-pix/user"
)

func TestParticipateInLottery(t *testing.T) {
	// Criar um usuário de teste
	_, err := user.CreateUser("Teste", "senha123")
	if err != nil {
		t.Fatalf("Erro ao criar usuário de teste: %v", err)
	}

	// Participar na loteria
	result, err := ParticipateInLottery("pix-Teste", 10)
	if err != nil {
		t.Fatalf("Erro ao participar da loteria: %v", err)
	}

	if result.Prize != 0 {
		t.Errorf("Esperado prêmio 0, mas obteve %.2f", result.Prize)
	}
}
