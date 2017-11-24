package main

import (
	"snaptron_api/src/query"
	"snaptron_api/src/data"
	"snaptron_api/src/ops"
	"fmt"
	"snaptron_api/src/examples"
	"snaptron_api/src/web"
)


func main() {
	// simulate_jir_calc()
	// simulate_tsv_calc()
	simulate_ssc_calc()

	// old_simulates()
}


func old_simulates() {
	 test_frame_ssc()

	 ssc()
	 simulate_flexibility(25)
	 simulate_region_and_filter()
	 simulate_restful()
}


func simulate_jir_calc() {
	region1 := query.Region()
	region2 := query.Region()

	region1.Chromosome("chr4").Start_Pos(20763023).End_Pos(20763023)
	region1.Either_End(true)
	region2.Chromosome("chr4").Start_Pos(20763098).End_Pos(20763098)
	region2.Either_Start(true)

	filter := query.Filter()
	filter.Coverage_Sum(">", 1)

	query_string_1 := query.Execute(region1, filter)
	query_string_2 := query.Execute(region2, filter)

	dataframe_1 := data.DataFrame().From_Query_String(query_string_1)
	dataframe_2 := data.DataFrame().From_Query_String(query_string_2)


	// Start intermediate
	frames_from_q1 := ops.Dataframe_To_Frames(dataframe_1)
	frames_from_q2 := ops.Dataframe_To_Frames(dataframe_2)

	frames_from_q1 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q1)
	frames_from_q2 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q2)

	frames_from_q1 = ops.Summarize(frames_from_q1, ops.Sum_Count)
	frames_from_q2 = ops.Summarize(frames_from_q2, ops.Sum_Count)

	frames_with_jir := ops.Calculate_Ratio(ops.JIR_Ratio, frames_from_q1, frames_from_q2)
	frames_with_jir = ops.Order(frames_with_jir, ops.Decr_Stat)

	url := fmt.Sprintf("http://snaptron.cs.jhu.edu/%v/samples?all=1", "srav2")
	frames_with_jir = ops.Load_Metadata_Into_Frames(frames_with_jir, url)

	ops.Print_jir(frames_with_jir)
}


func simulate_tsv_calc() {
	gnb1_group := ssc_gnb1_validated()
	tap2_group := ssc_tap2_validated()

	intersect_groups := ops.Intersect(gnb1_group, tap2_group)
	union_groups := ops.Union(gnb1_group, tap2_group)

	intersect_groups = ops.Order(intersect_groups, ops.Incr_Sample_ID)
	union_groups = ops.Order(union_groups, ops.Incr_Sample_ID)

	url := fmt.Sprintf("http://snaptron.cs.jhu.edu/%v/samples?all=1", "srav2")
	union_groups = ops.Load_Metadata_Into_Frames(union_groups, url)

	ops.Print_tsv(intersect_groups, union_groups)
}


func simulate_ssc_calc() {
	gnb1_group := ssc_gnb1_validated()
	p1k3cd_group := ssc_p1k3cd_nonvalidated()
	tap2_group := ssc_tap2_validated()

	all_groups := ops.Intersect(gnb1_group, p1k3cd_group, tap2_group)
	all_groups = ops.Summarize(all_groups, ops.Sum_Count)
	all_groups = ops.Order(all_groups, ops.Incr_Count, ops.Incr_Sample_ID)

	ops.Print_ssc(all_groups)
}


func ssc_tap2_validated() *[]ops.Frame {
	// chr6:32831148-32831148	2	strand=-&samples_count>=1	TAP2 validated
	// chr6:32831182-32831182	1	strand=-&samples_count>=1

	region1 := query.Region()
	region1.Chromosome("chr6").Start_Pos(32831148).End_Pos(32831148)
	region1.Either_End(true)

	filter1 := query.Filter()
	filter1.Strand_Minus(true).Samples_Count(">", 0)

	query_string_1 := query.Execute(region1, filter1)
	dataframe_1 := data.DataFrame().From_Query_String(query_string_1)


	region2 := query.Region()
	region2.Chromosome("chr6").Start_Pos(32831182).End_Pos(32831182)
	region2.Either_Start(true)

	filter2 := query.Filter()
	filter2.Strand_Minus(true).Samples_Count(">", 0)

	query_string_2 := query.Execute(region2, filter2)
	dataframe_2 := data.DataFrame().From_Query_String(query_string_2)


	// Start intermediate
	frames_from_q1 := ops.Dataframe_To_Frames(dataframe_1)
	frames_from_q2 := ops.Dataframe_To_Frames(dataframe_2)

	frames_from_q1 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q1)
	frames_from_q2 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q2)

	group := ops.Union(frames_from_q1, frames_from_q2)
	group = ops.Summarize(group, ops.Sum_Count)
	return group
}


func ssc_p1k3cd_nonvalidated() *[]ops.Frame {
	// chr1:9664595-9664595	2	strand=+&samples_count>=1	PIK3CD non_validated
	// chr1:9664759-9664759	1	strand=+&samples_count>=1	PIK3CD non_validated

	region1 := query.Region()
	region1.Chromosome("chr1").Start_Pos(9664595).End_Pos(9664595)
	region1.Either_End(true)

	filter1 := query.Filter()
	filter1.Strand_Plus(true).Samples_Count(">", 0)

	query_string_1 := query.Execute(region1, filter1)
	dataframe_1 := data.DataFrame().From_Query_String(query_string_1)


	region2 := query.Region()
	region2.Chromosome("chr1").Start_Pos(9664759).End_Pos(9664759)
	region2.Either_Start(true)

	filter2 := query.Filter()
	filter2.Strand_Plus(true).Samples_Count(">", 0)

	query_string_2 := query.Execute(region2, filter2)
	dataframe_2 := data.DataFrame().From_Query_String(query_string_2)


	// Start intermediate
	frames_from_q1 := ops.Dataframe_To_Frames(dataframe_1)
	frames_from_q2 := ops.Dataframe_To_Frames(dataframe_2)

	frames_from_q1 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q1)
	frames_from_q2 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q2)

	group := ops.Union(frames_from_q1, frames_from_q2)
	group = ops.Summarize(group, ops.Sum_Count)
	return group
}


func ssc_gnb1_validated() *[]ops.Frame {
	//chr1:1879786-1879786	2	strand=-&samples_count>=1
	//chr1:1879903-1879903	1	strand=-&samples_count>=1

	region1 := query.Region()
	region1.Chromosome("chr1").Start_Pos(1879786).End_Pos(1879786)
	region1.Either_End(true)

	filter1 := query.Filter()
	filter1.Strand_Minus(true).Samples_Count(">", 0)

	query_string_1 := query.Execute(region1, filter1)
	dataframe_1 := data.DataFrame().From_Query_String(query_string_1)


	region2 := query.Region()
	region2.Chromosome("chr1").Start_Pos(1879903).End_Pos(1879903)
	region2.Either_Start(true)

	filter2 := query.Filter()
	filter2.Strand_Minus(true).Samples_Count(">", 0)

	query_string_2 := query.Execute(region2, filter2)
	dataframe_2 := data.DataFrame().From_Query_String(query_string_2)


	// Start intermediate
	frames_from_q1 := ops.Dataframe_To_Frames(dataframe_1)
	frames_from_q2 := ops.Dataframe_To_Frames(dataframe_2)

	frames_from_q1 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q1)
	frames_from_q2 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q2)

	group := ops.Union(frames_from_q1, frames_from_q2)
	group = ops.Summarize(group, ops.Sum_Count)
	return group
}






func test_frame_ssc() {
	//chr6:1-514015&rfilter=samples_count>:100
	region1 := query.Region()
	region1.Chromosome("chr6").Start_Pos(1).End_Pos(514015)

	filter1 := query.Filter()
	filter1.Samples_Count(">", 100)

	q_str_1 := query.Execute(region1, filter1)
	df := data.DataFrame().From_Query_String(q_str_1)

	frames := ops.Dataframe_To_Frames(df)
	frames = ops.Summarize(frames, ops.Sum_Count)
	frames = ops.Order(frames, ops.Decr_Count, ops.Decr_Sample_ID)

	fmt.Print(*frames)
	fmt.Print("Done")
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

	df := data.DataFrame().From_Query_String(q_str_1)

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
