#in the algorithms meeting on 03/08, we discussed about different methods to find the index of where a specific kmer begins on a string. ian showed us how to use regex to do it and told asked me to use a different approach as homework

input_string = 'ATCGATCGACTGACT'

k_input = int(input("Enter the k-value: "))

def kmer_index(input, k):
    #create a dictionary where the key will be the k-mer, and the value will be the list of indexes where it appears
    kmer_dict = {}

    #go through the string (not sure if the range bound is correct)
    for x in range(len(input)-k+1):
        kmer_idx = input[x: x+k]

        #if the kmer is in the dictionary, record the index, if not, create it
        if kmer_idx in kmer_dict:
            kmer_dict[kmer_idx].append(x)
        else:
            kmer_dict[kmer_idx] = [x]
    
    #output the dictionary 
    return kmer_dict

print(kmer_index(input_string, k_input))

#Questions to ask:
#-----------------
# 1. is the output correct?
# 2. should the zero index be in the list?