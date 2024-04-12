package main

import (
	"encoding/json"
	"flag"
	"os"
	"runtime"

	"github.com/Nazhgam/sum.git/calculate"
	"github.com/Nazhgam/sum.git/domain"
	"github.com/labstack/gommon/log"
)

func main() {
	workers := flag.Int("workers", runtime.GOMAXPROCS(-1), "default")
	flag.Parse()

	log.Infof("workers count is: %d", *workers)

	numbers, err := ReadJson("payload_json/payload.json")
	if err != nil {
		return
	}

	calculate.Calculate(*workers, numbers)

}

func ReadJson(filePath string) ([]domain.Number, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Errorf("open file error: %w", err)
		return nil, err
	}

	defer f.Close()

	var numbers []domain.Number

	if err := json.NewDecoder(f).Decode(&numbers); err != nil {
		log.Errorf("decode error: %w", err)
		return nil, err
	}

	log.Infof("success read json file. length of slice: %d", len(numbers))
	return numbers, nil
}
