package common

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func ReadPkg(conn net.Conn) (msg Message, err error) {

	var buf = make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	//从conn读pkgLen个字节存到buf
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}
	fmt.Println("获取数据成功", msg)
	return
}

func WritePkg(conn net.Conn, data []byte) (err error) {

	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte

	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	n, err := conn.Write(buf[:4])

	if n != 4 || err != nil {
		fmt.Println("发送长度失败", err)
		return
	}
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送数据失败", err)
		return
	}
	fmt.Println("发送数据成功", string(data))
	return

}
