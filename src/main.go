package main

import (
	"snaptron_api/src/query"
	"snaptron_api/src/web"
)

func main() {
	simulate_flexibility(25)
	// simulate_region_and_filter()
	// simulate_restful()
}

func simulate_flexibility(i int) {
	start := 29446395 + i
	end := 30142858
	gt := ">"


	region := query.Region()
	region.Gene("abc")
	region.Reset_Gene()

	region.Chromosome("CD99")
	region.Reset_Chromosome()
	region.Chromosome("chr2").Start_Pos(start).End_Pos(end)

	filter := query.Filter()
	filter.Coverage_Sum(gt, 200).Samples_Count(gt, 150)

	q_str_1 := query.Execute(filter, region)

	print("Q_STR_1\n")
	print(q_str_1)


	// `regions=KMT2E&rfilter=samples_count>:5&sfilter=description:cortex"`
	new_region := query.Region()
	new_region.Gene("KMT2E")

	new_filter := query.Filter()
	new_filter.Samples_Count(">", 5)

	new_metadata := query.Metadata()
	new_metadata.Key("description").Value("cortex")

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


