package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/AlanAloha/read_record"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type mm_model struct {
	name  string
	k     int
	size  int
	score map[string]float64
}

type len_model struct {
	name  string
	size  int
	score []float64
	tail  float64
}

type pwm_model struct {
	name  string
	size  int
	score [][]float64
}

type feature struct {
	seq   string
	beg   int
	end   int
	score float64
}

type mRNA struct {
	seq     string
	beg     int
	end     int
	exons   []feature
	introns []feature
	atg     int
	score   float64
}

type isozone struct {
	dons     int
	accs     int
	trials   int
	isoforms int
	mrnas    []mRNA
}

func complexity(mrnas []mRNA) float64 {
	total := 0.0
	p := make([]float64, len(mrnas))
	for i, mrna := range mrnas {
		w := math.Pow(2, mrna.score)
		total += w
		p[i] = w
	}
	for i := range p {
		p[i] /= total
	}

	h := 0.0
	for i := range p {
		if p[i] > 0 {
			h -= p[i] * math.Log2(p[i])
		}
	}

	return h
}

func mm_cache(mm mm_model, seq string) []float64 {
	score := make([]float64, len(seq))
	for i := mm.k; i < len(seq)-mm.k; i++ {
		cur_kmer := seq[i:i+mm.k]
		if key, ok := mm.score[cur_kmer]; ok {
			_ = key
			score[i] = score[i-1] + mm.score[seq[i:i+mm.k]]
		} else {
			score[i] = score[i-1]
		}
	}

	return score
}

func score_emm(mm mm_model, mm_cache []float64, mrna mRNA) float64 {
	score := 0.0
	for _, exon := range mrna.exons {
		beg := exon.beg
		end := exon.end
		s := mm_cache[end-mm.k+1] - mm_cache[beg-1]
		score += s
	}
	return score
}

func score_imm(mm mm_model, mm_cache []float64, mrna mRNA, dpwm pwm_model,
	apwm pwm_model) float64 {
	score := 0.0
	for _, intron := range mrna.introns {
		beg := intron.beg + dpwm.size
		end := intron.end - apwm.size
		s := mm_cache[end-mm.k+1] - mm_cache[beg-1]
		score += s
	}
	return score
}

func read_mm(file string) mm_model {
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	var mm mm_model
	score := make(map[string]float64)
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, " ")
		if f[0] == "%" {
			mm.name = f[2]
			size, _ := strconv.Atoi(f[3])
			mm.size = size
		} else if len(f) == 2 {
			kmer := f[0]
			if mm.k == 0 {
				mm.k = len(kmer)
			}
			prob, _ := strconv.ParseFloat(f[1], 64)
			score[kmer] = log2score(prob)
		}
	}

	mm.score = score

	return mm
}

func find_tail(val float64, size int) float64 {
	lo := 0.0
	hi := 1000.0
	x := float64(size)
	var m float64

	for hi-lo > 1 {
		m = (hi + lo) / 2.0
		p := 1.0 / m
		f := math.Pow(1.0-p, x-1.0) * p
		if f < val {
			lo += (m - lo) / 2.0
		} else {
			hi -= (hi - m) / 2.0
		}
	}

	return m
}

func read_len(file string) len_model {
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	var len_model len_model
	score := []float64{}
	var size int
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, " ")
		if f[0] == "%" {
			len_model.name = f[2]
			size, _ = strconv.Atoi(f[3])
			len_model.size = size
		} else {
			s, _ := strconv.ParseFloat(f[0], 64)
			score = append(score, s)
		}
	}

	len_model.tail = find_tail(score[size-1], size)
	expect := 1 / float64(size)
	for i := range score {
		score[i] = math.Log2(score[i] / expect)
	}
	len_model.score = score
	return len_model
}

func score_len(len_model len_model, length int) float64 {
	if length < 0 {
		fmt.Println("Error in length model... Length < 0")
		os.Exit(4)
	}

	if length >= len_model.size {
		p := 1 / len_model.tail
		q := math.Pow(1-p, float64(length-1)) * p
		expect := 1 / float64(len_model.size)
		s := math.Log2(q / expect)
		return s
	} else {
		return len_model.score[length]
	}
}

func score_elen(len_model len_model, mrna mRNA) float64 {
	score := 0.0
	for _, exon := range mrna.exons {
		length := exon.end - exon.beg + 1
		s := score_len(len_model, length)
		score += s
	}
	return score
}

func score_ilen(len_model len_model, mrna mRNA) float64 {
	score := 0.0
	for _, intron := range mrna.introns {
		length := intron.end - intron.beg + 1
		s := score_len(len_model, length)
		score += s
	}
	return score
}

func log2score(prob float64) float64 {
	if prob == 0 {
		return -100
	}
	return math.Log2(prob / 0.25)
}

func read_pwm(file string) pwm_model {
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	var pwm pwm_model
	score := [][]float64{}
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Split(line, " ")
		if f[0] == "%" {
			pwm.name = f[2]
			pwm.size, _ = strconv.Atoi(f[3])
		} else {
			cur_row := make([]float64, 4)
			for i := range cur_row {
				p, _ := strconv.ParseFloat(f[i], 64)
				cur_row[i] = log2score(p)
			}
			score = append(score, cur_row)
		}
	}
	pwm.score = score

	return pwm
}

func score_pwm(pwm pwm_model, seq string, pos int) float64 {
	score := 0.0
	for i := 0; i < pwm.size; i++ {
		bp := string(seq[i+pos])
		if bp == "A" || bp == "a" {
			score += pwm.score[i][0]
		} else if bp == "C" || bp == "c" {
			score += pwm.score[i][1]
		} else if bp == "G" || bp == "g" {
			score += pwm.score[i][2]
		} else {
			score += pwm.score[i][3]
		}
	}
	return score
}

func score_apwm(pwm pwm_model, mrna mRNA) float64 {
	score := 0.0
	for i := 0; i < len(mrna.introns); i++ {
		feat := mrna.introns[i]
		s := score_pwm(pwm, feat.seq, feat.end-pwm.size+1)
		score += s
	}
	return score
}

func score_dpwm(pwm pwm_model, mrna mRNA) float64 {
	score := 0.0
	for i := 0; i < len(mrna.introns); i++ {
		feat := mrna.introns[i]
		s := score_pwm(pwm, feat.seq, feat.beg)
		score += s
	}
	return score
}

func no_dup(cur int, sites []int) bool {
	for _, site := range sites {
		if site == cur {
			return false
		}
	}

	return true
}

func canonical(beg int, end int, seq string) bool {
	if seq[beg:beg+2] != "GT" {
		return false
	}
	if seq[end-1:end+1] != "AG" {
		return false
	}
	return true
}

func gff_sites(seq string, file string, dons *[]int, accs *[]int, gtag bool) {
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	gff := bufio.NewScanner(fh)
	for gff.Scan() {
		line := gff.Text()
		f := strings.Split(line, "\t")
		src := f[1]
		feat := f[2]
		beg, _ := strconv.Atoi(f[3])
		end, _ := strconv.Atoi(f[4])
		beg -= 1
		end -= 1
		strd := f[6]
		if src != "RNASeq_splice" {
			continue
		}
		if strd != "+" {
			continue
		}
		if feat != "intron" {
			continue
		}
		if gtag && !canonical(beg, end, seq) {
			continue
		}
		if no_dup(beg, *dons) {
			*dons = append(*dons, beg)
		}
		if no_dup(end, *accs) {
			*accs = append(*accs, end)
		}
	}
	sort.Ints(*dons)
	sort.Ints(*accs)
}

func gtag_sites(seq string, beg int, end int, dons *[]int, accs *[]int) {
	for i := beg; i < end; i++ {
		if seq[i] == 'G' && seq[i+1] == 'T' {
			*dons = append(*dons, i)
		}
		if seq[i] == 'A' && seq[i+1] == 'G' {
			*accs = append(*accs, i+1)
		}
	}
}

func isoforms(seq string, emin int, imin int, smax int, gen int, gff string) *isozone {
	var iso isozone
	seqlen := len(seq)
	var dons []int
	var accs []int

	if gff != "" {
		gff_sites(seq, gff, &dons, &accs, true)
	} else {
		gtag_sites(seq, gen+emin, seqlen-gen-emin, &dons, &accs)
	}

	var nsites int
	if len(dons) < len(accs) {
		nsites = len(dons)
	} else {
		nsites = len(accs)
	}
	if nsites > smax {
		nsites = smax
	}

	trials := 0
	forms := 0
	var mrnas []mRNA
	for n := 1; n <= nsites; n++ {
		dcombos := get_combinations(dons, n)
		acombos := get_combinations(accs, n)

		for _, dsites := range dcombos {
			for _, asites := range acombos {
				if len(dsites) != len(asites) {
					fmt.Println("number of acceptor sites and number of donor sites for combination aren't equal")
					os.Exit(3)
				}
				trials += 1

				if short_intron(dsites, asites, imin) {
					continue
				}
				if short_exon(dsites, asites, seqlen, gen, emin) {
					continue
				}

				forms += 1
				mrna := build_mRNA(seq, gen, seqlen-gen, dsites, asites)
				mrnas = append(mrnas, mrna)
			}
		}

	}

	iso.dons = len(dons)
	iso.accs = len(accs)
	iso.trials = trials
	iso.isoforms = forms
	iso.mrnas = mrnas
	return &iso
}

func build_mRNA(seq string, beg int, end int, dsites []int, asites []int) mRNA {
	mrna := mRNA{seq: seq, beg: beg, end: end, score: 0}

	if len(dsites) == 0 {
		f := feature{seq: seq, beg: beg, end: end, score: 0}
		mrna.exons = append(mrna.exons, f)
		return mrna
	}

	//introns
	for i := range dsites {
		ibeg := dsites[i]
		iend := asites[i]
		f := feature{seq: seq, beg: ibeg, end: iend, score: 0}
		mrna.introns = append(mrna.introns, f)
	}

	//1st exon
	ei := feature{seq: seq, beg: beg, end: dsites[0] - 1, score: 0}
	mrna.exons = append(mrna.exons, ei)
	//internal exons
	for i := 1; i < len(asites); i++ {
		exon_beg := asites[i-1] + 1
		exon_end := dsites[i] - 1
		ex := feature{seq: seq, beg: exon_beg, end: exon_end, score: 0}
		mrna.exons = append(mrna.exons, ex)
	}
	//last exon
	el := feature{seq: seq, beg: asites[len(asites)-1] + 1, end: end, score: 0}
	mrna.exons = append(mrna.exons, el)

	return mrna
}

func short_intron(dsites []int, asites []int, imin int) bool {
	for i := range dsites {
		ilen := asites[i] - dsites[i] + 1
		if ilen < imin {
			return true
		}
	}
	return false
}

func short_exon(dsites []int, asites []int, seqlen int, gen int, emin int) bool {
	//first exon
	exon_beg := gen + 1
	exon_end := dsites[0] - 1
	elen := exon_end - exon_beg + 1
	if elen < emin {
		return true
	}

	//internal exon(s)
	for i := 1; i < len(asites); i++ {
		exon_beg = asites[i-1] + 1
		exon_end = dsites[i] - 1
		elen = exon_end - exon_beg
		if elen < emin {
			return true
		}
	}

	//last exon
	exon_beg = asites[len(asites)-1] + 1
	exon_end = seqlen - gen - 1
	elen = exon_end - exon_beg + 1
	if elen < emin {
		return true
	}

	return false
}

func get_combinations(arr []int, n int) [][]int {
	tmp := make([]int, n)
	l := len(arr)
	combo := [][]int{}
	combination(&combo, arr, n, tmp, 0, l-1, 0)

	return combo
}

func combination(combo *[][]int, arr []int, n int, tmp []int, start int, end int, index int) {
	if index == n {
		carrier := make([]int, n)
		for i := range tmp {
			carrier[i] = tmp[i]
		}
		*combo = append(*combo, carrier)
		return
	}

	i := start
	for i <= end && end-i+1 >= n-index {
		tmp[index] = arr[i]
		combination(combo, arr, n, tmp, i+1, end, index+1)
		i += 1
	}
}

func read_seq(fasta string) (string, string) {
	var seq string
	var idn string
	records := read_record.Read_record(fasta)
	record_num := 0
	for records.Next() {
		if record_num > 0 {
			fmt.Println("fasta file contains more than 1 sequence!")
			os.Exit(2)
		}
		record := records.Record()
		seq = record.Seq
		idn = record.Id
		record_num += 1
	}

	return idn, seq
}

func main() {
	flag_fasta := flag.String("fa", "", "path to fasta file (required)")
	flag_emin := flag.Int("min_exon", 25, "minimum exon length [25]")
	flag_imin := flag.Int("min_intron", 35, "minimum intron length [35]")
	flag_smax := flag.Int("max_splice", 3, "maximum number of splices [3]")
	flag_gen := flag.Int("flank", 99, "genomic flank lengths [99]")
	flag_gff := flag.String("introns", "", "use introns from GFF file")
	flag_head := flag.Int("limit", 20, "limit report to this many isoforms [20]")
	flag_apwm := flag.String("apwm", "", "use acceptor pwm")
	flag_dpwm := flag.String("dpwm", "", "use donor pwm")
	flag_elen := flag.String("elen", "", "use exon length model")
	flag_ilen := flag.String("ilen", "", "use intron length model")
	flag_emm := flag.String("emm", "", "use exon markov model")
	flag_imm := flag.String("imm", "", "use intron markov model")
	flag_ic := flag.Float64("icost", 0.0, "cost for each intron")
	flag_wapwm := flag.Float64("wapwm", 1.0, "weight of acceptor pwm [1.0]")
	flag_wdpwm := flag.Float64("wdpwm", 1.0, "weight of donor pwm [1.0]")
	flag_welen := flag.Float64("welen", 1.0, "weight of exon length [1.0]")
	flag_wilen := flag.Float64("wilen", 1.0, "weight of intron length [1.0]")
	flag_wemm := flag.Float64("wemm", 1.0, "weight of exon markov model [1.0]")
	flag_wimm := flag.Float64("wimm", 1.0, "weight of intron markov model [1.0]")

	flag.Parse()

	if *flag_fasta == "" {
		flag.Usage()
		os.Exit(1)
	}

	emin := *flag_emin
	imin := *flag_imin
	smax := *flag_smax
	gen := *flag_gen
	gff := *flag_gff
	apwm := *flag_apwm
	dpwm := *flag_dpwm
	elen := *flag_elen
	ilen := *flag_ilen
	emm := *flag_emm
	imm := *flag_imm
	ic := *flag_ic
	head := *flag_head
	wapwm := *flag_wapwm
	wdpwm := *flag_wdpwm
	welen := *flag_welen
	wilen := *flag_wilen
	wemm := *flag_wemm
	wimm := *flag_wimm

	// Find all isoforms
	idn, seq := read_seq(*flag_fasta)
	iso := isoforms(seq, emin, imin, smax, gen, gff)
	if head > iso.isoforms {
		head = iso.isoforms
	}

	// Load models
	var accpwm pwm_model
	var donpwm pwm_model
	var exolen len_model
	var intlen len_model
	var exomm mm_model
	var intmm mm_model
	var ecache []float64
	var icache []float64

	if apwm != "" {
		accpwm = read_pwm(apwm)
	} else {
		_ = accpwm
	}
	if dpwm != "" {
		donpwm = read_pwm(dpwm)
	} else {
		_ = donpwm
	}
	if elen != "" {
		exolen = read_len(elen)
	} else {
		_ = exolen
	}
	if ilen != "" {
		intlen = read_len(ilen)
	} else {
		_ = intlen
	}
	if emm != "" {
		exomm = read_mm(emm)
		ecache = mm_cache(exomm, seq)
	} else {
		_ = exomm
		_ = ecache
	}
	if imm != "" {
		intmm = read_mm(imm)
		icache = mm_cache(intmm, seq)
	} else {
		_ = intmm
		_ = icache
	}

	// Score models
	for i, mrna := range iso.mrnas {
		if apwm != "" {
			iso.mrnas[i].score += score_apwm(accpwm, mrna) * wapwm
		}
		if dpwm != "" {
			iso.mrnas[i].score += score_dpwm(donpwm, mrna) * wdpwm
		}
		if elen != "" {
			iso.mrnas[i].score += score_elen(exolen, mrna) * welen
		}
		if ilen != "" {
			iso.mrnas[i].score += score_ilen(intlen, mrna) * wilen
		}
		if emm != "" {
			iso.mrnas[i].score += score_emm(exomm, ecache, mrna) * wemm
		}
		if imm != "" {
			iso.mrnas[i].score += score_imm(intmm, icache, mrna, donpwm, accpwm) * wimm
		}
		iso.mrnas[i].score -= float64(len(mrna.introns)) * ic
	}

	// Sort isoforms
	sort.SliceStable(iso.mrnas, func(i, j int) bool {
		return iso.mrnas[i].score > iso.mrnas[j].score
	})

	fmt.Printf("# name: %s\n", idn)
	fmt.Printf("# length: %d\n", len(seq))
	fmt.Printf("# donors: %d\n", iso.dons)
	fmt.Printf("# acceptors: %d\n", iso.accs)
	fmt.Printf("# trials: %d\n", iso.trials)
	fmt.Printf("# isoforms: %d\n", iso.isoforms)
	fmt.Printf("# complexity: %.4f\n", complexity(iso.mrnas))

	// Get probability of each isoform from score
	max_score := iso.mrnas[0].score
	p := make([]float64, head)
	total := 0.0
	for i := 0; i < head; i++ {
		mrna := iso.mrnas[i]
		w := math.Pow(2, mrna.score-max_score)
		total += w
		p[i] = w
	}
	for i := range p {
		p[i] /= total
	}

	idnf := strings.Split(idn, " ")
	chrom := idnf[0]
	gbeg := iso.mrnas[0].exons[0].beg + 1
	gend := iso.mrnas[0].exons[len(iso.mrnas[0].exons)-1].end + 1
	fmt.Printf("%s\tisoformer\tgene\t%d\t%d\t.\t+\t.\tID=Gene-%s\n\n",
		chrom, gbeg, gend, chrom)

	for i := 0; i < head; i++ {
		mrna := iso.mrnas[i]
		// mRNA title
		fmt.Printf("%s\tisoformer\tmRNA\t%d\t%d\t%.4g\t+\t.\tID=tx-%s-%d;Parent=Gene-%s\n",
			chrom, mrna.beg+1, mrna.end+1, p[i], chrom, i+1, chrom)

		// exons
		for _, exon := range mrna.exons {
			fmt.Printf("%s\tisoformer\texon\t%d\t%d\t%.4g\t+\t.\tParent=tx-%s-%d\n",
				chrom, exon.beg+1, exon.end+1, p[i], chrom, i+1)
		}

		// introns
		for _, intron := range mrna.introns {
			fmt.Printf("%s\tisoformer\tintron\t%d\t%d\t%.4g\t+\t.\tParent=tx-%s-%d\n",
				chrom, intron.beg+1, intron.end+1, p[i], chrom, i+1)
		}

		fmt.Println()
	}
}
