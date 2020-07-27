package fileoperation

import (
	"fmt"
	"encoding/csv"
	"os"
)

func FileRead(){
	csvFile,err := os.Open("data.csv");
	if err != nil {
		fmt.Println("File open error :",err)
	}
	fileData := csv.NewReader(csvFile);
	data , e :=fileData.Read();
	if e!=nil {
		fmt.Println("data read error :",e)
	}
	fmt.Println(data[0]);
}