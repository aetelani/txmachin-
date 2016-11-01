package main
/*
#cgo LDFLAGS: -luuid
#cgo CFLAGS: -I /usr/include/
#include <uuid/uuid.h>
#include <stdlib.h>
#include <errno.h>
int p_sizeOfUuid = sizeof(uuid_t);
int p_bufLen = -1;
int p_nextTokenIndex = 0;
unsigned char * p_UuidBuf;

void allocUuidBuf(const int bufLen) {
	p_bufLen = bufLen;
	p_UuidBuf = malloc(p_sizeOfUuid * bufLen);
}

void freeUuidBuf() {
	free(p_UuidBuf);
}

void populateBuf() {
	for(int i=0; i < p_bufLen; i++) {
		uuid_generate_random(p_UuidBuf + i * p_sizeOfUuid);
	}
}

unsigned char * getNextToken() {
	if(p_nextTokenIndex == p_bufLen) {
		errno = EIO;
		return 0;
	}
	return (p_UuidBuf + p_nextTokenIndex++ * p_sizeOfUuid);
}
*/
import "C"
import "unsafe"
import "fmt"
//import "bytes"

type uuid struct { }

func (u uuid) New() ([]byte) {
	var _, err = C.allocUuidBuf(1)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	defer C.freeUuidBuf()
	_, err = C.populateBuf()
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	b := C.GoBytes(unsafe.Pointer(C.getNextToken()), C.p_sizeOfUuid)	
	return b
}

func main() {
	s := uuid{}
	fmt.Println("hello uuid:", s.New())	
	return
}
