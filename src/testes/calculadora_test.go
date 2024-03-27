package main

import "testing"

func TestSoma(t *testing.T) {
	teste := Soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := Soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultiplica1(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultiplica2(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 101

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubtrai1(t *testing.T) {
	teste := Subtrai(5, 10)
	resultado := 5

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubtrai2(t *testing.T) {
	teste := Subtrai(5, 10)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivide1(t *testing.T) {
	teste := Divide(5, 10)
	resultado := 2

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivide2(t *testing.T) {
	teste := Divide(5, 10)
	resultado := 3

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
