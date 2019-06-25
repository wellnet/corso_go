package main

import (
	"testing"
)

var testText = `The Doom 1 shareware WAD file is (C) Copyright id Software.

           LIMITED USE SOFTWARE LICENSE AGREEMENT

This Limited Use Software License Agreement (the "Agreement") is a legal
agreement between you, the end-user, and Id Software, Inc. ("ID").  By
continuing the installation of this game program, by loading or running
the game, or by placing or copying the game program onto your computer
hard drive, you are agreeing to be bound by the terms of this Agreement.

ID SOFTWARE LICENSE

   1.   Grant of License.  ID grants to you the right to use the
Id Software game program (the "Software"), which is the shareware version
or episode one  of the game program.  For purposes of this section, "use"
means loading the Software into RAM, as well as installation on a hard disk
or other storage  device.  You may not:  modify, translate, disassemble,
decompile, reverse engineer, or create derivative works based upon the
Software.  You agree thatd the Software will not be shipped, transferred or
exported into any country in violation of the U.S. Export Administration Act
and that you will not utilize, in any other manner, the Software in violation
of any applicable law.

   2.   Copyright.  The Software is owned by ID and is protected by United
States copyright laws and international treaty provisions.  You must treat
the Software like any other copyrighted material, except that you may make
copies of the Software to give to other persons.  You may not charge or
receive any consideration from any other person for the receipt or use of
the Software without receiving ID's prior written consent as specified in the
VENDOR.DOC file.  You agree to use your best efforts to see that any user of
the Software licensed hereunder complies with this Agreement.

   3.   Limited Warranty.  ID warrants that if properly installed and
operated on a computer for which it is designed, the Software will perform
substantially in accordance with its designed purpose for a period of ninety
(90) days from the date the Software is first obtained by an end-user.   ID's
entire liability and your exclusive remedy shall be, at ID's option, either
(a) return of the retail price paid, if any, or (b) repair or replacement of
the Software that does not meet ID's Limited Warranty.  To make a warranty
claim, return the Software to the point of purchase, accompanied by proof of
purchase, your name, your address, and a statement of defect, or return the
Software with the above information to ID.  This Limited Warranty is void if
failure of the Software has resulted in whole or in part from accident,
abuse, misapplication or violation of this Agreement.  Any replacement
Software will be warranted for the remainder of the original warranty period
or thirty (30) days, whichever is longer.  This warranty allocates risks of
product failure between Licensee and ID.  ID's product pricing reflects this
allocation of risk and the limitations of liability contained in this
warranty.

   4.   NO OTHER WARRANTIES.  ID DISCLAIMS ALL OTHER WARRANTIES, EITHER
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO, IMPLIED WARRANTIES OF
MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE WITH RESPECT TO THE
SOFTWARE AND THE ACCOMPANYING WRITTEN MATERIALS, IF ANY.  THIS LIMITED
WARRANTY GIVES YOU SPECIFIC LEGAL RIGHTS.  YOU MAY HAVE OTHERS WHICH VARY
FROM JURISDICTION TO JURISDICTION.  ID DOES NOT WARRANT THAT THE OPERATION
OF THE SOFTWARE WILL BE UNINTERRUPTED, ERROR FREE OR MEET LICENSEE'S
SPECIFIC REQUIREMENTS.  THE WARRANTY SET FORTH ABOVE IS IN LIEU OF ALL OTHER
EXPRESS WARRANTIES WHETHER ORAL OR WRITTEN.  THE AGENTS, EMPLOYEES,
DISTRIBUTORS, AND DEALERS OF ID ARE NOT AUTHORIZED TO MAKE MODIFICATIONS TO
THIS WARRANTY, OR ADDITIONAL WARRANTIES ON BEHALF OF ID.  ADDITIONAL
STATEMENTS SUCH AS DEALER ADVERTISING OR PRESENTATIONS, WHETHER ORAL OR
WRITTEN, DO NOT CONSTITUTE WARRANTIES BY ID AND SHOULD NOT BE RELIED UPON.

   5.   Exclusive Remedies.  You agree that your exclusive remedy against
ID, its affiliates, contractors, suppliers, and agents for loss or damage
caused by any defect or failure in the Software regardless of the form of
action, whether in contract, tort, including negligence, strict liability or
otherwise, shall be the return of the retail purchase price paid, if any, or
replacement of the Software.   This Agreement shall be construed in
accordance with and governed by the laws of the State of Texas.  Copyright
and other proprietary matters will be governed by United States laws and
international treaties.  IN ANY CASE, ID SHALL NOT BE LIABLE FOR LOSS OF
DATA, LOSS OF PROFITS, LOST SAVINGS, SPECIAL, INCIDENTAL, CONSEQUENTIAL,
INDIRECT OR OTHER SIMILAR DAMAGES ARISING FROM BREACH OF WARRANTY, BREACH OF
CONTRACT, NEGLIGENCE, OR OTHER LEGAL THEORY EVEN IF ID OR ITS AGENT HAS BEEN
ADVISED OF THE POSSIBILITY OF SUCH DAMAGES, OR FOR ANY CLAIM BY ANY OTHER
PARTY.  Some jurisdictions do not allow the exclusion or limitation of
incidental or consequential damages, so the above limitation or exclusion
may not apply to you.

   6.   General Provisions.  Neither this Agreement nor any part or portion
hereof shall be assigned or sublicensed, except as described herein.  Should
any provision of this Agreement be held to be void, invalid, unenforceable or
illegal by a court, the validity and enforceability of the other provisions
shall not be affected thereby.  If any provision is determined to be
unenforceable, you agree to a modification of such provision to provide for
enforcement of the provision's intent, to the extent permitted by applicable
law.  Failure of a party to enforce any provision of this Agreement shall not
constitute or be construed as a waiver of such provision or of the right to
enforce such provision.  If you fail to comply with any terms of this
Agreement, YOUR LICENSE IS AUTOMATICALLY TERMINATED.

   YOU ACKNOWLEDGE THAT YOU HAVE READ THIS AGREEMENT, YOU UNDERSTAND THIS
AGREEMENT, AND UNDERSTAND THAT BY CONTINUING THE INSTALLATION OF THE
SOFTWARE, BY LOADING OR RUNNING THE SOFTWARE, OR BY PLACING OR COPYING THE
SOFTWARE ONTO YOUR COMPUTER HARD DRIVE, YOU AGREE TO BE BOUND BY THIS
AGREEMENT'S TERMS AND CONDITIONS.  YOU FURTHER AGREE THAT, EXCEPT FOR WRITTEN
SEPARATE AGREEMENTS BETWEEN ID AND YOU, THIS AGREEMENT IS A COMPLETE AND
EXCLUSIVE STATEMENT OF THE RIGHTS AND LIABILITIES OF THE PARTIES.  THIS
AGREEMENT SUPERSEDES ALL PRIOR ORAL AGREEMENTS, PROPOSALS OR UNDERSTANDINGS,
AND ANY OTHER COMMUNICATIONS BETWEEN ID AND YOU RELATING TO THE SUBJECT
MATTER OF THIS AGREEMENT.

The above license does not appear to grant distribution permission.  Email
from John Carmack of ID Software provided this clarification:

X-Sender: johnc@mail.idsoftware.com
X-Mailer: Windows Eudora Pro Version 3.0 (32)
Date: Sat, 23 Oct 1999 20:01:30 -0500
To: Joe Drew <hoserhead@woot.net>
From: johnc@idsoftware.com (John Carmack)
Subject: Re: Doom shareware WAD license

At 08:02 PM 10/23/99 -0400, you wrote:
>Can you give me a definite license on the doom 1 shareware wad? I find certain
>things that say "freely distribute" and others that say "get vendor's license"
>... All I need to have is a license so I can package it up for Debian.
>Thanks.
>Joe

The DOOM shareware wad is freely distributable.  No Quake data is freely
distributable.

John Carmack
`

func BenchmarkFindWords(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Do something
		_ = FindWords(testText)
	}
}

func BenchmarkFindWords2(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Do something
		_ = FindWords2(testText)
	}
}

func BenchmarkSplitOnNonLetters(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Do something
		_ = SplitOnNonLetters(testText)
	}
}

func BenchmarkSplitOnNonLettersOrNumbers(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Do something
		_ = SplitOnNonLettersOrNumbers(testText)
	}
}

func BenchmarkSplitOnWords(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// Do something
		_ = SplitOnWords(testText)
	}
}

func TestResults(t *testing.T) {
	a := FindWords(testText)
	b := SplitOnWords(testText)

	if len(a) != len(b) {
		t.Fail()
	}

}
