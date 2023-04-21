#another implementation of the kmer index problem, using sets instead of dictionaries

input_string = 'ATCGATCGACTGACT'

k_input = int(input("Enter the k-value: "))

def kmer_index(input, k):
  #set for the unique k-mers seen so far
  kmers = set()

  #stores the starting index of the k-mer
  kmer_index = set() 

  for x in range(len(input)-1+k):
    kmer = input[x:x+k]

    #if the kmer being read has already been read, then add the index to the kmer_index set. if it hasn't, then add it to the kmers set
    if kmer in kmers:
        kmer_index.add(x)
    else:
        kmers.add(kmer)

  return kmer_index

print(kmer_index(input_string, k_input))

#need to figure out how to output both the kmer and its starting index. only outputting indexes right now. also, do they need to be sorted?




