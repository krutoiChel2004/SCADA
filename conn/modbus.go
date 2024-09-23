package conn

import (
	"log"
	"time"

	"github.com/krutoiChel2004/modbusGo"
)

// Структура для Modbus TCP соединения
type ModbusTCP struct {
	Name, Address string
	handler       *modbusGo.TCPHandler
}

// Функция для получения Modbus TCP соединения
func GetModbusConnection(address string) (*ModbusTCP, error) {
	modbusConn := &ModbusTCP{Address: address}

	// Подключаемся к устройству по Modbus TCP
	handler := modbusGo.NewTCPHandler(address, 20*time.Second)
	if err := handler.Connect(); err != nil {
		return nil, err
	}

	modbusConn.handler = handler
	log.Println("Подключение к Modbus TCP по адресу", address, "успешно установлено")
	return modbusConn, nil
}

// Метод для закрытия соединения
func (m *ModbusTCP) Close() {
	if m.handler != nil {
		m.handler.Close()
		log.Println("Соединение с", m.Address, "закрыто")
	}
}

// Пример использования функции для чтения регистров
func (m *ModbusTCP) ReadRegisters(slaveID, functionCode, address, quantity uint16) ([]byte, error) {
	results, err := m.handler.ReadRegisters(slaveID, functionCode, address, quantity)
	if err != nil {
		return nil, err
	}
	return results, nil
}
