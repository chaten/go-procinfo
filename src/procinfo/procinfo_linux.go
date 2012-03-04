/*
  © 2012 Michael Chaten
  This Source Code Form is subject to the terms of the Mozilla Public
  License, v. 2.0. If a copy of the MPL was not distributed with this
  file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package procinfo
import (
	"os"
)
func GetProcessExecutable() (string,error) {
	return os.Readlink("/proc/self/exe")
}
