Logger
======

Yet another logger in go.

This is a simple logger with specific log levels so that you can simply switch log level from a configuration file.
There are 5 levels in order of increasing verboseness:

* Error
* Warning
* Info
* Debug
* Trace


You must initialize it with a specified log level and a path:

	logger.Initialize(logger.LogInfo,"/path/tofile")
	
this will then send to the specified file all messages of level greater or equals to LogInfo.

Doing:

	logger.Error("This is error number %d", 1234)
	logger.Warning("This is a warning")
	logger.Info("This is an info message")
	logger.Debug("This is a debug message")
	logger.Trace("This is a very low level trace")
	
will output the messages to the file.

