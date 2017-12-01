#SRAv2's sample ID range is: 0-50097
#some will always have 0 coverage

library(RCurl)
library(Matrix)
library(GenomicRanges)
library(dplyr)


getSnaptronDataFrame <- function(url, verbose=FALSE)
{
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

  #this block is from Leo as well
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

  #start of cwilks' code; Leo might not want to own this part :)
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

url_1 <- 'http://snaptron.cs.jhu.edu/srav2/snaptron?regions=chr4:20763023-20763023&either=2&rfilter=coverage_sum>:1'
url_2 <- 'http://snaptron.cs.jhu.edu/srav2/snaptron?regions=chr4:20763098-20763098&either=1&rfilter=coverage_sum>:1'

sdf_1 <- getSnaptronDataFrame(url_1,verbose=TRUE)
sdf_2 <- getSnaptronDataFrame(url_2,verbose=TRUE)

sc_1 <- sdf_1['sample_counts']
sc_2 <- sdf_2['sample_counts']


df_1 <- sc_1[[1]]
df_2 <- sc_2[[1]]


# set operations
intersect_df <- dplyr::inner_join(df_1, df_2, by="sample")
union_df <- dplyr::full_join(df_1, df_2, by="sample")



# shared sample count
df_1 %>% 
group_by(junction, sample) %>%
summarise(frequency = sum(coverage)) %>%
filter(frequency > 1) %>%
arrange(desc(frequency))



# tissue specificity

df %>%
group_by(sample) %>%
mutate(presence = f(v, d))

# if in interesect, return 1. else return 0
f <- function(v, d) {
  # v[d == "b"] - v[d == "a"]
}


# how do across 2 dataframes?


# junction inclusion ratio

df %>%
group_by(sample) %>%
mutate(jir = jir_f(v, d))

jir_f <- function(c1, c2) ((c1 - c2) / (c1 + c2 + 1))

