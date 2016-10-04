package main
/*
#cgo LDFLAGS: -luuid
#cgo CFLAGS: -I /usr/include/
#include <uuid/uuid.h>
#include <stdlib.h>
int p_sizeOfUuid = sizeof(uuid_t);
int p_bufLen = -1;
int p_nextTokenIndex = 0;
unsigned char * p_UuidBuf;
unsigned char * allocUuidBuf(const int bufLen) {
	p_bufLen = bufLen;
	p_UuidBuf = malloc(p_sizeOfUuid * bufLen);
	return p_UuidBuf;
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
	return (p_UuidBuf + p_nextTokenIndex++ * p_sizeOfUuid);
}
*/
import "C"
import "unsafe"
import "fmt"

func main() {
	C.allocUuidBuf(10)
	C.populateBuf()	
	fmt.Println("hello uuid:", C.GoBytes(unsafe.Pointer(C.getNextToken()), C.p_sizeOfUuid))
	fmt.Println("hello uuid:", C.GoBytes(unsafe.Pointer(C.getNextToken()), C.p_sizeOfUuid))
	fmt.Println("hello uuid:", C.GoBytes(unsafe.Pointer(C.getNextToken()), C.p_sizeOfUuid))
	fmt.Println("hello uuid:", C.GoBytes(unsafe.Pointer(C.getNextToken()), C.p_sizeOfUuid))
	defer C.freeUuidBuf()
	return
}
