/*
	Пакет storage реализует примитивное хранилище в фаловой системе.
 	Так как файловые системы плохо работают с большим количеством файлов
	в одной директории, в долгосрочной перспективе это означает, что проекту
	нужна другая реализация хранилища перед релизом.
	Также нужно оценить насколько высока вероятность конфликтов при нейминге
	файлов по временной метке.
*/
package storage

import (
	"log"
	"os"
	"strconv"

	"github.com/Meexe/blockchain/internal/models"
)

// storage - хранилище блокчейна
type storage struct {
	dir string
}

// New инициализирует хранилище для блокчейна в файловой системе
func New(dir string) *storage {
	return &storage{dir}
}

// SaveBlockchain дампит блокчейн в файловую систему
func (s *storage) SaveBlockchain(bc models.Blockchain) (err error) {
	for _, block := range bc.Chain {
		if err = s.SaveBlock(block); err != nil {
			return
		}
	}
	return
}

// LoadBlockchain подгружает блокчейн из файловой системы
func (s *storage) LoadBlockchain() (chain models.Blockchain, err error) {
	files, err := os.ReadDir(s.dir)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			log.Print("file is dir - skipping")
			continue
		}

		var (
			body  []byte
			block models.Block
		)
		if _, err = strconv.ParseInt(file.Name(), 10, 64); err != nil {
			log.Print("invalid block file name")
			continue
		}

		if body, err = os.ReadFile(s.dir + file.Name()); err != nil {
			return
		}

		if err = block.UnmarshalJSON(body); err != nil {
			return
		}

		chain.Chain = append(chain.Chain, block)
	}
	return
}

// SaveBlock дампит блок в файловую систему
// если файл с меткой уже существует - перезаписывает его
func (s *storage) SaveBlock(b models.Block) (err error) {
	var data []byte
	if data, err = b.MarshalJSON(); err != nil {
		return
	}
	return os.WriteFile(s.dir+strconv.FormatInt(b.Ts, 10), data, 0666)
}

// LoadBlock подгружает блок из файловой системы
func (s *storage) LoadBlock(ts int64) (b models.Block, err error) {
	var data []byte
	if data, err = os.ReadFile(s.dir + strconv.FormatInt(ts, 10)); err != nil {
		return
	}

	err = b.UnmarshalJSON(data)
	return
}
