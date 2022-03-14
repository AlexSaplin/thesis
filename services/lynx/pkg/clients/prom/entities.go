package prom

import "github.com/prometheus/common/model"

type MetricsResponse struct {
	RunCount  string `json:"run_count"`
	RunTime   string `json:"run_time"`
	LoadCount string `json:"load_count"`
	LoadTime  string `json:"load_time"`
}

type FullMetricsResponse struct {
	Metrics []MetricsResponse `json:"metrics"`
}

func tryGet(matrix model.Matrix, i int) string {
	if matrix.Len() > 0 {
		return matrix[0].Values[i].String()
	}
	return ""
}

func NewFullMetricsResponse(rTime model.Value, rCnt model.Value, lTime model.Value, lCnt model.Value) (result FullMetricsResponse) {
	if lTime.(model.Matrix).Len() == 0 {
		return
	}
	for i := range lTime.(model.Matrix)[0].Values {
		result.Metrics = append(result.Metrics, MetricsResponse{
			RunTime:   rTime.(model.Matrix)[0].Values[i].Value.String(),
			RunCount:  rCnt.(model.Matrix)[0].Values[i].Value.String(),
			LoadTime:  lTime.(model.Matrix)[0].Values[i].Value.String(),
			LoadCount: lCnt.(model.Matrix)[0].Values[i].Value.String(),
		})
	}
	return
}
