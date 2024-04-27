package sysInfo

func ScanStatNet() {
	AsyncExecuteCommand("/bin/sh", "-c", "sudo tcpdump", "-ntq", "-i", "any", "-Q", "inout", "-l")
}
