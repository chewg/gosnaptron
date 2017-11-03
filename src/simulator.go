package main

import (
	"snaptron_api/src/query"
	"snaptron_api/src/web"
	"strings"
	"snaptron_api/src/data"
	"snaptron_api/src/examples"
	"snaptron_api/src/ops"
	"fmt"
)

func main() {
	// test_frame_jir()
	test_frame_ssc()

	// ssc()
	// simulate_flexibility(25)
	// simulate_region_and_filter()
	// simulate_restful()
}

func test_frame_jir() {
	region1 := query.Region()
	region1.Chromosome("chr4").Start_Pos(20763023).End_Pos(20763023)
	region1.Either_End(true)

	filter1 := query.Filter()
	filter1.Coverage_Sum(">", 1)

	query_string_1 := query.Execute(region1, filter1)
	dataframe_1 := query_string_to_dataframe(query_string_1)


	region2 := query.Region()
	region2.Chromosome("chr4").Start_Pos(20763098).End_Pos(20763098)
	region2.Either_Start(true)

	filter2 := query.Filter()
	filter2.Coverage_Sum(">", 1)

	query_string_2 := query.Execute(region2, filter2)
	dataframe_2 := query_string_to_dataframe(query_string_2)


	// Start intermediate
	frames_from_q1 := ops.Import_Dataframe(dataframe_1)
	frames_from_q2 := ops.Import_Dataframe(dataframe_2)

	group := ops.Group(frames_from_q1, frames_from_q2)
	group = ops.Summarize(group, ops.Sum_Count_By_Sample_ID)


}



func test_frame_ssc() {
	//chr6:1-514015&rfilter=samples_count>:100
	region1 := query.Region()
	region1.Chromosome("chr6").Start_Pos(1).End_Pos(514015)

	filter1 := query.Filter()
	filter1.Samples_Count(">", 100)

	q_str_1 := query.Execute(region1, filter1)

	df := query_string_to_dataframe(q_str_1)

	frames := ops.Import_Dataframe(df)

	frames = ops.Summarize(frames, ops.Sum_Count_By_Sample_ID)

	frames = ops.Arrange(frames, ops.Decr_Count, ops.Decr_Sample_ID)

	fmt.Print(*frames)
	fmt.Print("Done")
}


func query_string_to_dataframe(str string) data.Dataframe {
	df := data.DataFrame()

	lines := strings.Split(str, "\n")
	// print(lines)

	for _, line := range lines[1 : len(lines) - 1] {
		data_cells := strings.Split(line, "\t")
		df.Load_DataFrames(data_cells)
	}

	return df
}


func ssc() {
	region1 := query.Region()
	region1.Chromosome("chr1").Start_Pos(1879786).End_Pos(1879786)
	region1.Either_End(true)

	filter1 := query.Filter()
	filter1.Strand_Minus(true).Samples_Count(">", 0)

	region2 := query.Region()
	region2.Chromosome("chr1").Start_Pos(1879903).End_Pos(1879903)
	region2.Either_Start(true)

	filter2 := query.Filter()
	filter2.Strand_Minus(true).Samples_Count(">=", 1)		//server todo server, tell Chris: fails if >=

	q_str_1 := query.Execute(region1, filter1)

	df := data.DataFrame()
	lines := strings.Split(q_str_1, "\n")
	print(lines)

	for _, line := range lines[1 : len(lines) - 1] {
		data_cells := strings.Split(line, "\t")
		df.Load_DataFrames(data_cells)
	}

	ssc := examples.Shared_Sample_Count(df)

	for k, v := range ssc {
		print("\n")
		print(k)
		print(" : ")
		print(v)
	}

	print(ssc)

}


func simulate_flexibility(i int) {
	start := 29446395 + i
	end := 30142858
	gt := ">"


	region_1 := query.Region()
	region_1.Gene("abc")

	region_2 := query.Region()
	region_2.Chromosome("CD99")
	region_2.Chromosome("chr2").Start_Pos(start).End_Pos(end)

	filter := query.Filter()
	filter.Coverage_Sum(gt, 200).Samples_Count(gt, 150)

	q_str_1 := query.Execute(filter, region_2)

	print("Q_STR_1\n")
	print(q_str_1)


	// `regions=KMT2E&rfilter=samples_count>:5&sfilter=description:cortex"`
	new_region := query.Region()
	new_region.Gene("KMT2E")

	new_filter := query.Filter()
	new_filter.Samples_Count(gt, 5)

	new_metadata := query.Metadata()
	new_metadata.Key("description").Value("","cortex")

	query_result_2 := query.Execute(new_metadata, new_region, new_filter)
	print("\nQ_STR_2\n")
	print(query_result_2)
}


func simulate_region_and_filter() {
	region := query.Region()
	region.Chromosome("chr2").Start_Pos(29446395).End_Pos(30142858)

	filter := query.Filter()
	filter.Samples_Count(">", 100).Coverage_Sum(">", 150)

	q_str := query.Execute(filter, region)
	print(q_str)
}


func simulate_restful() {
	answer := web.Get("http://snaptron.cs.jhu.edu/srav2/snaptron?regions=chr6:1-514015&rfilter=samples_count:100")
	print(answer)
}


