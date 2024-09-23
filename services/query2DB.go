package services

import (
	"fmt"
	"strings"
	"test/conn"
	"test/models"
	"time"

	"github.com/krutoiChel2004/modbusGo"
)

func CreateInsertQuery(config models.Config, modbusConn *conn.ModbusTCP) (string, map[string]interface{}, error) {
	data, err := ReadReg(16484, 16700, 120, modbusConn)
	if err != nil {
		return "", nil, fmt.Errorf("error reading Modbus registers: %v", err)
	}

	// Формирование частей SQL-запроса
	var columns []string
	var placeholders []string
	queryParams := make(map[string]interface{})

	for i, reg := range config.Reg {
		// Извлечение данных
		v := modbusGo.ByteToFloat32(data[reg.Data["AddresReg"]*4 : (reg.Data["AddresReg"]*4)+4])

		// Формирование имен столбцов и placeholders для SQL
		columnName := fmt.Sprintf("col_%d", i+1) // Для уникальности имён в запросе
		columns = append(columns, reg.Name)      // Добавляем оригинальное имя
		placeholders = append(placeholders, fmt.Sprintf(":%s", columnName))

		// Добавляем данные в map для передачи в NamedExec
		queryParams[columnName] = v
	}

	queryParams["date_time"] = time.Now()

	query := fmt.Sprintf(
		"INSERT INTO container_data (%s, date_time) VALUES (%s, :date_time)",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, queryParams, nil
}
