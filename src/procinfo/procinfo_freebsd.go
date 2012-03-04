package procinfo

//TODO: Use sysctl(CTL_KERN,KERN_PROC,KERN_PROC_PATHNAME,-1)
func GetProcessExecutable() (string,error) {
	panic("Not yet implemented")
}
