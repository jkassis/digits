// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package requester

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sort"
	"time"
)

const (
	barChar = "■"
)

// We report for max 1M results.
const maxRes = 1000000

type report struct {
	avgTotal float64
	fastest  float64
	slowest  float64
	average  float64
	rps      float64

	offsets []float64

	results chan *result
	done    chan bool
	total   time.Duration

	errorDist map[string]int
	lats      []float64
	sizeTotal int64
	numRes    int64
	output    string

	w io.Writer
}

func newReport(w io.Writer, results chan *result, output string, n int) *report {
	cap := min(n, maxRes)
	return &report{
		output:    output,
		results:   results,
		done:      make(chan bool, 1),
		errorDist: make(map[string]int),
		w:         w,
		lats:      make([]float64, 0, cap),
		offsets:   make([]float64, 0, cap),
	}
}

func runReporter(r *report) {
	// Loop will continue until channel is closed
	for res := range r.results {
		r.numRes++
		if res.err != nil {
			r.errorDist[res.err.Error()]++
		} else {
			r.avgTotal += res.duration.Seconds()
			if len(r.lats) < maxRes {
				r.lats = append(r.lats, res.duration.Seconds())
				r.offsets = append(r.offsets, res.offset.Seconds())
			}
		}
	}
	// Signal reporter is done.
	r.done <- true
}

func (r *report) finalize(total time.Duration) {
	r.total = total
	r.rps = float64(r.numRes) / r.total.Seconds()
	r.average = r.avgTotal / float64(len(r.lats))
	r.print()
}

func (r *report) print() {
	buf := &bytes.Buffer{}
	if err := newTemplate(r.output).Execute(buf, r.snapshot()); err != nil {
		log.Println("error:", err.Error())
		return
	}
	r.printf(buf.String())

	r.printf("\n")
}

func (r *report) printf(s string, v ...interface{}) {
	fmt.Fprintf(r.w, s, v...)
}

func (r *report) snapshot() Report {
	snapshot := Report{
		AvgTotal:  r.avgTotal,
		Average:   r.average,
		Rps:       r.rps,
		SizeTotal: r.sizeTotal,
		Total:     r.total,
		ErrorDist: r.errorDist,
		NumRes:    r.numRes,
		Lats:      make([]float64, len(r.lats)),
		Offsets:   make([]float64, len(r.lats)),
	}

	if len(r.lats) == 0 {
		return snapshot
	}

	snapshot.SizeReq = r.sizeTotal / int64(len(r.lats))

	copy(snapshot.Lats, r.lats)
	copy(snapshot.Offsets, r.offsets)

	sort.Float64s(r.lats)
	r.fastest = r.lats[0]
	r.slowest = r.lats[len(r.lats)-1]

	snapshot.Histogram = r.histogram()
	snapshot.LatencyDistribution = r.latencies()

	snapshot.Fastest = r.fastest
	snapshot.Slowest = r.slowest

	return snapshot
}

func (r *report) latencies() []LatencyDistribution {
	pctls := []int{10, 25, 50, 75, 90, 95, 99}
	data := make([]float64, len(pctls))
	j := 0
	for i := 0; i < len(r.lats) && j < len(pctls); i++ {
		current := i * 100 / len(r.lats)
		if current >= pctls[j] {
			data[j] = r.lats[i]
			j++
		}
	}
	res := make([]LatencyDistribution, len(pctls))
	for i := 0; i < len(pctls); i++ {
		if data[i] > 0 {
			res[i] = LatencyDistribution{Percentage: pctls[i], Latency: data[i]}
		}
	}
	return res
}

func (r *report) histogram() []Bucket {
	bc := 10
	buckets := make([]float64, bc+1)
	counts := make([]int, bc+1)
	bs := (r.slowest - r.fastest) / float64(bc)
	for i := 0; i < bc; i++ {
		buckets[i] = r.fastest + bs*float64(i)
	}
	buckets[bc] = r.slowest
	var bi int
	var max int
	for i := 0; i < len(r.lats); {
		if r.lats[i] <= buckets[bi] {
			i++
			counts[bi]++
			if max < counts[bi] {
				max = counts[bi]
			}
		} else if bi < len(buckets)-1 {
			bi++
		}
	}
	res := make([]Bucket, len(buckets))
	for i := 0; i < len(buckets); i++ {
		res[i] = Bucket{
			Mark:      buckets[i],
			Count:     counts[i],
			Frequency: float64(counts[i]) / float64(len(r.lats)),
		}
	}
	return res
}

// Report is the output
type Report struct {
	AvgTotal float64
	Fastest  float64
	Slowest  float64
	Average  float64
	Rps      float64

	AvgConn  float64
	AvgDNS   float64
	AvgReq   float64
	AvgRes   float64
	AvgDelay float64
	ConnMax  float64
	ConnMin  float64
	ReqMax   float64
	ReqMin   float64

	Lats      []float64
	ReqLats   []float64
	DelayLats []float64
	Offsets   []float64

	Total time.Duration

	ErrorDist map[string]int
	SizeTotal int64
	SizeReq   int64
	NumRes    int64

	LatencyDistribution []LatencyDistribution
	Histogram           []Bucket
}

// LatencyDistribution does what you think
type LatencyDistribution struct {
	Percentage int
	Latency    float64
}

// Bucket does what you think
type Bucket struct {
	Mark      float64
	Count     int
	Frequency float64
}
