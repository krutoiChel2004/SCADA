package services

import (
	"log"
	"math"
	"test/conn"
)

func ReadReg(startReg, endReg, maxRegsPerRequest int, modbusConn *conn.ModbusTCP) ([]byte, error) {

	// Общий массив для хранения всех прочитанных регистров
	var allData []byte

	// Общее количество регистров, которые нужно прочитать
	totalRegs := endReg - startReg + 1

	// Количество запросов
	numRequests := int(math.Ceil(float64(totalRegs) / float64(maxRegsPerRequest)))

	for i := 0; i < numRequests; i++ {
		// Вычисляем начальный регистр для текущего запроса
		currStartReg := startReg + i*maxRegsPerRequest

		// Вычисляем количество регистров для текущего запроса
		remainingRegs := endReg - currStartReg + 1
		numRegs := int(math.Min(float64(remainingRegs), float64(maxRegsPerRequest)))

		// Выполняем запрос
		data, err := modbusConn.ReadRegisters(1, 3, uint16(currStartReg), uint16(numRegs))
		if err != nil {
			log.Printf("Ошибка чтения регистров 1: %v", err)
		}

		// Добавляем данные в общий массив
		allData = append(allData, data...)
	}

	return allData, nil
}
