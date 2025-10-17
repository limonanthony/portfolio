package logger

const (
	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"

	colorBrightRed = "\033[91m"
)

var levelColors = map[Level]string{
	LevelDebug:   colorBlue,   // conventional: debug → blue
	LevelInfo:    colorCyan,   // conventional: info → cyan
	LevelSuccess: colorGreen,  // conventional: success → green
	LevelWarn:    colorYellow, // conventional: warn → yellow
	LevelError:   colorRed,    // conventional: error → red
	LevelPanic:   colorBrightRed,
}
