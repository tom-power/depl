package depl

var Commands = map[string]string{
	"buildLib":       buildLib,
	"copyLibs":       copyLibs,
	"copyLibsNested": copyLibsNested,
	"test":           test,
	"push":           push,
	"remote":         remote,
	"pull":           pull,
	"deployCopyLibs": deployCopyLibs,
	"shadowJar":      shadowJar,
	"dockerUp":       dockerUp,
	"deploy":         deploy,
	"catLog":         catLog,
	"catEvent":       catEvent}