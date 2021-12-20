#!/usr/bin/env python3
from subprocess import check_output

# Given a comma-separated set of tuple of integer values,
# calculate the minimum, average, maximum readings, and total of the readings.
#
# Usage example:
#
#    ./test.py 2 instances/r_simple2.go
#
# For example, given this:
#
#    2, 0, 3
#    2, 3, 0
#
# the output would be:
#
#    2.0, 0.0, 0.0  <--- min values for each reading
#    2.0, 1.5, 1.5  <--- avg values
#    2.0, 3.0, 3.0  <--- max values
#    4.0, 3.0, 3.0  <--- total values
#
import datetime
import os
import random
import sys


counter = 0   # test case counter

inputfiles1 = []
inputfiles2 = []

NEWLINE = "\n"
START_PORT = 34117

def generate_header(SERVERS, BUSY_THRESHOLD, UPPER_THRESHOLD, P_ACCEPT):

	spaces = "\n".join(f"var s{i} Space" for i in range(1, SERVERS+2))
	new_spaces = "\n    ".join(f's{i} = NewSpace("tcp://localhost:{START_PORT+i}/s{i}")' for i in range(1, SERVERS+2)) 
	go_peers = "\n    ".join(f"go P{i}()" for i in range(2, SERVERS+2))


	return f"""
package main

import (
	"fmt"
	"os"
	"sync"
	"time"
	"math"
	"math/rand"
	. "github.com/pspaces/gospace"
)

const REQUESTS = {100*SERVERS} 
const SERVERS = {SERVERS}
const REQ_LOAD = 1 
const BUSY_THRESHOLD = {BUSY_THRESHOLD} 
const UPPER_THRESHOLD = {UPPER_THRESHOLD}


/************
REQUESTS = Number of requests that will be generated
SERVERS = Number of servers
REQ_LOAD = Load of an individual request
BUSY_THRESHOLD = Server declares to be "busy" when its load hits this threshold
UPPER_THRESHOLD = Server stops queueing requests when its load hits this threshold
**************/

const debug = false

// Probability with which a server accepts a request
// (Probability of completing a request is 1-P_ACCEPT)
const P_ACCEPT = {P_ACCEPT}

var l sync.Mutex
var wg sync.WaitGroup
{spaces}

// Counters for stats
var FORWARDED_REQUESTS int = 0
var STOLEN_REQUESTS int = 0
var COMPLETED_REQUESTS int = 0
// Logical timestamp to measure request acceptance time
var TIME int64 = 0
// Statistics about acceptance time
var totalTime int64 = 0
var minTime int64 = math.MaxInt64
var maxTime int64 = 0

func main() {{
	// n servers + 1 request generator
	wg.Add(SERVERS + 1)
	{new_spaces}

	fmt.Println("Start")
	go P1_Generator()
	{go_peers}

	wg.Wait()
	fmt.Println("Stop")
	fmt.Println("all,completed,local,forwarded,stolen,minTime,maxTime,avgTime,replicas,evictions,   maxMem")
	fmt.Fprintf(os.Stdout, "%d,%d,%d,%d,%d,%d,%d,%f,%d,%d,%d\\n",
		REQUESTS,
		COMPLETED_REQUESTS,
		COMPLETED_REQUESTS-STOLEN_REQUESTS,
		FORWARDED_REQUESTS,
		STOLEN_REQUESTS,
		minTime,
		maxTime,
		float64(totalTime)/float64(COMPLETED_REQUESTS),
		-1, -1, -1)
}}

"""

def generate_footer():
	return """
func delay() {
	time.Sleep(time.Duration(rand.Int63n(75))*time.Millisecond)
}

func biasedRandBool(bias float32) bool {
	// Assuming that 0 <= bias <= 1
	// Returns true with probability "bias" and
	// 0 with probability (1-bias).
	return rand.Float32() < bias
}

func log(format string, a ...interface{}) {
	if debug {
		fmt.Fprintf(os.Stdout, format+"\\n", a ...)
	}
}

func updateTimeStats (newTime int64) {
	totalTime += newTime
	if (newTime > maxTime) {
		maxTime = newTime
	}
	if (newTime < minTime) {
		minTime = newTime
	}
}

func getTIME() int64 {
	result := TIME
	TIME += 1
	return result
}
"""

def generate_P1(SERVERS):
	def ifblock(i): 
		return f"""
		if server == {i} {{
			s{i}.Put({i}, service, tstamp);
		}}
		"""

	return f"""
func P1_Generator() {{
	// This process only generates service requests.
	// A request is a pair (srv, svc) where
	// srv is the ID of a server (1 < srv <= SERVERS+1);
	// svc is the ID of the required service (1 < svc <= SERVERS+1).
	// Notice that server i is the default target for the i-th service.
	defer wg.Done()
	var i int;
	var service int;
	var server int;
	var tstamp int64;

	for i = 0; i<REQUESTS; i++ {{
		// Both service and server must be between 2 and SERVERS+2
		service = rand.Intn(SERVERS) + 2
		server = rand.Intn(SERVERS) + 2
		tstamp = getTIME()
		l.Lock()
		{NEWLINE.join(ifblock(i) for i in range(2, SERVERS+2))}
		l.Unlock()
		log("[1]\\tput req %d to server %d", service, server)
	}}
}}
"""

def generate_peer(i, SERVERS):
	def check_busy(j):
		return f"""
			if service == {j} {{
				s{j}.QueryP({j}, &s)
				if s == "busy" {{
					updateTimeStats(getTIME() - tstamp)
					stolenReq = true
					load += REQ_LOAD
				}} else {{
					s{j}.Put({j}, {j}, tstamp)
				}}
			}}
	"""

	return f"""
func P{i}() {{
	defer wg.Done()
	var tid int = {i}
	var load int = 0
	var service int
	var wasBusyBefore bool
	var isBusyNow bool
	var tstamp int64
	var server int

	steps := 0
	// The (steps < 10*REQUESTS) condition is just in case we
	// drop requests in the replicated case
	for COMPLETED_REQUESTS < REQUESTS && (steps < 10*REQUESTS) {{
		steps += 1
		l.Lock()
		service = -1
		wasBusyBefore = isBusyNow

		if load < UPPER_THRESHOLD && biasedRandBool(P_ACCEPT) {{
			s{i}.GetP(&server, &service, &tstamp)
			if service == tid {{
				updateTimeStats(getTIME() - tstamp)
				load += REQ_LOAD
				if server != tid {{
					// We have stolen a request for OUR service which was sent
					// to ANOTHER server. Only happens in the replicated case!
					STOLEN_REQUESTS += 1
					log("[%d]\\tstole req %d from server %d", tid, service, server)
				}} else {{
					log("[%d]\\taccepted req for service %d", tid, service)
				}}
			}} else {{
				// This request is for a service j != tid.
				// However, if server j declares to be busy,
				// steal the request.
				// (Notice that it doesnt't matter whether the "thief" is busy or not.)
				// If server j is not busy, simply forward the request to it.
				var s string
				stolenReq := false
				{NEWLINE.join(check_busy(j) for j in range(2, SERVERS+2))}
				

				isBusyNow = load >= BUSY_THRESHOLD
				if isBusyNow && !wasBusyBefore {{
					// Declare that this server is now busy
					s{i}.Put({i}, "busy")
					log("[%d]\\tnow busy: load %d >= %d",  tid, load, BUSY_THRESHOLD)
				}}
				// Update counters
				if service != -1 {{
					if stolenReq {{
						STOLEN_REQUESTS += 1
						log("[%d]\\tstole req %d from server %d", tid, service, server)
					}} else {{
						FORWARDED_REQUESTS += 1
						log("[%d]\\tfwded req %d to server %d", tid, service, service)
					}}
				}}
			}}
		}} else {{
			if load >= REQ_LOAD {{
				// Complete a request
				delay()
				log("[%d]\\tcompleting a request", tid)
				load -= REQ_LOAD
				if load < BUSY_THRESHOLD && (load+REQ_LOAD >= BUSY_THRESHOLD) {{
					// Declare that this server is no longer busy
					isBusyNow = false
					s{i}.GetP({i}, "busy")
					log("[%d]\\tno longer busy: load %d < %d",  tid, load, BUSY_THRESHOLD)
				}}
				COMPLETED_REQUESTS += 1
			}}
		}}
		l.Unlock()
	}}
}}
"""

# def prova(SERVERS=3, UPPER_THRESHOLD=4, BUSY_THRESHOLD=2, P_ACCEPT="0.75"):
# 	print(generate_header(SERVERS, UPPER_THRESHOLD, BUSY_THRESHOLD, P_ACCEPT))
# 	print(generate_P1(SERVERS))
# 	for i in range(2, SERVERS+2):
# 		print(generate_peer(i, SERVERS))
# 	print(generate_footer())


def generate_test_case(SERVERS, BUSY_THRESHOLD, UPPER_THRESHOLD, P_ACCEPT, REPLLIMIT, POLICY):
	global counter
	
	output = (
		generate_header(SERVERS, BUSY_THRESHOLD, UPPER_THRESHOLD, P_ACCEPT) +
		generate_P1(SERVERS) +
		"\n".join(generate_peer(i, SERVERS) for i in range(2, SERVERS+2)) +
		generate_footer())


	# # Save the test case (GoSpace version)

	filename = f"ws/test{counter}_N{SERVERS}_Tb{BUSY_THRESHOLD}_Tu{UPPER_THRESHOLD}_Pa{int(P_ACCEPT*100)}_R{REPLLIMIT}_RP{POLICY}.go"
	filename2 = filename.replace("/test", "/r_test")

	with open(filename, 'w') as outputfile:
		outputfile.write(output)

	# Generate the replicated version of the test case
	cmd = f'go run *.go -quiet -i {filename}'
	output = check_output(cmd, env=os.environ, shell=True).decode()


	# Add replication metrics (replicas, evictions)
	output = output.replace("-1, -1, -1)", "Getwcount(), GetEvictionCount(), GetReplicaMax())")


	if REPLLIMIT > 0:
		output = output.replace(
			"Replispace{Sp: Sp}",
			f'Replispace{{Sp: Sp, ReplLimit: {REPLLIMIT}, ReplacementPolicy: "{POLICY}"}}'
		)

	with open(filename2, 'w') as outputfile:
		outputfile.write(output)

	inputfiles1.append(filename)
	inputfiles2.append(filename2)
	counter += 1


def execute_test_cases(filenames,times):
	first = True

	inf = float("inf")

	avg = []
	low = []
	high = []
	total = []

	for filename in filenames:
		####################print('    '+filename)  # print name of each test case
		cmd = 'go run ' +filename

		# Invoke the same command several times and get the readings.
		for i in range(0,int(times)):
			output = check_output(cmd, env=os.environ, shell=True).decode()
			# stream = os.popen(cmd)
			# output = stream.read()

		####################print(output)   # print results of each run on each test case

			for line in output.splitlines():
				entries = line.split(',')

				if not entries[0].strip().isdigit(): continue

				if first:
					avg = [0.0] * len(entries)
					low = [inf] * len(entries)
					high = [0.0] * len(entries)
					total = [0.0] * len(entries)
					first = False

				for j in range(0,len(avg)):
					value = (
						float(entries[j])
						if "." in entries[j]
						else int(entries[j]))
					total[j] += value
					low[j] = float(min(low[j],value))
					high[j] = float(max(high[j],value))

	for i in range(0,len(avg)):
		avg[i] = float(total[i]) / (float(times)*len(filenames))

	# Generate output values (min,avg,max,total) to print out.
	values = '  min '

	for i in range(0,len(avg)): values += '%s, ' % '{0: >8.1f}'.format(low[i])

	values = values[:-2]
	values += '\n  avg '

	for i in range(0,len(avg)): values += '%8.1f, ' % (avg[i])

	values = values[:-2]
	values += '\n  max '

	for i in range(0,len(avg)): values += '%s, ' % '{0: >8.1f}'.format(high[i])

	values = values[:-2]
	values += '\n  tot '

	for i in range(0,len(avg)): values += '%s, ' % '{0: >8.1f}'.format(total[i])
	values = values[:-2]

	print(values)
	print()


def main(args, test=0):
	global inputfiles1, inputfiles2

	test_cases_per_configuration = 1
	simulations_per_test_case = 20

	# S =	no. of servers 
	# 		(keep in mind the total no. of processes will be S+1, 
	# 		namely the request generator)
	# B = 	Busy threshold
	# P = 	probability of accepting a request vs. handling an accepted request
	# U = 	Upper threshold (load at which the peer stops accepting requests)
	#		Here, we fix it to 150% of the busy threshold

	# ONLY USED IN THE REPLICATED PROGRAM:
	# R =	Replica limit (0 means no limit)
	# RP = 	Replacement policy ("fifo", "lru", or "random")

	if test == 0:
		# Test 1: Increasing busy threshold (no memory limit, fifo policy)
		S =		[  3, 	3, 	 3,   3]
		B = 	[  2, 	4,   6,   8]
		P =		[0.75, 0.75, 0.75, 0.75]
		R = 	[0] * len(S)
		RP = 	["fifo"] * len(S)
		U = 	[int(1.5*Tb) for Tb in B]

	elif test == 1:

		# Test 2: Increasing memory limit, fifo policy
		S =		[  3, 	3, 	 3,   3]
		B = 	[  2, 	2,   2,   2]
		P =		[0.75, 0.75, 0.75, 0.75]
		R = 	[  1, 	2, 	 4,	  8]
		RP = 	["fifo"] * len(S)
		U = 	[int(1.5*Tb) for Tb in B]

	elif test == 2:
		# Test 3: Increasing memory limit, lru policy
		S =		[  3, 	3, 	 3,   3]
		B = 	[  2, 	2,   2,   2]
		P =		[0.75, 0.75, 0.75, 0.75]
		R = 	[  1, 	2, 	 4,	  8]
		RP = 	["lru"] * len(S)
		U	= 	[int(1.5*Tb) for Tb in B]

	else:
		# Test 4: Increasing memory limit, random policy
		S =		[  3, 	3, 	 3,   3]
		B = 	[  2, 	2,   2,   2]
		P =		[0.75, 0.75, 0.75, 0.75]
		R = 	[  1, 	2, 	 4,	  8]
		RP = 	["random"] * len(S)
		U	= 	[int(1.5*Tb) for Tb in B]



	print('Overall configurations:        %d' % len(S))
	print('Test cases per configuration:  %d' % test_cases_per_configuration)
	print('Simulations per test case:     %d' % simulations_per_test_case)
	runs = test_cases_per_configuration*simulations_per_test_case*2
	print('Simulations per configuration: %d' % runs)

	print('')

	for a in range(0,len(S)):

		print(f"Configuration no. {a+1} (N={S[a]:01} Tb={B[a]:02} Tu={U[a]:03} Pa={int(P[a]*100):02}% R={R[a]:02} RP={RP[a]})")
		for b in range(0,test_cases_per_configuration):
			generate_test_case(S[a],B[a],U[a],P[a],R[a],RP[a])

		header = "         total,  handled,    local,    fwded,   stolen     t_min,    t_max,    t_avg,*replicas, *evicted,   maxMem"
			
		print('  Without replication:')
		print(header)
		execute_test_cases(inputfiles1,simulations_per_test_case)

		print('  With replication:')
		print(header)
		execute_test_cases(inputfiles2,simulations_per_test_case)


		inputfiles1 = []
		inputfiles2 = []

	print (' * = this measure only counts for replicated programs')
	# print ('** = this measure only to be used as a safety check for test case generation')

if __name__ == "__main__":
	# generate_test_case(3, 2)
	# execute_test_cases(inputfiles1, 5)
	# prova()
	for i in range(4):
		print(f"Executing experimental suite {i+1}...")
		main(sys.argv[0:], i)



