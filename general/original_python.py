from decimal import Decimal
from typing import NamedTuple
import unittest

"""
The problem

At Cedar, we aim to find the best messaging that will motivate the patient to engage with us.

You have three sets of data:

- A list of patient bills.

- A list of bill communications sent to patients. channel_type can be "text", "email", or "paper mail"

- A list of payments that Cedar received from patients.

Your goal is to produce output that allows us to see which channel type is most effective at driving payments.

We ask that you:

1) Define "effectiveness". This is open-ended, but we ask that you spend no more than 5 minutes choosing a metric. We are not looking for something complex, instead please choose a metric that you can implement in 30 minutes or less.

2) Implement your metric.

3) Write tests for your code to prove it works as expected. Ideally we'd like your tests to cover at least three different scenarios. If you are running out of time, your interviewer will ask you to describe your testing approach during the last 5-10 minutes of the interview.

You are given some initial test data for your development, as well as a sample unit test.

You may ask questions to your interviewer, and you may use any resources on the Internet (except AI coding assistants) as needed while coding!
"""

class Bill(NamedTuple):
	patient_id: str
	amount: Decimal
	created_timestamp: int

class Communication(NamedTuple):
	patient_id: str
	channel_type: str
	sent_timestamp: int

class Payment(NamedTuple):
	patient_id: str
	amount: Decimal
	paid_timestamp: int

BILLS = [
	Bill("p1", Decimal(100), 160310 * 10000),
	Bill("p2", Decimal(500), 160320 * 10000),
	Bill("p2", Decimal(350), 160320 * 10000),
	Bill("p3", Decimal(450), 160340 * 10000),
	Bill("p1", Decimal(100), 160500 * 10000),
]

COMMUNICATIONS = [
	Communication("p1", "email", 160325 * 10000),
	Communication("p2", "text", 160320 * 10000),
	Communication("p2", "paper_mail", 160330 * 10000),
	Communication("p2", "paper_mail", 160340 * 10000),
	Communication("p3", "paper_mail", 160345 * 10000),
]

PAYMENTS = [
	Payment("p2", Decimal(350), 160345 * 10000),
	Payment("p3", Decimal(450), 160360 * 10000),
	Payment("p1", Decimal(100), 160600 * 10000),
]

def calculate_effectiveness(bills, communications, payments):
	emails = 0
	texts = 0
	paper = 0

	for bill in BILLS:
		patient = bill[0]
		bill_timestamp = bill[2]
		for payment in PAYMENTS:
			payment_timestamp = payment[2]
			if payment[0] == patient and bill_timestamp <= payment[2]:
			# then we want to collect all the communications
			# between the timestamp and the payment timestamp
			# and we want to score them
				for communication in COMMUNICATIONS:
					communication_timestamp = communication[2]
					if patient == communication[0] and bill_timestamp <= communication_timestamp and communication_timestamp <= payment_timestamp:
					communication_style = communication[1]
					if communication_style == "text":
						texts += 1
					elif communication_style == "paper_mail":
						paper += 1
					else:
						emails += 1
	
	return {
	"email": emails,
	"text": texts,
	"paper_mail": paper
	}

# ----------------------------------------------------------------------

# Test Cases

class CommunicationEffectivenessTestCase(unittest.TestCase):

def test_calculate_effectiveness(self):

bills = BILLS

communications = COMMUNICATIONS

payments = PAYMENTS

effectiveness = calculate_effectiveness(bills, communications, payments)

self.assertEqual(effectiveness["email"], 2)

self.assertEqual(effectiveness["text"], 2)

self.assertEqual(effectiveness["paper_mail"], 3)

  

# ----------------------------------------------------------------------

# Run all of the tests (do not modify or delete this line)

# ----------------------------------------------------------------------

unittest.main()

"""
We want to calculate how effective a communication is in getting a bill paid

and this could be across patients in general

1000 patients could respond better when a text is sent we don't know yet.

effectivenmess how much time between sending the communication and a payment being received. However, this may not be related but we have to

in a way guess, because we'll see other behaviors for other pateitns

not receiving that same communication

overall, we want to calculate what the difference in time is between

what communications are sent to the patient vs. when the payment was received and then score the communication methods according to that time

let's say we create a bill, the patient probably doesn't know about it yet, until they get the communication.

p1 has a bill for 100, and then 1 time unit later, an email is sent, and then the bill is paid immediately, i.e. the difference between the communication is 0, email would :"win"

p2 has a bill 850, and then 0 units between text, 10 units for paper, and 20 units for paper, and the payment is received at 25 units, 25 for email, and 15 for paper; p2 hs 1 text, and 2 paper mails.

between p1 and p2, the average email was .5, text .5, and 1 paper mail, for avg email, p2

each communication has a credit score of 1, email text paper

for patient,

emails = 0
texts = 0
paper = 0

for bill in BILLs:
	patient = bill[0]
	bill_timestamp = bill[2]
	for payments in PAYMENTS:
		payment_timestamp = payment[2]
		
		if payment[0] == patient and bill_timestamp <= payment[2]:
		
			# then we want to collect all the communications
			
			# between the timestamp and the payment timestamp
			
			# and we want to score them
			
			for communication in COMMUNICATIONS:
			
			communication_timestamp = communication[2]
			
			if patient == communication[0] and bill_timestamp <= communication_timestamp and communication_timestamp <= payment_timestamp:
				
				communication_style = communication[1]
				
				if communication_style == "text":
				
				texts += 1
				
				elif communication_style == "paper_mail":
				
				paper += 1
				
				else:
				
				emails += 1
	
"""