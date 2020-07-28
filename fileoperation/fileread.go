package fileoperation

import (
	"fmt"
	"encoding/csv"
	"os"
	"io"
)

func FileRead(datachan chan string,exitstatus *string){
	csvFile,err := os.Open("data.csv");
	if err != nil {
		fmt.Println("File open error :",err)
	}
	fileData := csv.NewReader(csvFile);
	for{
		data , e :=fileData.Read();
		if e == io.EOF{
			break;
		}
		datachan<-data[0];
		if e!=nil {
			fmt.Println("data read error :",e)
		}
	}
	datachan<-"endoffile"
	*exitstatus="true";
}