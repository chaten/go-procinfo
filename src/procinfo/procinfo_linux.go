package procinfo
import (
	"os"
)
func GetProcessExecutable() (string,error) {
	return os.Readlink("/proc/self/exe")
}
