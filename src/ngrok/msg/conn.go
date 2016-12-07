package msg

import (
	"encoding/binary"
	"errors"
	"fmt"
	"ngrok/conn"
)

func readMsgShared(c conn.Conn) (buffer []byte, err error) {
	c.Debug("等待读取消息")

	var sz int64
	err = binary.Read(c, binary.LittleEndian, &sz)
	if err != nil {
		return
	}
	c.Debug("读取消息长度: %d", sz)

	buffer = make([]byte, sz)
	n, err := c.Read(buffer)
	c.Debug("读取消息 %s", buffer)

	if err != nil {
		return
	}

	if int64(n) != sz {
		err = errors.New(fmt.Sprintf("预期要读取 %d bytes, 但只读取 %d", sz, n))
		return
	}

	return
}

func ReadMsg(c conn.Conn) (msg Message, err error) {
	buffer, err := readMsgShared(c)
	if err != nil {
		return
	}

	return Unpack(buffer)
}

func ReadMsgInto(c conn.Conn, msg Message) (err error) {
	buffer, err := readMsgShared(c)
	if err != nil {
		return
	}
	return UnpackInto(buffer, msg)
}

func WriteMsg(c conn.Conn, msg interface{}) (err error) {
	buffer, err := Pack(msg)
	if err != nil {
		return
	}

	c.Debug("写消息: %s", string(buffer))
	err = binary.Write(c, binary.LittleEndian, int64(len(buffer)))

	if err != nil {
		return
	}

	if _, err = c.Write(buffer); err != nil {
		return
	}

	return nil
}
