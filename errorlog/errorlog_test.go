package errorlog

import (
	"testing"
)

func TestElogInsertStmnt(t *testing.T) {
	l := Elog{}
	message := "fake error in line:32"
	payload := `{"test": 1, "tess2": "ss"}`

	actual := l.insertStmnt(message, payload)
	expected := `INSERT INTO error_logs (repository, error, payload)
		VALUES ('', 'fake error in line:32', '{"test": 1, "tess2": "ss"}')`
	if actual != expected {
		t.Errorf("expected statement: %s, but actual: %s", expected, actual)
	}

	t.Run("when there is escaping character in message", func(t *testing.T) {
		l := Elog{}
		message := "fake error 'c' in line:32"
		payload := `{"test": 1, "tess2": "ss"}`

		actual := l.insertStmnt(message, payload)
		expected := `INSERT INTO error_logs (repository, error, payload)
		VALUES ('', 'fake error \'c\' in line:32', '{"test": 1, "tess2": "ss"}')`
		if actual != expected {
			t.Errorf("expected statement: %s, but actual: %s", expected, actual)
		}
	})
}

func TestElogEnv(t *testing.T) {
	t.Run("when there is no env given", func(t *testing.T) {
		l := Elog{}
		actual := l.isProd()
		expected := false
		if actual != expected {
			t.Errorf("expected statement: %t, but actual: %t", expected, actual)
		}
	})
	t.Run("when production env is given", func(t *testing.T) {
		l := Elog{environment: "production"}
		actual := l.isProd()
		expected := true
		if actual != expected {
			t.Errorf("expected statement: %t, but actual: %t", expected, actual)
		}
	})
}
