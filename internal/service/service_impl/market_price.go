package service_impl

import (
	"context"
	"errors"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service"
)

type MarketPriceService struct{}

const BatchSize = 10000
const BlanksRatioError = 0.6

func (m MarketPriceService) RetrieveInterval(ctx context.Context, marketItemId int64, start int64, end int64, repo repository.MarketPriceRepository) *domain.MarketPricesRetrieve {
	return repo.RetrieveInterval(ctx, marketItemId, start, end)
}

func (m MarketPriceService) RetrieveLast(ctx context.Context, marketItemIds []int64, repo repository.MarketPriceRepository) *domain.MarketPricesRetrieve {
	lastMarketPrices := make([]domain.MarketPrice, 0)
	for _, miId := range marketItemIds {
		lastMarketPriceRetrieved := repo.RetrieveLast(ctx, miId)
		if lastMarketPriceRetrieved.Error == nil && lastMarketPriceRetrieved.MarketPrice != nil {
			lastMarketPrices = append(lastMarketPrices, *lastMarketPriceRetrieved.MarketPrice)
		}
	}
	marketPricesRetrieve := domain.MarketPricesRetrieve{
		MarketPrices: lastMarketPrices,
		Error:        nil,
	}
	return &marketPricesRetrieve
}

func (m MarketPriceService) Create(ctx context.Context, marketPrice *domain.MarketPrice, repo repository.MarketPriceRepository) (bool, error) {
	return repo.Create(ctx, marketPrice)
}

func (m MarketPriceService) BulkCreate(ctx context.Context, marketPrices *[]domain.MarketPrice, repo repository.MarketPriceRepository) (int64, error) {
	batchSize := BatchSize
	wholeSize := len(*marketPrices)
	createdRows := int64(0)

	for i := 0; wholeSize > 0; i++ {
		batch := (*marketPrices)[i*batchSize : i*batchSize+service.MinInt(batchSize, wholeSize)]
		rows, err := repo.BulkCreate(ctx, &batch)
		if err != nil {
			return createdRows, err
		}
		wholeSize -= batchSize
		createdRows += rows
	}

	return createdRows, nil
}

func (m MarketPriceService) FillBlanks(ctx context.Context, marketPrices []domain.MarketPrice, intervals []int64) ([]domain.MarketPrice, int) {
	resultMarketPrices := make([]domain.MarketPrice, len(intervals))
	indexesUsed := make([]int, 0)

	if len(intervals) > 0 && len(marketPrices) > 0 {
		currentIndex := 0
		endTS := &intervals[currentIndex]

		currentMinIndex := 0
		currentMinDiff := service.AbsInt64(marketPrices[currentMinIndex].Timestamp - *endTS)
		currentMPIndex := currentMinIndex

		for currentIndex < len(intervals) && currentMPIndex < len(marketPrices) {
			currentDiff := service.AbsInt64(marketPrices[currentMPIndex].Timestamp - *endTS)
			if currentMinDiff < currentDiff {
				resultMarketPrices[currentIndex] = domain.MarketPrice{
					MarketItemId: marketPrices[currentMinIndex].MarketItemId,
					Price:        marketPrices[currentMinIndex].Price,
					Timestamp:    *endTS,
				}

				currentIndex++
				if currentIndex < len(intervals) {
					endTS = &intervals[currentIndex]
					currentMinDiff = service.AbsInt64(marketPrices[currentMinIndex].Timestamp - *endTS)
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

func (m MarketPriceService) RetrieveMarketItemPrices(
	ctx context.Context,
	marketItemId int64,
	startTimeStamp int64,
	endTimeStamp int64,
	interval int64,
	repo repository.MarketPriceRepository,
) *domain.MarketPricesRetrieve {
	intervals := service.GetIntervals(startTimeStamp, endTimeStamp, interval)
	marketPricesRetrieved := m.RetrieveInterval(ctx, marketItemId, startTimeStamp, endTimeStamp, repo)
	if marketPricesRetrieved.Error != nil {
		return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: marketPricesRetrieved.Error}
	}
	marketPrices, blanksCount := m.FillBlanks(ctx, marketPricesRetrieved.MarketPrices, intervals)
	blanksRatio := float64(blanksCount) / float64(len(intervals))

	if blanksRatio > BlanksRatioError {
		return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: errors.New("empty or not enough data")}
	}

	return &domain.MarketPricesRetrieve{MarketPrices: marketPrices, Error: nil}
}
