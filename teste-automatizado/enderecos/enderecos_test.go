package enderecos

import "testing"

// Teste Unitários
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado, eu esperava %s e veio %s",
			tipoDeEnderecoRecebido, tipoDeEnderecoEsperado)
	}
}
