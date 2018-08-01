package println

import "println/logs"

func Error(i ...interface{}) {
	logs.Error(i)
}

func Errorf(format string, arg ...interface{}) {
	logs.Errorf(format, arg)
}

func Info(i ...interface{}) {
	logs.Info(i)
}

func Infof(format string, arg ...interface{}) {
	logs.Infof(format, arg)
}

func Warn(i ...interface{}) {
	logs.Warn(i)
}

func Warnf(format string, arg ...interface{}) {
	logs.Warnf(format, arg)
}
