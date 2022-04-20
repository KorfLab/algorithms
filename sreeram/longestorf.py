
import argparse
from readfasta import read_record

arg_parser = argparse.ArgumentParser(
    description='Returns the protein sequence with the longest open reading frame')

arg_parser.add_argument('fasta_file', type=str, metavar='file',
                        help='enter the path to the fasta file you want to use')

arg_parser.add_argument('-frame', required=False, type=int, default=3,
                        metavar='3 or 6', help='enter the reading frame of 3 or 6 (default is 3)')

arg = arg_parser.parse_args()


"""
general algo (notes for myself):
1. find all the indices of the start codons in each seq
2. find corresponding index of the stop codon for each start index
3. get all the orfs corresponding to the start and stop in a list?
4. if frame is 6 then do the same for the complement and compare with the single frame reading before returning the longest orf
"""


"""
function that finds the indices of all the start codons in the given sequence

input: sequence
output: list of indices where the start codons are found
"""
def find_starts(seq) -> list:

    start_indices = []
    cur_start = 0
    
    # index loops through the seq and returns the index of the first location where it finds ATG or it raises an exception if it doesn't find it
    try:
        cur_start = seq.index('ATG')
    except ValueError:
        return start_indices
    
    while cur_start < len(seq):
        start_indices.append(cur_start)

        # now we find a new start codon after the previous start codon index
        try:
            cur_start = seq.index('ATG', cur_start+1)
        except ValueError:
            return start_indices


"""
function that finds the index of the first stop codon in the sequence

input: sequence
output: index where the stop codon ends or if no stop codon returns -1
"""
def find_stop(seq) -> int:
    
    i = 0
    while i < len(seq)-2 and seq[i:i+3] not in ['TAG', 'TAA', 'TGA']:
        i += 3

    # check if the i is within the bounds of the sequence
    if i < len(seq) - 2:
        # return the index of the end of the stop codon
        return i+3

    return -1


"""
function that finds the indices of all the orfs of a given sequence

input: sequence
output: list of tuples where each tuple is the start and stop index of an orf in the sequence
"""
def get_orfs(seq):

    # get all the start indices of the sequence
    start_indices = find_starts(seq)
    stop_indices = []

    # loop through start indices and get the stop indices of each start index
    for start in start_indices:
        stop_ = find_stop(seq[start:])
        actual_stop = start + stop_
        stop_indices.append(actual_stop)
    
    orf = []
    # make a tuple of each start and stop and append to list
    for start, stop in zip(start_indices, stop_indices):
        orf.append((start, stop))
     
    return orf


"""
function to find the complement of a given sequence

input: sequence
output: string that is the complement to the given sequence
"""
def seq_comp(seq) -> str:
    # dict containing what we want to replace
    replace_dict = {
        'A' : 'T',
        'T' : 'A',
        'C' : 'G',
        'G' : 'C'
    }
    # the translation table converts dict into ascii numbers
    transtable = str.maketrans(replace_dict)
    # translate function uses the table to replace chars when found
    return seq.translate(transtable)


"""
function that finds the longest orf given a list of orfs

input: list of orfs (tuples of start and stop indices)
output: the length of the longest orf and the start and stop index
"""
def longest_orf(seq_orfs) -> int:
    # variables to keep track of the longest orf
    len_longest = -1
    longest_start = -1
    longest_stop = -1
    # loop through the orfs and find the longest one
    for orf in seq_orfs:
        start = orf[0]
        stop = orf[1]
        if stop-start > len_longest:
            len_longest = stop - start
            longest_start = start
            longest_stop = stop
    
    return len_longest, longest_start, longest_stop


"""
main loop
"""
for id, seq in read_record(arg.fasta_file):
    
    final_orfs = []
    # get the longest orf of normal frame
    orfs = get_orfs(seq)
    longest, start, stop = longest_orf(orfs)

    # append the longest orf to the final_orfs list
    if longest != -1:
        final_orfs.append(seq[start:stop])
    


    if arg.frame == 6:
        # get complement's longest orf
        complement = seq_comp(seq)
        orfs_c = get_orfs(complement)
        longest_c, start_c, stop_c = longest_orf(orfs_c)

        # append the complement's longest orf to the final_orfs list
        if longest_c != -1:
            final_orfs.append(complement[start_c:stop_c])

    # check to see if we have two orfs in the final_orfs list and print accordingly
    if len(final_orfs) == 2:
        print('>'+ id)
        print(max(final_orfs, key=len))
    
    elif len(final_orfs) == 1:
        print('>'+ id)
        print(final_orfs[0])

    else:
        print('>'+ id)
        print("No ORF Found")






# simple regex expression that does everything in one step (found online)

# for id, seq in read_record(arg.fasta_file):
#     cur_seq_all = re.findall(r'(?=(ATG(?:(?!TAA|TAG|TGA)...)*(?:TAA|TAG|TGA)))',seq)
#     if len(cur_seq_all) > 0:
#         print(f'>{id}' + '\n' + max(cur_seq_all, key=len))
