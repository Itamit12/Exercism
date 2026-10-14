package phonenumber
import "errors"



func Number(phoneNumber string) (string, error) {
    var result string
    number:=0
	if string(phoneNumber[0])=="1"{
        for i:=1;i<len(phoneNumber);i++{
            if int(phoneNumber[i])>=48&&int(phoneNumber[i])<=57{
                result+=string(phoneNumber[i])
                number++
            }
        }
    }else if phoneNumber[0:2]=="+1"{
        for i:=2;i<len(phoneNumber);i++{
            if int(phoneNumber[i])>=48&&int(phoneNumber[i])<=57{
                result+=string(phoneNumber[i])
                number++
            }
        }
    }else{
        for i:=0;i<len(phoneNumber);i++{
            if int(phoneNumber[i])>=48&&int(phoneNumber[i])<=57{
                result+=string(phoneNumber[i])
                number++
            }
        }
    }
    
    if number<10||number>11{
        return "",errors.New("error")
    }
    if number ==11&&(string(phoneNumber[0])!="1"||(string(phoneNumber[0])!="+"&&string(phoneNumber[1])!="1")){
        return "",errors.New("error")
    }
    if string(result[0])=="0"||string(result[0])=="1"{
        return "",errors.New("error")
    }else if string(result[3])=="0"||string(result[3])=="1"{
        return "",errors.New("error")
    }
	return result,nil
}

func AreaCode(phoneNumber string) (string, error) {
	result,err:=Number(phoneNumber)
    if err != nil{
        return "",errors.New("error")
    }
    var res string
    for i:=0;i<3;i++{
        res+=string(result[i])
    }
    return res,nil
}

func Format(phoneNumber string) (string, error) {
	result,err:=Number(phoneNumber)
    if err != nil{
        return "",errors.New("error")
    }
    return "("+result[0:3]+") "+result[3:6]+"-"+result[6:],nil
}
