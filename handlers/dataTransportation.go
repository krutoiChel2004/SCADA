package handlers

import (
	"log"
	"test/conn"
	"test/database"
	"test/models"
	"test/services"
	"time"

	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/contrib/websocket"
)

// func createJSON(config models.Config, modbusConn *conn.ModbusTCP) []map[string]float32 {
// 	var data []map[string]float32

// 	vReg1, err := modbusConn.ReadRegisters(1, 3, 16484, 108)
// 	vReg2, err := modbusConn.ReadRegisters(1, 3, 16592, 108)

// 	if err != nil {
// 		log.Printf("Ошибка чтения регистров: %v", err)
// 	}

// 	combined := append(vReg1, vReg2...)

// 	for _, Reg := range config.Reg {

// 		v := modbusGo.ByteToFloat32(combined[Reg.Data["AddresReg"]*4 : (Reg.Data["AddresReg"]*4)+4])

// 		m := make(map[string]float32)
// 		m[Reg.Name] = v

// 		data = append(data, m)
// 	}

// 	return data
// }

// Обработчик WebSocket
func C2C(c *websocket.Conn) {
	name := c.Params("name")

	plan, err := ioutil.ReadFile("configs/config-" + name + ".json")
	if err != nil {
		log.Printf("Ошибка открытия файла: %v", err)
		return
	}

	var config models.Config

	if err := json.Unmarshal(plan, &config); err != nil {
		log.Printf("Ошибка десериализации JSON: %v", err)
		return
	}

	modbusConn, err := conn.GetModbusConnection(config.IP)
	if err != nil {
		log.Printf("Ошибка подключения к Modbus TCP: %v", err)
		return
	}
	defer modbusConn.Close()

	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	defer db.Close()

	for {

		query, queryParams, err := services.CreateInsertQuery(config, modbusConn)

		if err != nil {
			log.Fatalf("Ошибка формирования запроса: %v", err)
			break
		}

		_, err = db.NamedExec(query, queryParams)
		if err != nil {
			log.Fatalf("Ошибка вставки данных: %v", err)
			break
		}

		time.Sleep(5 * time.Second)
	}

	log.Println("Завершаем работу и закрываем WebSocket соединение")
	c.Close()
}
