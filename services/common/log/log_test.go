package log

import "testing"

func Test_Info(t *testing.T) {
	Info("Info")
}

func Test_Infof(t *testing.T) {
	Infof("Info")
}

func Test_Infoln(t *testing.T) {
	Infoln("Info")
}

func Test_Error(t *testing.T) {
	Error("Error")
}

func Test_Errorf(t *testing.T) {
	Errorf("Error")
}

func Test_Errorln(t *testing.T) {
	Errorln("Error")
}

func Test_Warn(t *testing.T) {
	Warn("Warn")
}

func Test_Warnf(t *testing.T) {
	Warnf("%s", "Warn")
}

func Test_Warnln(t *testing.T) {
	Warnln("Warn")
}
func Test_Debug(t *testing.T) {
	Debug("Debug")
}

func Test_Debugf(t *testing.T) {
	Debugf("Debug")
}

func Test_Debugln(t *testing.T) {
	Debugln("Debug")
}

func Benchmark_Info(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Info(i)
	}
}

func Benchmark_Infof(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Infof("%v", i)
	}
}

func Benchmark_Infoln(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Infoln(i)
	}
}

func Benchmark_Error(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Error(i)
	}
}

func Benchmark_Errorf(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Errorf("%v", i)
	}
}

func Benchmark_Errorln(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Errorln(i)
	}
}
