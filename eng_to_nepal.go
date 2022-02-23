package convert

import (
	"errors"
	"strconv"
)

func GetNepalMonth(m int) string {
	n_month := ""
	switch m {
	case 1:
		n_month = "वैशाख"
		break

	case 2:
		n_month = "जेठ"
		break

	case 3:
		n_month = "असार"
		break

	case 4:
		n_month = "साउन"
		break

	case 5:
		n_month = "भदौ"
		break

	case 6:
		n_month = "असोज"
		break

	case 7:
		n_month = "कात्तिक"
		break

	case 8:
		n_month = "मंसीर"
		break

	case 9:
		n_month = "पुस"
		break

	case 10:
		n_month = "माघ"
		break

	case 11:
		n_month = "फागुन"
		break

	case 12:
		n_month = "चैत्र"
		break
	}
	return n_month
}

func isRangeEng(yy, mm, dd int) bool {
	if yy < 1944 || yy > 2033 {
		return false
	}

	if mm < 1 || mm > 12 {
		return false
	}

	if dd < 1 || dd > 31 {
		return false
	}

	return true
}

func EngToNep(yy, mm, dd int) (error, map[string]string) {
	if isRangeEng(yy, mm, dd) == false {
		return errors.New("Range not supported"), nil
	} else {
		month := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		lmonth := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

		def_eyy := 1944 //spear head english date...
		def_nyy := 2000
		def_nmm := 9
		def_ndd := 17 - 1 //spear head nepali date...
		total_eDays := 0
		total_nDays := 0
		a := 0
		day := 7 - 1 //all the initializations...
		m := 0
		y := 0
		i := 0
		j := 0
		//	numDay := 0

		// count total no. of days in-terms of year
		for i := 0; i < (yy - def_eyy); i++ { //total days for month calculation...(english)
			if isLeap(def_eyy+i) == true {
				for j := 0; j < 12; j++ {
					total_eDays += lmonth[j]
				}
			} else {
				for j := 0; j < 12; j++ {
					total_eDays += month[j]
				}
			}
		}

		// count total no. of days in-terms of month
		for i := 0; i < (mm - 1); i++ {
			if isLeap(yy) == true {
				total_eDays += lmonth[i]
			} else {
				total_eDays += month[i]
			}
		}

		// count total no. of days in-terms of date
		total_eDays += dd

		i = 0
		j = def_nmm
		total_nDays = def_ndd
		m = def_nmm
		y = def_nyy

		// count nepali date from array
		for total_eDays != 0 {
			a = Date[i][j]
			total_nDays++ //count the days
			day++         //count the days interms of 7 days
			if total_nDays > a {
				m++
				total_nDays = 1
				j++
			}
			if day > 7 {
				day = 1
			}
			if m > 12 {
				y++
				m = 1
			}
			if j > 12 {
				j = 1
				i++
			}
			total_eDays--
		}

		Nep_date["year"] = strconv.Itoa(y)
		Nep_date["month"] = strconv.Itoa(m)
		Nep_date["date"] = strconv.Itoa(total_nDays)
		Nep_date["day"] = GetDayOfWeek(day)
		Nep_date["nmonth"] = GetNepalMonth(m)

		return nil, Nep_date
	}
}
