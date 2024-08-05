package alarm

import (
	"fmt"
	"sort"

	"github.com/EDDYCJY/go-gin-example/models"
)

type AlarmService struct{}

func (a *AlarmService) GetNewAlarms() ([]*models.Alarms, error) {

	return models.GetNewAlarms()

}

type AlarmTypeItem struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type ByValue []AlarmTypeItem

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value > a[j].Value }

func (a *AlarmService) GetAlarmTypeStat() ([]AlarmTypeItem, error) {
	alarms, err := models.GetResolvedAlarms()
	if err != nil {
		return nil, err
	}

	alarmMap := map[string]int{}
	for _, item := range alarms {
		alarmMap[item.Name]++
	}

	output := make([]AlarmTypeItem, 0)
	for k, v := range alarmMap {
		output = append(output, AlarmTypeItem{k, v})
	}
	sort.Sort(ByValue(output))

	otherNum := 0
	for i := 5; i < len(output); i++ {
		otherNum += output[i].Value
	}
	if len(output) > 5 {
		output = output[:5]
		output = append(output, AlarmTypeItem{
			Key:   "其他",
			Value: otherNum,
		})
	}

	return output, nil
}

type CompanyStat struct {
	Company     string `json:"company"`
	Num         int    `json:"num"`
	TotalSecond int    `json:"totalSecond"`
	Avg         int    `json:"avg"`
	AvgStr      string `json:"avgStr"`
}

type ByCompanyAvg []*CompanyStat

func (a ByCompanyAvg) Len() int           { return len(a) }
func (a ByCompanyAvg) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCompanyAvg) Less(i, j int) bool { return a[i].Avg > a[j].Avg }

func (a *AlarmService) GetCompanyStat() ([]*CompanyStat, error) {
	alarms, err := models.GetResolvedAlarms()
	if err != nil {
		return nil, err
	}

	companyMap := make(map[string]*CompanyStat)
	for _, item := range alarms {
		if _, ok := companyMap[item.Company]; !ok {
			companyMap[item.Company] = &CompanyStat{
				Company: item.Company,
			}
		}
		companyMap[item.Company].Num++
		companyMap[item.Company].TotalSecond += int(item.AlarmEndAt.Sub(item.AlarmStartAt).Seconds())
	}
	for k, v := range companyMap {
		v.Company = k
		v.Avg = v.TotalSecond / v.Num
		v.AvgStr = secondsToDHM(v.Avg)
	}

	output := make([]*CompanyStat, 0)
	for _, item := range companyMap {
		output = append(output, item)
	}

	sort.Sort(ByCompanyAvg(output))
	return output, nil
}

type CaptainStat struct {
	Captain     string `json:"captain"`
	Num         int    `json:"num"`
	TotalSecond int    `json:"totalSecond"`
	Avg         int    `json:"avg"`
	AvgStr      string `json:"avgStr"`
}

type ByCaptainAvg []*CaptainStat

func (a ByCaptainAvg) Len() int           { return len(a) }
func (a ByCaptainAvg) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCaptainAvg) Less(i, j int) bool { return a[i].Avg > a[j].Avg }

func (a *AlarmService) GetCaptainStat() ([]*CaptainStat, error) {
	alarms, err := models.GetResolvedAlarms()
	if err != nil {
		return nil, err
	}

	captainMap := make(map[string]*CaptainStat)
	for _, item := range alarms {
		if _, ok := captainMap[item.Captain]; !ok {
			captainMap[item.Captain] = &CaptainStat{
				Captain: item.Captain,
			}
		}
		captainMap[item.Captain].Num++
		captainMap[item.Captain].TotalSecond += int(item.AlarmEndAt.Sub(item.AlarmStartAt).Seconds())
	}
	for k, v := range captainMap {
		v.Captain = k
		v.Avg = v.TotalSecond / v.Num
		v.AvgStr = secondsToDHM(v.Avg)
	}

	output := make([]*CaptainStat, 0)
	for _, item := range captainMap {
		output = append(output, item)
	}

	sort.Sort(ByCaptainAvg(output))
	return output[:5], nil
}

func secondsToDHM(seconds int) (output string) {
	// 1秒钟
	const second = 1
	// 1分钟
	const minute = 60 * second
	// 1小时
	const hour = 60 * minute
	// 1天
	const day = 24 * hour

	days := seconds / day
	seconds -= days * day
	hours := seconds / hour
	seconds -= hours * hour
	minutes := seconds / minute
	seconds -= minutes * minute

	if day != 0 {
		return fmt.Sprintf("%d 天 %d 小时 %d 分钟 %d 秒", days, hours, minutes, seconds)
	}
	if hour != 0 {
		return fmt.Sprintf("%d 小时 %d 分钟 %d 秒", hours, minutes, seconds)
	}
	if minute != 0 {
		return fmt.Sprintf("%d 分钟 %d 秒", minutes, seconds)
	}
	return fmt.Sprintf("%d 秒", seconds)
}
