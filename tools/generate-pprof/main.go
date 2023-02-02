package main

import (
	"log"
	"os"
	"time"

	"github.com/google/pprof/profile"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	const mainBinary = "/usr/bin/pprof-demo"

	var cpuM = []*profile.Mapping{
		{
			ID:              1,
			Start:           0x10000,
			Limit:           0x40000,
			File:            mainBinary,
			HasFunctions:    true,
			HasFilenames:    true,
			HasLineNumbers:  true,
			HasInlineFrames: true,
		},
	}

	var cpuF = []*profile.Function{
		{ID: 1, Name: "main", SystemName: "main", Filename: "main.go"},
		{ID: 2, Name: "doALot", SystemName: "doALot", Filename: "main.go"},
		{ID: 3, Name: "doLittle", SystemName: "doLittle", Filename: "main.go"},
		{ID: 4, Name: "prepare", SystemName: "prepare", Filename: "main.go"},
	}

	var cpuL = []*profile.Location{
		{
			ID:      1,
			Mapping: cpuM[0],
			Address: 0x1,
			Line: []profile.Line{
				{Function: cpuF[0], Line: 4},
			},
		},
		{
			ID:      2,
			Mapping: cpuM[0],
			Address: 0x2,
			Line: []profile.Line{
				{Function: cpuF[1], Line: 15},
			},
		},
		{
			ID:      3,
			Mapping: cpuM[0],
			Address: 0x3,
			Line: []profile.Line{
				{Function: cpuF[2], Line: 20},
			},
		},
		{
			ID:      4,
			Mapping: cpuM[0],
			Address: 0x4,
			Line: []profile.Line{
				{Function: cpuF[3], Line: 10},
			},
		},
	}

	var testProfile = &profile.Profile{
		PeriodType:    &profile.ValueType{Type: "cpu", Unit: "nanoseconds"},
		Period:        1000000,
		TimeNanos:     time.Now().UnixNano(),
		DurationNanos: 1e9,
		SampleType: []*profile.ValueType{
			{Type: "samples", Unit: "count"},
		},
		Sample: []*profile.Sample{
			{
				Location: []*profile.Location{cpuL[0]},
				Value:    []int64{3},
			},
			{
				Location: []*profile.Location{cpuL[1], cpuL[0]},
				Value:    []int64{20},
			},
			{
				Location: []*profile.Location{cpuL[3], cpuL[1], cpuL[0]},
				Value:    []int64{5},
			},
			{
				Location: []*profile.Location{cpuL[2], cpuL[0]},
				Value:    []int64{5},
			},
			{
				Location: []*profile.Location{cpuL[3], cpuL[2], cpuL[0]},
				Value:    []int64{5},
			},
		},
		Location: cpuL,
		Function: cpuF,
		Mapping:  cpuM,
	}

	f, err := os.Create("cpu.pb.gz")
	if err != nil {
		return err
	}
	defer f.Close()

	return testProfile.Write(f)
}
