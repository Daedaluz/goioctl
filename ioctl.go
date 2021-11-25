package ioctl

import "syscall"

const (
	iocNrBits   = 8
	iocTypeBits = 8
	iocSizeBits = 14
	iocDirBits  = 2

	iocNrShift   = 0
	iocTypeShift = iocNrShift + iocNrBits
	iocSizeShift = iocTypeShift + iocTypeBits
	iocDirShift  = iocSizeShift + iocSizeBits

	iocNone  = 0
	iocWrite = 1
	iocRead  = 2
)

func IO(t, nr uintptr) uintptr {
	return IOC(iocNone, t, nr, 0)
}

func IOR(t, nr, size uintptr) uintptr {
	return IOC(iocRead, t, nr, size)
}

func IOW(t, nr, size uintptr) uintptr {
	return IOC(iocWrite, t, nr, size)
}

func IOWR(t, nr, size uintptr) uintptr {
	return IOC(iocRead|iocWrite, t, nr, size)
}

func IOC(dir, t, nr, size uintptr) uintptr {
	return (dir << iocDirShift) | (t << iocTypeShift) | (nr << iocNrShift) | (size << iocSizeShift)
}

func Ioctl(fd int, request uintptr, value uintptr) error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), request, value)
	if e == syscall.Errno(0) {
		return nil
	}
	return e
}
