package godecode2

import "testing"

func TestDecode(t *testing.T) {
	m := make(map[string]string)
	m["hello world"] = "hello world"
	m["南无阿弥陀佛"] = "Nan Wu A Mi Tuo Fo"
	m["Κνωσός"] = "Knosos"
	m["あみだにょらい"] = "amidaniyorai"
	for k, v := range m {
		if Decode(k) != v {
			t.Errorf("[ %s ] Error: [ %s ] != [ %s ]", k, v, Decode(k))
		}
	}
}
func TestInitals(t *testing.T) {
	m := make(map[string]string)
	m["Hello world."] = "Hw"
	m["南无阿弥陀佛"] = "NWAMTF"
	m["Κνωσός"] = "K"
	m["あみだにょらい"] = "a"
	m["小小姑娘\n清早起床\n\r提着花篮\t上市场。"] = "XXGN\nQZQC\n\rTZHL\tSSC"
	for k, v := range m {
		if Initials(k) != v {
			t.Errorf("[ %s ] Error: [ %s ] != [ %s ]", k, v, Initials(k))
		}
	}
}
