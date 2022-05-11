package service_impl

import (
	"errors"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
	"math"
)

type MarketPriceService struct{}

//func (m MarketPriceService) RetrieveInterval(marketItemId int64, start int64, end int64, repo repository.MarketPriceRepository) <-chan domain.MarketPricesRetrieve {
//	channel := make(chan domain.MarketPricesRetrieve)
//
//	go func() {
//		channel <- *repo.RetrieveInterval(marketItemId, start, end)
//		close(channel)
//	}()
//
//	return channel
//}
//
//func (m MarketPriceService) Create(marketPrice *domain.MarketPrice, repo repository.MarketPriceRepository) <-chan error {
//	channel := make(chan error)
//
//	go func() {
//		channel <- repo.Create(marketPrice)
//		close(channel)
//	}()
//
//	return channel
//}

func absInt64(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}

func (m MarketPriceService) RetrieveInterval(marketItemId int64, start int64, end int64, repo repository.MarketPriceRepository) domain.MarketPricesRetrieve {
	return *repo.RetrieveInterval(marketItemId, start, end)
}

func (m MarketPriceService) Create(marketPrice *domain.MarketPrice, repo repository.MarketPriceRepository) error {
	return repo.Create(marketPrice)
}

func (m MarketPriceService) FillBlanks(marketPrices []domain.MarketPrice, intervals []int64) ([]domain.MarketPrice, int) {
	resultMarketPrices := make([]domain.MarketPrice, len(intervals))
	indexesUsed := make([]int, 0)

	if len(intervals) > 0 && len(marketPrices) > 0 {
		currentIndex := 0
		endTS := &intervals[currentIndex]

		currentMinIndex := 0
		currentMinDiff := absInt64(marketPrices[currentMinIndex].Timestamp - *endTS)
		currentMPIndex := currentMinIndex

		for currentIndex < len(intervals) && currentMPIndex < len(marketPrices) {
			currentDiff := absInt64(marketPrices[currentMPIndex].Timestamp - *endTS)
			if currentMinDiff < currentDiff {
				resultMarketPrices[currentIndex] = domain.MarketPrice{
					MarketItemId: marketPrices[currentMinIndex].MarketItemId,
					Price:        marketPrices[currentMinIndex].Price,
					Timestamp:    *endTS,
				}

				//fmt.Println(
				//	time.Unix(resultMarketPrices[currentIndex].Timestamp, 0),
				//	time.Unix(marketPrices[currentMinIndex].Timestamp, 0),
				//	resultMarketPrices[currentIndex].Price,
				//)

				currentIndex++
				if currentIndex < len(intervals) {
					endTS = &intervals[currentIndex]
					currentMinDiff = absInt64(marketPrices[currentMinIndex].Timestamp - *endTS)
				}
				if len(indexesUsed) == 0 || currentMinIndex != indexesUsed[len(indexesUsed)-1] {
					indexesUsed = append(indexesUsed, currentMinIndex)
				}
			} else {
				currentMinDiff = currentDiff
				currentMinIndex = currentMPIndex
				currentMPIndex++
			}
		}
		if len(resultMarketPrices) > 0 {
			for currentIndex < len(intervals) {
				resultMarketPrices[currentIndex] = domain.MarketPrice{
					MarketItemId: resultMarketPrices[currentIndex-1].MarketItemId,
					Price:        resultMarketPrices[currentIndex-1].Price,
					Timestamp:    intervals[currentIndex],
				}
				currentIndex++
			}
		}
	}

	blanksCount := len(intervals) - len(indexesUsed)
	return resultMarketPrices, blanksCount
}

func (m MarketPriceService) GetIntervals(startTimeStamp int64, endTimeStamp int64, interval int64) []int64 {
	intervalsCount := int64(0)
	difference := endTimeStamp - startTimeStamp
	if difference > 0 {
		intervalsCount = int64(math.Ceil(float64(difference) / float64(interval)))
	}

	intervals := make([]int64, intervalsCount)
	for i := int64(0); i < intervalsCount; i++ {
		intervals[i] = startTimeStamp + i*interval
	}

	return intervals
}

func (m MarketPriceService) GetMarketItemPrices(
	marketItemId int64,
	startTimeStamp int64,
	endTimeStamp int64,
	interval int64,
	repo repository.MarketPriceRepository,
) *domain.MarketPricesRetrieve {

	intervals := m.GetIntervals(startTimeStamp, endTimeStamp, interval)
	marketPricesRetrieved := m.RetrieveInterval(marketItemId, startTimeStamp, endTimeStamp, repo)
	if marketPricesRetrieved.Error != nil {
		return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: marketPricesRetrieved.Error}
	}
	marketPrices, blanksCount := m.FillBlanks(marketPricesRetrieved.MarketPrices, intervals)
	blanksRatio := float64(blanksCount) / float64(len(intervals))
	if blanksRatio > 0.6 {
		return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: errors.New("empty or not enough data")}
	}

	return &domain.MarketPricesRetrieve{MarketPrices: marketPrices, Error: nil}
}
