#!/usr/bin/env python2
from __future__ import print_function

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

# VAL = value, ADDR = address, LOCAL = local process, TARGET = target process (depending on the address)
block1 = '	delay()\n'\
	     '	l.Lock()\n'\
	     '	value = <VAL>\n'\
	     '	s<TARGET>.Put(<ADDR>,value)\n'\
	     '	log("p<LOCAL> w:%d", value)\n'\
	     '	countwrites(value,<LOCAL>,<TARGET>)\n'\
	     '	l.Unlock()'

block2 = '	delay()\n'\
	     '	l.Lock()\n'\
	     '	value = -1\n'\
	     '	s<TARGET>.QueryP(<ADDR>,&value)\n'\
	     '	log("p<LOCAL> r:%d", value)\n'\
	     '	countreads(value,<LOCAL>,<TARGET>)\n'\
	     '	l.Unlock()'


'''
	n = no. of processes (must match the template for now)
    m = overall memory size (must be a multiple of n)
    o = no. of operations per process
    p = (expected) percentace of Put operations (as opposite to QueryP)
'''
#def generate_test_case(n=32, m=256, o=16, p=100):
#def generate_test_case(n=32, m=512, o=16, p=100): # to use n=32,m=512,o=16
def generate_test_case(n=4, m=32, o=16, p=100):   # to use n=4,m=32,o=16
#def generate_test_case(n=4, m=64, o=16, p=100):   # to use n=4,m=64,o=16
	global counter

	#with open('instances/template_n32.go', 'r') as template: # for n=32
	with open('instances/template.go', 'r') as template:      # for n=4
		output = template.read()

	random.seed(datetime.datetime.now())

	# For each process
	for j in range(0,n):
		local = j+1

		# Guess o operations with their address
		for i in range(0,o):
			addr = random.randint(1,m)   # random address, from 1 to m
			target = ((addr-1)/(m/n))+1      # destination space (node 1 stores addresses from 1 to m/n, and so on)
			value = ((addr-1)/(m/n))+1       # same range as destination address (but irrelevant for now)
			op = 'Put' if random.randint(1,100)<=p else 'QueryP'

			#print("operation %s, addr %d, target %d" % (op, addr, target))

			if op == 'Put': block = block1
			else:           block = block2

			block = block.replace('<ADDR>', '%d' %addr)
			block = block.replace('<LOCAL>', '%d' %local)
			block = block.replace('<VAL>', '%d' %value)
			block = block.replace('<TARGET>', '%d' %target)

			key = '	<P%d>' %local
			output = output.replace(key, block+'\n\n'+key)

		output = output.replace('\n\n'+key, '')

	# Save the test case (GoSpace version)
	c1 = str(counter).zfill(3)
	n1 = str(n).zfill(1)
	m1 = str(m).zfill(2)
	o1 = str(o).zfill(3)
	p1 = str(p).zfill(3)

	filename = 'instances/test%s_n%s_m%s_o%s_p%s.go' %(c1,n1,m1,o1,p1)
	filename2 = 'instances/r_test%s_n%s_m%s_o%s_p%s.go' %(c1,n1,m1,o1,p1)

	with open(filename, 'w') as outputfile:
		outputfile.write(output)

	# Generate the replicated version of the test case
	cmd = 'go run *.go --quiet -i %s' % (filename)

	output = os.popen(cmd).read()
	output = output.replace('writes_replicated = 0', 'writes_replicated = Getwcount()-writes_local')

	with open(filename2, 'w') as outputfile:
		outputfile.write(output)

	inputfiles1.append(filename)
	inputfiles2.append(filename2)
	counter += 1


def execute_test_cases(filenames,times):
	first = True

	avg = []
	low = []
	high = []
	total = []

	for filename in filenames:
		####################print('    '+filename)  # print name of each test case
		cmd = 'go run ' +filename

		# Invoke the same command several times and get the readings.
		for i in range(0,int(times)):
			stream = os.popen(cmd)
			output = stream.read()

		####################print(output)   # print results of each run on each test case

			for line in output.splitlines():
				entries = line.split(',')

				if not entries[0].strip().isdigit(): continue

				if first:
					avg = [0.0] * len(entries)
					low = [sys.maxint] * len(entries)
					high = [0.0] * len(entries)
					total = [0.0] * len(entries)
					first = False

				for j in range(0,len(avg)):
					value = int(entries[j])
					total[j] += value
					low[j] = float(min(low[j],value))
					high[j] = float(max(high[j],value))

	for i in range(0,len(avg)): avg[i] = float(total[i]) / (float(times)*len(filenames))

	# Generate output values (min,avg,max,total) to print out.
	values = '  min '

	for i in range(0,len(avg)): values += '%s, ' % '{0: >8}'.format(low[i])

	values = values[:-2]
	values += '\n  avg '

	for i in range(0,len(avg)): values += '%8.1f, ' % (avg[i])

	values = values[:-2]
	values += '\n  max '

	for i in range(0,len(avg)): values += '%s, ' % '{0: >8}'.format(high[i])

	values = values[:-2]
	values += '\n  tot '

	for i in range(0,len(avg)): values += '%s, ' % '{0: >8}'.format(total[i])
	values = values[:-2]

	print(values)
	print('')


def main(args):
	global inputfiles1, inputfiles2

	test_cases_per_configuration = 10
	simulations_per_test_case = 10

	#n = no. of processes (must match the template for now)
    #m = overall memory size (must be a multiple of n)
    #o = no. of operations per process
    #p = (expected) percentace of Put operations (as opposite to QueryP)

    # Test 1:
    #
    #     +/- read/write ratio
    #

 	# experiment 1: the scenario with n=4, m=32, o=16

	n = [ 4, 4,  4,  4,  4,  4,  4,  4,  4]
	m = [32, 32, 32, 32, 32, 32, 32, 32, 32]
	o = [16, 16, 16, 16, 16, 16, 16, 16, 16]
	p = [10, 20, 30, 40, 50, 60, 70, 80, 90]

	# experiment 2: the scenario with n=4, m=64, o=16

	# n = [ 4, 4,  4,  4,  4,  4,  4,  4,  4]
	# m = [64, 64, 64, 64, 64, 64, 64, 64, 64]
	# o = [16, 16, 16, 16, 16, 16, 16, 16, 16]
	# p = [10, 20, 30, 40, 50, 60, 70, 80, 90]

	 # experiment 3: the scenario with n=32, m=256, o=16

	# n = [32, 32, 32, 32, 32, 32, 32, 32, 32]
	# m = [256, 256, 256, 256, 256, 256, 256, 256, 256]
	# o = [16, 16, 16, 16, 16, 16, 16, 16, 16]
	# p = [10, 20, 30, 40, 50, 60, 70, 80, 90]


	 # experiment 4: the scenario with n=32, m=512, 0=16

	# n = [32, 32, 32, 32, 32, 32, 32, 32, 32]
	# m = [512, 512, 512, 512, 512, 512, 512, 512, 512]
	# o = [16, 16, 16, 16, 16, 16, 16, 16, 16]
	# p = [10, 20, 30, 40, 50, 60, 70, 80, 90]


	print('Overall configurations:        %d' % len(n))
	print('Test cases per configuration:  %d' % test_cases_per_configuration)
	print('Simulations per test case:     %d' % simulations_per_test_case)
	runs = test_cases_per_configuration*simulations_per_test_case*2
	print('Simulations per configuration: %d' % runs)

	print('')

	for a in range(0,len(n)):
		n1 = str(n[a]).zfill(1)
		m1 = str(m[a]).zfill(2)
		o1 = str(o[a]).zfill(3)
		p1 = str(p[a]).zfill(3)

		print('Configuration no. %d (n=%s m=%s o=%s p=%s)' %(a+1,n1,m1,o1,p1))
		for b in range(0,test_cases_per_configuration):
			generate_test_case(n[a],m[a],o[a],p[a])

		print('  Without replication:')
		print('        *loc w,   *rem w,   repl w,   *loc r,   *rem r,  **tot w, **succ r, **fail r')
		execute_test_cases(inputfiles1,simulations_per_test_case)

		print('        *loc w,   *rem w,   repl w,   *loc r,   *rem r,  **tot w, **succ r, **fail r')
		print('  With replication:')
		execute_test_cases(inputfiles2,simulations_per_test_case)


		inputfiles1 = []
		inputfiles2 = []

	print (' * = this measure only counts for non-replicated programs')
	print ('** = this measure only to be used as a safety check for test case generation')

if __name__ == "__main__":
	main(sys.argv[0:])



