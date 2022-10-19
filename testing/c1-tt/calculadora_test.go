package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 2
	num2 := 1
	resultadoEsperado := 1
	resultado := Restar(num1, num2)
	if resultado != resultadoEsperado {
		t.Error("el resultado es incorrecto")
	}
}

func TestOrdenar(t *testing.T) {
	nums := []int{1, 3, 2, 6, 4}
	resultadoEsperado := []int{1, 2, 3, 4, 6}
	resultado := Ordenar(nums)
	assert.Equal(t, resultado, resultadoEsperado, "no coinciden")
}

func TestDividir(t *testing.T) {
	num1 := 2
	num2 := 0
	resultadoEsperado := 0
	resultado, err := Dividir(num1, num2)
	assert.Equal(t, resultado, resultadoEsperado)
	assert.NotNil(t, err)
}
