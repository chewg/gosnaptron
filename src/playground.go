package main

import (
	"snaptron_api/src/query"
	"snaptron_api/src/ops"
	"snaptron_api/src/server"
)


func main() {
	//jir_tss_hg19()

	ssc_example()

	//tsv_example()
}


func jir_tss_hg19() {

	DATASOURCE := "srav1"

	//	region	contains	filters	group
	//	chr2:29446395-30142858	1	strand=-	A_NormalTSS
	//	chr2:29416789-29446394	1	strand=-	B_AltTSS

	// BASIC LEVEL
	A_NormalTSS := query.Region()
	B_AltTSS := query.Region()

	A_NormalTSS.Chromosome("chr2").Start_Pos(29446395).End_Pos(30142858).Contains(true)
	B_AltTSS.Chromosome("chr2").Start_Pos(29416789).End_Pos(29446394).Contains(true)

	filter := query.Filter()
	filter.Strand_Minus(true)

	params_A := query.Build_Params(A_NormalTSS, filter)
	params_B := query.Build_Params(B_AltTSS, filter)

	server_data_A := server.Get_mRNA_From_Server(params_A, DATASOURCE)
	server_data_B := server.Get_mRNA_From_Server(params_B, DATASOURCE)

	dataframe_A := query.DataFrame().From_Server_Data(server_data_A)
	dataframe_B := query.DataFrame().From_Server_Data(server_data_B)


	// Intermediate Level
	frames_A := ops.Dataframe_To_Frames(dataframe_A)
	frames_B := ops.Dataframe_To_Frames(dataframe_B)

	frames_A = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_A)
	frames_B = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_B)

	frames_A = ops.Summarize(frames_A, ops.Sum_Count)
	frames_B = ops.Summarize(frames_B, ops.Sum_Count)

	frames_with_jir := ops.Calculate_Ratio(ops.JIR_Ratio, frames_A, frames_B)
	frames_with_jir = ops.Order(frames_with_jir, ops.Decr_Stat)
	frames_with_jir = ops.Load_Metadata_Into_Frames(frames_with_jir, DATASOURCE, 77)

	ops.Print_jir(frames_with_jir)
}


func ssc_example() {
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

	region_1 := query.Region()
	region_2 := query.Region()

	region_1.Chromosome("chr6").Start_Pos(32831148).End_Pos(32831148).Either_End(true)
	region_2.Chromosome("chr6").Start_Pos(32831182).End_Pos(32831182).Either_Start(true)

	filter := query.Filter()
	filter.Strand_Minus(true).Samples_Count(">", 0)

	params_1 := query.Build_Params(region_1, filter)
	params_2 := query.Build_Params(region_2, filter)

	return get_ssc_frames(params_1, params_2)
}


func ssc_p1k3cd_nonvalidated() *[]ops.Frame {
	// chr1:9664595-9664595	2	strand=+&samples_count>=1	PIK3CD non_validated
	// chr1:9664759-9664759	1	strand=+&samples_count>=1	PIK3CD non_validated

	region_1 := query.Region()
	region_2 := query.Region()

	region_1.Chromosome("chr1").Start_Pos(9664595).End_Pos(9664595).Either_End(true)
	region_2.Chromosome("chr1").Start_Pos(9664759).End_Pos(9664759).Either_Start(true)

	filter := query.Filter()
	filter.Strand_Plus(true).Samples_Count(">", 0)

	params_1 := query.Build_Params(region_1, filter)
	params_2 := query.Build_Params(region_2, filter)

	return get_ssc_frames(params_1, params_2)
}


func ssc_gnb1_validated() *[]ops.Frame {
	//chr1:1879786-1879786	2	strand=-&samples_count>=1
	//chr1:1879903-1879903	1	strand=-&samples_count>=1

	region_1 := query.Region()
	region_2 := query.Region()

	region_1.Chromosome("chr1").Start_Pos(1879786).End_Pos(1879786).Either_End(true)
	region_2.Chromosome("chr1").Start_Pos(1879903).End_Pos(1879903).Either_Start(true)

	filter := query.Filter()
	filter.Strand_Minus(true).Samples_Count(">", 0)

	params_1 := query.Build_Params(region_1, filter)
	params_2 := query.Build_Params(region_2, filter)

	return get_ssc_frames(params_1, params_2)
}


func get_ssc_frames(params_1, params_2 string) *[]ops.Frame {
	DATASOURCE := "gtex"

	server_data_1 := server.Get_mRNA_From_Server(params_1, DATASOURCE)
	server_data_2 := server.Get_mRNA_From_Server(params_2, DATASOURCE)

	dataframe_1 := query.DataFrame().From_Server_Data(server_data_1)
	dataframe_2 := query.DataFrame().From_Server_Data(server_data_2)

	// Start intermediate
	frames_from_q1 := ops.Dataframe_To_Frames(dataframe_1)
	frames_from_q2 := ops.Dataframe_To_Frames(dataframe_2)

	frames_from_q1 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q1)
	frames_from_q2 = ops.Group_By(ops.Group_Frames_By_Sample_ID, frames_from_q2)

	group := ops.Union(frames_from_q1, frames_from_q2)
	group = ops.Summarize(group, ops.Sum_Count)

	return group
}


func tsv_example() {
	gnb1_group := ssc_gnb1_validated()
	tap2_group := ssc_tap2_validated()

	intersect_groups := ops.Intersect(gnb1_group, tap2_group)
	union_groups := ops.Union(gnb1_group, tap2_group)

	intersect_groups = ops.Order(intersect_groups, ops.Incr_Sample_ID)
	union_groups = ops.Order(union_groups, ops.Incr_Sample_ID)

	DATASOURCE := "gtex"
	union_groups = ops.Load_Metadata_Into_Frames(union_groups, DATASOURCE, 60)

	ops.Print_tsv(intersect_groups, union_groups)
}

