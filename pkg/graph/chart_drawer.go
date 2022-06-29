package graph

import (
	"bytes"
	"fmt"
	"github.com/abrbird/portfolio_bot/pkg/api"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"math"
	"time"
)

type ChartDrawer struct{}

func New() *ChartDrawer {
	return &ChartDrawer{}
}

func (chPl *ChartDrawer) MarketItem(
	marketItem *api.MarketItem,
	marketItemPrices []*api.MarketPrice,
	portfolioMarketItem *api.PortfolioItem,
) (*bytes.Buffer, error) {

	xValues := make([]time.Time, len(marketItemPrices))
	yValues := make([]float64, len(marketItemPrices))
	yPValues := make([]float64, len(marketItemPrices))
	yMin := math.MaxFloat64
	yMax := -math.MaxFloat64
	for i, mip := range marketItemPrices {
		xValues[i] = time.Unix(mip.GetTimestamp(), 0)
		yValues[i] = mip.GetPrice()
		yPValues[i] = portfolioMarketItem.GetPrice()

		if mip.GetPrice() > yMax {
			yMax = mip.GetPrice()
		}
		if mip.GetPrice() < yMin {
			yMin = mip.GetPrice()
		}
	}

	priceSeries := chart.TimeSeries{
		Name: marketItem.GetCode(),
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xValues,
		YValues: yValues,
	}

	smaSeries := chart.SMASeries{
		Name: "SMA",
		Style: chart.Style{
			StrokeColor:     drawing.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: priceSeries,
	}

	bbSeries := &chart.BollingerBandsSeries{
		Name: "Bollinger Bands",
		Style: chart.Style{
			StrokeColor: drawing.ColorFromHex("efefef"),
			FillColor:   drawing.ColorFromHex("efefef").WithAlpha(64),
		},
		InnerSeries: priceSeries,
	}
	_ = bbSeries
	_ = smaSeries

	series := []chart.Series{
		priceSeries,
		//bbSeries,
		//smaSeries,
	}

	if portfolioMarketItem != nil {
		pSeries := &chart.MinSeries{
			Name: "Purchase price",
			Style: chart.Style{
				StrokeColor:     chart.ColorAlternateGray,
				StrokeDashArray: []float64{5.0, 5.0},
				//Show:            true,
			},
			InnerSeries: chart.TimeSeries{
				XValues: xValues,
				YValues: yPValues,
			},
		}
		series = append(series, pSeries)
	}

	graph := chart.Chart{
		Title: fmt.Sprintf("%s", marketItem.GetCode()),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Canvas: chart.Style{
			FillColor: drawing.ColorFromHex("dedede"),
		},
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeHourValueFormatter,
		},
		YAxis: chart.YAxis{
			Name: "Price",
			Range: &chart.ContinuousRange{
				Max: yMax + math.Abs(yMax)*0.05,
				Min: yMin - math.Abs(yMin)*0.05,
			},
		},
		Series: series,
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (chPl *ChartDrawer) PortfolioSummary(
	baseShift float64,
	portfolioItems []*api.PortfolioItem,
	itemsPricesMap map[int64][]*api.MarketPrice,
) (*bytes.Buffer, error) {

	portfolioItemsMap := make(map[int64]*api.PortfolioItem, 0)
	for _, pi := range portfolioItems {
		portfolioItemsMap[pi.GetMarketItemId()] = pi
	}

	profitStyle := chart.Style{
		FillColor:   drawing.ColorFromHex("13c158"),
		StrokeColor: drawing.ColorFromHex("13c158"),
		StrokeWidth: 0,
	}

	lossStyle := chart.Style{
		FillColor:   drawing.ColorFromHex("c11313"),
		StrokeColor: drawing.ColorFromHex("c11313"),
		StrokeWidth: 0,
	}

	bars := make([]chart.Value, len(itemsPricesMap[portfolioItems[0].GetMarketItemId()]))
	for marketItemId, portfolioItemPrices := range itemsPricesMap {
		for i, pip := range portfolioItemPrices {
			bars[i].Value += pip.GetPrice() * portfolioItemsMap[marketItemId].GetVolume()
			bars[i].Label = time.Unix(pip.GetTimestamp(), 0).Format("Jan 02")
		}
	}
	yMin := math.MaxFloat64
	yMax := -math.MaxFloat64
	for i, _ := range bars {
		bars[i].Value -= baseShift
		bars[i].Value = bars[i].Value / baseShift * 100

		if bars[i].Value < 0 {
			bars[i].Style = lossStyle
		} else {
			bars[i].Style = profitStyle
		}

		if bars[i].Value > yMax {
			yMax = bars[i].Value
		}
		if bars[i].Value < yMin {
			yMin = bars[i].Value
		}
	}
	if yMin > 0 {
		yMin = 0
	}
	if yMax < 0 {
		yMax = 0
	}

	sbc := chart.BarChart{
		Title: "Portfolio PNL, %",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Canvas: chart.Style{
			FillColor: drawing.ColorFromHex("dedede"),
		},
		Height:   512,
		BarWidth: 50,
		YAxis: chart.YAxis{
			Range: &chart.ContinuousRange{
				Max: yMax + math.Abs(yMax)*0.01,
				Min: yMin - math.Abs(yMin)*0.01,
			},
		},
		UseBaseValue: true,
		BaseValue:    0.0,
		Bars:         bars,
	}

	buffer := bytes.NewBuffer([]byte{})
	err := sbc.Render(chart.PNG, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
