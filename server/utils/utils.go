package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gotest/chat/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) ReadPkg() (msg common.Message, err error) {

	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4])
	//从conn读pkgLen个字节存到buf
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	json.Unmarshal(t.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}
	fmt.Println("获取数据成功", msg)
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {

	pkgLen := uint32(len(data))

	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen)

	n, err := t.Conn.Write(t.Buf[:4])

	if n != 4 || err != nil {
		fmt.Println("发送长度失败", err)
		return
	}
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送数据失败", err)
		return
	}
	fmt.Println("发送数据成功", string(data))
	return

}
