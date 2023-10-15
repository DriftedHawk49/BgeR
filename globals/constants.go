package globals

type OS_TYPE string
type DE_TYPE string

const (
	OS_LINUX          OS_TYPE = "LINUX"
	OS_WINDOWS        OS_TYPE = "WINDOWS"
	OS_MAC            OS_TYPE = "MAC"
	OS_INVALID        OS_TYPE = "INVALID"
	DE_GNOME          DE_TYPE = "GNOME"
	DE_KDE            DE_TYPE = "KDE"
	DE_XFCE           DE_TYPE = "XFCE"
	DE_INVALID        DE_TYPE = "INVALID"
	DUR_DEFAULT       int     = 60
	TEMP_FAILURE_FILE string  = "bgerFailure.txt"
)
