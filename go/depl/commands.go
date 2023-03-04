package depl

var Commands = map[string]string{
	"buildLib":       buildLib,
	"copyLibs":       copyLibs,
	"build":          build,
	"test":           test,
	"push":           push,
	"remote":         remote,
	"pull":           pull,
	"deployCopyLibs": deployCopyLibs,
	"shadowJar":      shadowJar,
	"dockerUp":       dockerUp,
	"dockerUpDev":    dockerUpDev,
	"deploy":         deploy,
	"catLog":         catLog,
	"catEvent":       catEvent}
