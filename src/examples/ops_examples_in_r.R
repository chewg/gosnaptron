#SRAv2's sample ID range is: 0-50097
#some will always have 0 coverage

library(RCurl)
library(Matrix)
library(GenomicRanges)
library(dplyr)


getSnaptronDataFrame <- function(url, verbose=FALSE)
{
  # based off of Leo's code
  to_chr_list <- function(x) {
    r <- strsplit(x, ',')
    i <- which(sapply(r, function(y) { y[[1]] == '0' }))
    if(length(i) > 0) r[i] <- NA
    return(CharacterList(r))
  }

  start_time <- Sys.time()

  csv_data <- read.csv(textConnection(getURI(url)), sep='\t')
  
  end_time <- Sys.time()
  
  elapsed_time <- end_time - start_time
  if(verbose) message(paste(elapsed_time,"seconds to query Snaptron and download",dim(csv_data)[1],"junction rows"))

  
  colnames(csv_data) <- c('type', 'snaptron_id', 'chromosome', 'start', 'end',
                     'length', 'strand', 'annotated', 'left_motif', 'right_motif',
                     'left_annotated', 'right_annotated', 'samples',
                     'samples_count', 'coverage_sum', 'coverage_avg', 'coverage_median',
                     'source_dataset_id')
  snaptron_data <- GRanges(seqnames = csv_data[, 'chromosome'], 
                    IRanges(as.numeric(csv_data[, 'start']), as.numeric(csv_data[, 'end'])),
                    strand = csv_data[, 'strand'])
  snaptron_data$type <- as.factor(csv_data[, 'type'])
  snaptron_data$snaptron_id <- as.integer(csv_data[, 'snaptron_id'])
  snaptron_data$annotated <- to_chr_list(as.character(csv_data[, 'annotated']))
  snaptron_data$left_motif <- csv_data[, 'left_motif']
  snaptron_data$right_motif <- csv_data[, 'right_motif']
  snaptron_data$left_annotated <- to_chr_list(as.character(csv_data[, 'left_annotated']))
  snaptron_data$right_annotated <- to_chr_list(as.character(csv_data[, 'right_annotated']))
  snaptron_data$samples_count <- as.integer(csv_data[, 'samples_count'])
  snaptron_data$coverage_sum <- as.integer(csv_data[, 'coverage_sum'])
  snaptron_data$coverage_avg <- as.numeric(csv_data[, 'coverage_avg'])
  snaptron_data$coverage_median <- as.numeric(csv_data[, 'coverage_median'])
  snaptron_data$source_dataset_id <- as.integer(csv_data[, 'source_dataset_id'])

  # based off of cwilks' code
  nr<-nrow(csv_data)
  nc<-ncol(csv_data)
  rows<-NULL
  cols<-NULL
  vals<-NULL
  
  start_time <- Sys.time()
  
  i <- 1
  
  while(i <= nr)
  {
    c2 <- strsplit(as.character(csv_data$samples[i]),",")
    aa <- unlist(lapply(c2[[1]],function(x) { strsplit(x,':')[[1]][1] }))
    ab <- unlist(lapply(c2[[1]],function(x) { strsplit(x,':')[[1]][2] }))
    
    rows <- c(rows,rep(as.integer(csv_data[i, 'snaptron_id']),length(aa[2:length(aa)])))
    cols <- c(cols,aa[2:length(aa)])
    vals <- c(vals,ab[2:length(ab)])
    i <- i + 1
  }

  end_time <- Sys.time()
  elapsed_time <- end_time - start_time
  if(verbose) message(paste(elapsed_time,"seconds to setup sample count matrix"))
  
  #start at 0 since the sample ID's do
  #convert to sparseMatrix for expansion & storage benefits
  sample_counts <- data.frame(rows,as.numeric(cols),as.numeric(vals),stringsAsFactors=FALSE)
  colnames(sample_counts) <- c("junction","sample","coverage")
  
  return(list("ranges"=snaptron_data,"sample_counts"=sample_counts))
}


# charlie's ops

url_1 <- 'http://snaptron.cs.jhu.edu/srav2/snaptron?regions=chr4:20763023-20763023&either=2&rfilter=coverage_sum>:1'
url_2 <- 'http://snaptron.cs.jhu.edu/srav2/snaptron?regions=chr4:20763098-20763098&either=1&rfilter=coverage_sum>:1'

sdf_1 <- getSnaptronDataFrame(url_1,verbose=TRUE)
sdf_2 <- getSnaptronDataFrame(url_2,verbose=TRUE)

sc_1 <- sdf_1['sample_counts']
sc_2 <- sdf_2['sample_counts']


df_1 <- sc_1[[1]]
df_2 <- sc_2[[1]]


#
# shared sample count
#

ssc_df <- df_1 %>% 
group_by(junction, sample) %>%
summarise(frequency = sum(coverage)) %>%
filter(frequency > 1) %>%
arrange(desc(frequency))

#
# end
#



#
# tissue specificity v2
#

keep_cols <- c("sample", "coverage")
nojunction_df_1 <- df_1[keep_cols]
nojunction_df_2 <- df_2[keep_cols]

intersect_df <- dplyr::inner_join(nojunction_df_1, nojunction_df_2, by="sample")
intersect_df$coverage <- rowSums(intersect_df[, c(2, 3)])
keep_cols <- c("sample", "coverage")
intersect_df <- intersect_df[keep_cols]


union_df <- dplyr::full_join(nojunction_df_1, nojunction_df_2, by="sample")
union_df$coverage.x[is.na(union_df$coverage.x)] <- 0
union_df$coverage.y[is.na(union_df$coverage.y)] <- 0
union_df$coverage <- rowSums(union_df[, c(2, 3)])
union_df <- union_df[keep_cols]

ts_df <- dplyr::left_join(union_df, intersect_df, by="sample")

# create new TS present column
ts_df$present <- rowSums(ts_df[, c(2, 3)])
ts_df$present <- ifelse(ts_df$present > 0, 1, 0)
ts_df$present[is.na(ts_df$present)] <- 0

#
# end
#



#
# junction inclusion ratio
#

keep_cols <- c("sample", "coverage")
nojunction_df_1 <- df_1[keep_cols]
nojunction_df_2 <- df_2[keep_cols]

union_df <- dplyr::full_join(nojunction_df_1, nojunction_df_2, by="sample")

union_df$coverage.x[is.na(union_df$coverage.x)] <- 0
union_df$coverage.y[is.na(union_df$coverage.y)] <- 0

jir_df <- union_df %>% group_by(sample) %>% mutate(jir = (coverage.x - coverage.y)/(coverage.x + coverage.y + 1))

#
# end
#
