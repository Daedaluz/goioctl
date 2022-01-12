# goioctl

Tiny golang ioctl package.

Simply exposes IOCTL syscall function and some functions for calculating
the request number, often found in linux headers as macros.

* IO
* IOR
* IOW
* IOWR
* IOC


* `func Ioctl(fd, request, value uintptr) (error)`
* `func IoctlX(fd, request, value uintptr) (int64, error)`